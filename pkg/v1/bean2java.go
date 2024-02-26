package v1_generator

import (
	"log"
	"strings"

	"github.com/cbroglie/mustache"
	"github.com/galasa-dev/cli/pkg/files"
)

func GenerateBeans(fs files.FileSystem, storeFilepath string, apiFilePath string, packageName string) error {
	var err error
	var apiyaml string
	var beanList []Bean

	err = generateDirectories(fs, storeFilepath)
	if err == nil {
		apiyaml, err = fs.ReadTextFile(apiFilePath)
		if err == nil {
			if strings.Contains(apiyaml, "schema:") {
				apiyaml = strings.Split(apiyaml, "schema:")[1]
			}
			beanList, err = getBeansFromYaml([]byte(apiyaml), packageName)
			if err == nil {
				for _, bean := range beanList {
					nonFatalErr := createBeanFile(bean, fs, storeFilepath)
					if err != nil {
						log.Println("Error: failed to create bean: " + bean.Object.GetName() + ". Reason is: " + nonFatalErr.Error())
					}
				}
			}
		}
	}

	return err
}

// Cleans and/or creates the store file
func generateDirectories(fs files.FileSystem, storeFilepath string) error {
	log.Println("Cleaning generated beans directory: " + storeFilepath)
	exists, err := fs.DirExists(storeFilepath)
	if err == nil {
		if exists {
			fs.DeleteDir(storeFilepath)
		}
		err = fs.MkdirAll(storeFilepath)
	}
	return err
}

// Uses go's templating engine to generate Beans u
func createBeanFile(bean Bean, fs files.FileSystem, storeFilepath string) error {
	log.Println("Creating bean: " + bean.Object.varName + ".java")
	generatedBeanFileContents, err := mustache.RenderFile("../../templates/Javabean", bean)
	if err == nil {
		err = fs.WriteTextFile(storeFilepath+"/"+bean.Object.varName+".java", generatedBeanFileContents)
	}
	return err
}
