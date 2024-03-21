package cmd

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/techcobweb/openapi2beans/pkg/utils"
	"gopkg.in/yaml.v3"
)

func NewConfigCommand(factory utils.Factory, flags Openapi2beansFlagStore) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "CLI for openapi2beans",
		Long:  "A tool for generating java beans from an openapi yaml file.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return exectuteConfigCommand(factory, &flags)
		},
	}

	cmd.Flags().BoolP("help", "h", false, "Displays the options for the 'config' command.")
	addFlags(cmd, &flags)
	cmd.MarkPersistentFlagRequired("config")
	cmd.MarkFlagsOneRequired("yaml", "package", "output")

	return cmd
}

func exectuteConfigCommand(factory utils.Factory, flags *Openapi2beansFlagStore) error {
	var err error
	var configFlags Openapi2beansFlagStore
	var newConfigContents []byte

	fs := factory.GetFileSystem()
	exists, err := fs.Exists(flags.configFileName)

	if err == nil {
		if exists {
			var configFile []byte
			configFile, err = fs.ReadBinaryFile(flags.configFileName)
			if err == nil {
				err = yaml.Unmarshal(configFile, configFlags)
			}
		} else {
			var w io.WriteCloser
			w, err = fs.Create(flags.configFileName)
			w.Close()
		}
	}

	if err == nil {
		setConfigFlags(flags, &configFlags)
		newConfigContents, err = yaml.Marshal(configFlags)
		if err == nil {
			err = fs.WriteBinaryFile(flags.configFileName, newConfigContents)
		}
	}

	return err
}

func setConfigFlags(cliFlags *Openapi2beansFlagStore, configFlags *Openapi2beansFlagStore) {
	if cliFlags.apiFilePath != "" {
		configFlags.apiFilePath = cliFlags.apiFilePath
	}
	if cliFlags.logFileName != "" {
		configFlags.logFileName = cliFlags.logFileName
	}
	if cliFlags.packageName != "" {
		configFlags.packageName = cliFlags.packageName
	}
	if cliFlags.storeFilePath != "" {
		configFlags.storeFilePath = cliFlags.storeFilePath
	}
}