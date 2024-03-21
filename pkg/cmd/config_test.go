package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/techcobweb/openapi2beans/pkg/utils"
)

func TestSetConfigFlagOverridesConfigFlags(t *testing.T) {
	// Given...
	configFlags := Openapi2beansFlagStore {
		apiFilePath: "overrideapi",
		packageName: "override",
		storeFilePath: "shouldnt be here",
		logFileName: "oopedie doopedie",
	}
	cliFlags := Openapi2beansFlagStore {
		apiFilePath: "apifile",
		packageName: "package name",
		storeFilePath: "store",
		logFileName: "-",
	}

	// When...
	setConfigFlags(&cliFlags, &configFlags)

	// Then...
	assert.Equal(t, cliFlags.apiFilePath, configFlags.apiFilePath)
	assert.Equal(t, cliFlags.packageName, configFlags.packageName)
	assert.Equal(t, cliFlags.storeFilePath, configFlags.storeFilePath)
	assert.Equal(t, cliFlags.logFileName, configFlags.logFileName)
}

// func TestExectuteConfigCommandWithoutExistingConfigFileReturnsOk(t *testing.T) {
// 	// Given...
// 	factory := utils.NewMockFactory()
// 	cliFlags := Openapi2beansFlagStore {
// 		apiFilePath: "apifile",
// 		packageName: "package name",
// 		storeFilePath: "store",
// 		logFileName: "-",
// 		configFileName: "config.yaml",
// 	}


// 	// When...
// 	err := exectuteConfigCommand(factory, &cliFlags)

// 	// Then...
// 	assert.Nil(t, err)
// 	fs := factory.GetFileSystem()
// 	output, err := fs.ReadTextFile(cliFlags.configFileName)
// 	assert.Nil(t, err)
// 	expectedOutput := `
// apiFilePath: "apifile"
// packageName: "package name"
// storeFilePath: "store"
// logFileName: "-"
// `
// 	assert.Equal(t, expectedOutput, output)
// }

func TestConfigWillOutputHelp(t *testing.T) {
	// Given...
	args := []string{"config"}
	mockFactory := utils.NewMockFactory()

	// When...
	err := Execute(mockFactory, args)

	// Then...
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "required flag(s) \"config\" not set")
	checkOutput("", "Error: required flag(s) \"config\" not set", mockFactory, t)
}