package generator

import (
	"testing"

	"github.com/galasa-dev/cli/pkg/files"
	"github.com/stretchr/testify/assert"
)

func TestGenerateFilesProducesFileFromSingleGenericObjectSchema(t *testing.T) {
	// Given...
	packageName := "generated"
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "dev/wyvinar"
	apiFilePath := "test-resources/single-bean.yaml"
	objectName := "MyBeanName"
	generatedCodeFilePath := storeFilepath + "/" + packageName + "/" + objectName + ".java"
	testapiyaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
`
	mockFileSystem.WriteTextFile(apiFilePath, testapiyaml)
	
	// When...
	err := GenerateFiles(mockFileSystem, storeFilepath, apiFilePath, packageName)

	// Then...
	assert.Nil(t, err)
	assert.Nil(t, errList)
	generatedClassFile := openGeneratedFile(t, mockFileSystem, generatedCodeFilePath)
	assertClassFileGeneratedOk(t, generatedClassFile, "MyBeanName")
}