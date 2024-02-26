package v1_generator

import (
	"testing"

	"github.com/cbroglie/mustache"
	"github.com/galasa-dev/cli/pkg/files"
	"github.com/stretchr/testify/assert"
)

const (
	TARGET_JAVA_PACKAGE = "main"
)

func AssertFileGeneratedOk(t *testing.T, mockFileSystem files.FileSystem, storeFilepath string, generatedCodeFilepath string, objectName string) {
	exists, err := mockFileSystem.Exists(generatedCodeFilepath)
	assert.Nil(t, err)
	assert.True(t, exists)
	generatedFile, err := mockFileSystem.ReadTextFile(generatedCodeFilepath)
	assert.Nil(t, err)
	assert.Contains(t, generatedFile, "public class "+objectName)
	assert.Contains(t, generatedFile, "public "+objectName+" ()")
}

func TestGeneratorCreatesGeneratedDirectory(t *testing.T) {
	mockFileSystem := files.NewMockFileSystem()
	err := generateDirectories(mockFileSystem, "../generated")

	assert.Nil(t, err)
	exists, err := mockFileSystem.DirExists("../generated")
	assert.Nil(t, err)
	assert.True(t, exists)
}

func TestGeneratorReturnsNoErrorWhenFilepathDirectoryAlreadyExists(t *testing.T) {
	mockFileSystem := files.NewMockFileSystem()
	mockFileSystem.MkdirAll("../generated")
	err := generateDirectories(mockFileSystem, "../generated")

	assert.Nil(t, err)
	exists, err := mockFileSystem.DirExists("../generated")
	assert.Nil(t, err)
	assert.True(t, exists)
}

func TestGeneratorReturnsNoErrorWhenFilepathDirectoryAlreadyExistsWithFileIn(t *testing.T) {
	// Given...
	storeFilepath := "generated"
	mockFileSystem := files.NewMockFileSystem()
	mockFileSystem.MkdirAll(storeFilepath)
	mockFileSystem.WriteTextFile(storeFilepath+"/test.txt", "this is but a test, good luck and godspeed")

	// When...
	err := generateDirectories(mockFileSystem, storeFilepath)

	// Then...
	assert.Nil(t, err)
	exists, err := mockFileSystem.DirExists(storeFilepath)
	assert.Nil(t, err)
	assert.True(t, exists)
	exists, err = mockFileSystem.Exists(storeFilepath + "/test.txt")
	assert.Nil(t, err)
	assert.False(t, exists)
}

func TestGenerateBeansCreatesEmptyObjectBean(t *testing.T) {
	// Given...
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	apiFilePath := "test-resources/single-bean.yaml"
	objectName := "MyBeanName"
	generatedCodeFilePath := storeFilepath + "/" + objectName + ".java"
	testapiyaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
`
	mockFileSystem.WriteTextFile(apiFilePath, testapiyaml)

	// When...
	err := GenerateBeans(mockFileSystem, storeFilepath, apiFilePath, TARGET_JAVA_PACKAGE)

	// Then...
	assert.Nil(t, err)
	AssertFileGeneratedOk(t, mockFileSystem, storeFilepath, generatedCodeFilePath, objectName)
}

func TestTemplateAcceptsBeanStructure(t *testing.T) {
	// Given...
	mockFileSystem := files.NewMockFileSystem()
	storeFilepath := "generated"
	objectName := "JsonError"
	generatedCodeFilePath := storeFilepath + "/" + objectName + ".java"

	var bean Bean
	bean.Object.varName = objectName
	bean.BeanPackage = "generated"
	bean.Object.description = "this is a blank bean"
	bean.Object.varTypeName = "object"

	// When...
	err := createBeanFile(bean, mockFileSystem, storeFilepath)

	// Then...
	assert.Nil(t, err)
	AssertFileGeneratedOk(t, mockFileSystem, storeFilepath, generatedCodeFilePath, objectName)
}

type Planet struct {
	PlanetName string
}
func (planet Planet) GetName() (string) {
	return planet.PlanetName
}

func TestExploringMoustacheMore(t *testing.T) {
	objectName := "BeanName"
	var bean Bean
	bean.Object.varName = objectName
	bean.BeanPackage = TARGET_JAVA_PACKAGE
	bean.Object.description = "this is a blank bean"

	randoMap := map[string]string{
		"planet": "earth",
	}
	result, err := mustache.Render("{{planet}}", randoMap)

	assert.Nil(t, err)
	assert.Equal(t, result, "earth")


	planet := Planet {
		PlanetName: "earth",
	}

	result, err = mustache.Render("We are on {{PlanetName}}", planet)

	assert.Nil(t, err)
	assert.Equal(t, "We are on earth", result)
}