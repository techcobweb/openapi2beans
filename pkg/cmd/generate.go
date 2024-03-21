package cmd

import (
	"github.com/galasa-dev/cli/pkg/files"
	galasaUtils "github.com/galasa-dev/cli/pkg/utils"
	"github.com/spf13/cobra"
	openapi2beans_errors "github.com/techcobweb/openapi2beans/pkg/errors"
	"github.com/techcobweb/openapi2beans/pkg/generator"
	"github.com/techcobweb/openapi2beans/pkg/utils"
	"gopkg.in/yaml.v3"
)

func NewGenerateCommand(factory utils.Factory, flags Openapi2beansFlagStore) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generates java from openapi yaml",
		Long:  "command used to generate java from an openapi yaml input.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return executeGenerateCmd(factory, &flags)
		},
	}

	addFlags(cmd, &flags)
	cmd.Flags().BoolP("help", "h", false, "Displays the options for the 'openapi2beans' command.")
	cmd.MarkFlagsOneRequired("yaml", "config")
	cmd.MarkFlagsMutuallyExclusive("yaml", "config")
	cmd.MarkFlagsOneRequired("package", "config")
	cmd.MarkFlagsMutuallyExclusive("package", "config")
	cmd.MarkFlagsOneRequired("output", "config")
	cmd.MarkFlagsMutuallyExclusive("output", "config")

	return cmd
}

func executeGenerateCmd(factory utils.Factory, flags *Openapi2beansFlagStore) error {
	var err error
	fs := factory.GetFileSystem()
	err = galasaUtils.CaptureLog(fs, flags.logFileName)
	if err == nil {
		handleConfig(fs, flags)
		err = generator.GenerateFiles(fs, flags.storeFilePath, flags.apiFilePath, flags.packageName)
	}
	return err
}

func handleConfig(fs files.FileSystem, flags *Openapi2beansFlagStore) error {
	var err error
	var configFlagValues *Openapi2beansFlagStore

	configFlagValues, err = processConfigFile(fs, flags.configFileName)

	if flags.apiFilePath == "" && configFlagValues.apiFilePath != "" {
		flags.apiFilePath = configFlagValues.apiFilePath
	} else {
		err = openapi2beans_errors.NewError("handleConfig: unable to find yaml value from neither flag nor config.")
	}

	if flags.packageName == "" && configFlagValues.packageName != "" {
		flags.packageName = configFlagValues.packageName
	} else {
		err = openapi2beans_errors.NewError("handleConfig: unable to find package value from neither flag nor config.")
	}

	if flags.storeFilePath == "" && configFlagValues.storeFilePath != "" {
		flags.storeFilePath = configFlagValues.storeFilePath
	} else {
		err = openapi2beans_errors.NewError("handleConfig: unable to find output value from neither flag nor config.")
	}

	return err
}

func processConfigFile(fs files.FileSystem, configFileName string) (*Openapi2beansFlagStore, error) {
	var err error
	var isConfigFileExisting bool
	var configFile string
	var configFlagValues *Openapi2beansFlagStore

	isConfigFileExisting, err = fs.Exists(configFileName)
	if err == nil {
		if isConfigFileExisting {
			configFile, err = fs.ReadTextFile(configFileName)
			if err == nil {
				err = yaml.Unmarshal([]byte(configFile), &configFlagValues)
			}
		}
	}
	return configFlagValues, err
}
