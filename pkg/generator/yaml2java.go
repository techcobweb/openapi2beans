package generator

import (
	"log"

	"github.com/galasa-dev/cli/pkg/files"
)

func GenerateFiles(fs files.FileSystem, storeFilepath string, apiFilePath string, packageName string) error {
	var err error
	var apiyaml string

	err = generateDirectories(fs, storeFilepath)
	if err == nil {
		apiyaml, err = fs.ReadTextFile(apiFilePath)
		if err == nil {
			var schemaTypes map[string]*SchemaType
			schemaTypes, err = getSchemaTypesFromYaml([]byte(apiyaml))
			javaPackage := translateSchemaTypesToJavaPackage(schemaTypes, packageName)
			convertJavaPackageToJavaFiles(javaPackage, fs, storeFilepath)
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
		log.Printf("Creating output directory: %s\n", storeFilepath)
		err = fs.MkdirAll(storeFilepath)
	}
	return err
}