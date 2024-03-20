package generator

import (
	"log"
	"strings"

	"github.com/galasa-dev/cli/pkg/files"
)

func GenerateFiles(fs files.FileSystem, projectFilePath string, apiFilePath string, packageName string) error {
	var fatalErr error
	var apiyaml string
	var errList []error

	storeFilePath := generateStoreFilePath(projectFilePath, packageName)
	fatalErr = generateDirectories(fs, storeFilePath)
	if fatalErr == nil {
		apiyaml, fatalErr = fs.ReadTextFile(apiFilePath)
		if fatalErr == nil {
			var schemaTypes map[string]*SchemaType
			schemaTypes, errList, fatalErr = getSchemaTypesFromYaml([]byte(apiyaml))
			if fatalErr == nil {
				javaPackage := translateSchemaTypesToJavaPackage(schemaTypes, packageName)
				convertJavaPackageToJavaFiles(javaPackage, fs, storeFilePath)
			}
		}
	}

	handleErrList(errList)
	return fatalErr
}

// Cleans and/or creates the store file
func generateDirectories(fs files.FileSystem, storeFilePath string) error {
	log.Println("Cleaning generated beans directory: " + storeFilePath)
	exists, err := fs.DirExists(storeFilePath)
	if err == nil {
		if exists {
			fs.DeleteDir(storeFilePath)
		}
		log.Printf("Creating output directory: %s\n", storeFilePath)
		err = fs.MkdirAll(storeFilePath)
	}
	return err
}

func handleErrList(errList []error) {
	log.Println("Listing non-fatal errors:")
	for _, individualError := range errList {
		log.Println(individualError.Error())
	}
}

func generateStoreFilePath(projectFilePath string, packageName string) string {
	packageFilePath := strings.ReplaceAll(packageName, ".", "/")
	return projectFilePath + "/" + packageFilePath
}