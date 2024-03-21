package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/techcobweb/openapi2beans/pkg/utils"
)

func TestCGenerateWillExecute(t *testing.T) {
	// Given...
	args := []string{"generate", "--help"}
	mockFactory := utils.NewMockFactory()

	// When...
	err := Execute(mockFactory, args)

	// Then...
	assert.Nil(t, err)
	checkOutput("Usage:", "", mockFactory, t)
}

func TestGenerateWillReturnErrorsWithoutPackageOrConfigSet(t *testing.T) {
	// Given...
	args := []string{"generate", "--yaml", "somewhere.yaml", "--output", "target.project"}
	mockFactory := utils.NewMockFactory()

	// When...
	err := Execute(mockFactory, args)

	// Then...
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "at least one of the flags in the group [package config] is required")
	checkOutput("", "Error: at least one of the flags in the group [package config] is required", mockFactory, t)
}

func TestGenerateWillReturnErrorsWithoutOutputOrConfigSet(t *testing.T) {
	// Given...
	args := []string{"generate", "--yaml", "somewhere.yaml", "--package", "packageName"}
	mockFactory := utils.NewMockFactory()

	// When...
	err := Execute(mockFactory, args)

	// Then...
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "at least one of the flags in the group [output config] is required")
	checkOutput("", "Error: at least one of the flags in the group [output config] is required", mockFactory, t)
}

func TestGenerateWillReturnErrorsWithoutYamlOrConfigSet(t *testing.T) {
	// Given...
	args := []string{"generate", "--package", "packageName", "--output", "target.project"}
	mockFactory := utils.NewMockFactory()

	// When...
	err := Execute(mockFactory, args)

	// Then...
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "at least one of the flags in the group [yaml config] is required")
	checkOutput("", "Error: at least one of the flags in the group [yaml config] is required", mockFactory, t)
}

func TestGenerateWillReturnOkWithAllFlagsSet(t *testing.T) {
	// Given...
	args := []string{"generate", "--package", "packageName", "--output", "target/project", "--yaml", "test.yaml"}
	mockFactory := utils.NewMockFactory()
	mockFileSystem := mockFactory.GetFileSystem()
	testapiyaml := `openapi: 3.0.3
components:
  schemas:
    MyBeanName:
      type: object
`
	mockFileSystem.WriteTextFile("test.yaml", testapiyaml)


	// When...
	err := Execute(mockFactory, args)

	// Then...
	assert.Nil(t, err)
	checkOutput("", "", mockFactory, t)
}