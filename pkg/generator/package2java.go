package generator

import (
	"log"

	"github.com/cbroglie/mustache"
	"github.com/galasa-dev/cli/pkg/files"
)

// func generateDirectories(fs files.FileSystem, storeFilepath string) error {
// 	log.Println("Cleaning generated beans directory: " + storeFilepath)
// 	exists, err := fs.DirExists(storeFilepath)
// 	if err == nil {
// 		if exists {
// 			fs.DeleteDir(storeFilepath)
// 		}
// 		err = fs.MkdirAll(storeFilepath)
// 	}
// 	return err
// }

func convertJavaPackageToJavaFiles(javaPackage *JavaPackage, fs files.FileSystem, storeFilePath string) {
	for _, javaClass := range javaPackage.Classes {
		createJavaClassFile(javaClass, fs, storeFilePath)
	}
	for _, javaEnum := range javaPackage.Enums {
		createJavaEnumFile(javaEnum, fs, storeFilePath)
	}
}

// plugs a java class into a moustache template and saves the resulting string in a file
func createJavaClassFile(javaClass *JavaClass, fs files.FileSystem, storeFilepath string) error {
	log.Println("Creating class: " + javaClass.Name + ".java")
	generatedBeanFileContents, err := mustache.RenderFile("../../templates/JavaClassTemplate.mustache", javaClass)
	if err == nil {
		err = fs.WriteTextFile(storeFilepath+"/"+javaClass.Name+".java", generatedBeanFileContents)
		if err == nil {
			log.Print("Successfully created enum: " + javaClass.Name + ".java")
		} else {
			log.Printf("Error: failed to create: %s. Reason is: %s", javaClass.Name + ".java", err.Error())
		}
	} else {
		log.Print("Error rendering file.")
	}
	return err
}

func createJavaEnumFile(javaEnum *JavaEnum, fs files.FileSystem, storeFilepath string) error {
	log.Println("Creating enum: " + javaEnum.Name + ".java")
	generatedBeanFileContents, err := mustache.RenderFile("../../templates/JavaEnumTemplate.mustache", javaEnum)
	if err == nil {
		err = fs.WriteTextFile(storeFilepath+"/"+javaEnum.Name+".java", generatedBeanFileContents)
		if err == nil {
			log.Print("Successfully created enum: " + javaEnum.Name + ".java")
		} else {
			log.Printf("Error: failed to create: %s. Reason is: %s", javaEnum.Name + ".java", err.Error())
		}
	} else {
		log.Print("Error rendering file.")
	}
	return err
}