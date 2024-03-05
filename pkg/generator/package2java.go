package generator

import (
	"log"

	"github.com/cbroglie/mustache"
	"github.com/galasa-dev/cli/pkg/files"
)

// func createJavaPackageFiles(fs files.FileSystem, storeFilepath string, apiFilePath string, javaPackage JavaPackage) {

// }

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

// plugs a java class into a moustache template and saves the resulting string in a file
func createJavaClassFile(javaClass JavaClass, fs files.FileSystem, storeFilepath string) error {
	log.Println("Creating bean: " + javaClass.Name + ".java")
	generatedBeanFileContents, err := mustache.RenderFile("../../templates/JavaClassTemplate.mustache", javaClass)
	if err == nil {
		err = fs.WriteTextFile(storeFilepath+"/"+javaClass.Name+".java", generatedBeanFileContents)
	}
	return err
}