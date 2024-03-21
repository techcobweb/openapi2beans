package cmd

import (
	"github.com/spf13/cobra"
	"github.com/techcobweb/openapi2beans/pkg/utils"
)

type Openapi2beansFlagStore struct {
	apiFilePath    string
	packageName    string
	storeFilePath  string
	logFileName    string
	configFileName string
}

func NewRootCommand(factory utils.Factory, flags Openapi2beansFlagStore) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "openapi2beans",
		Short: "CLI for openapi2beans",
		Long:  "A tool for generating java beans from an openapi yaml file.",
		SilenceUsage: true,
	}

	cmd.SetErr(factory.GetStdErrConsole())
	cmd.SetOut(factory.GetStdOutConsole())

	cmd.Flags().BoolP("help", "h", false, "Displays the options for the 'openapi2beans' command.")
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	addFlags(cmd, &flags)

	addChildCommands(factory, flags, cmd)

	return cmd
}

func addChildCommands(factory utils.Factory, flags Openapi2beansFlagStore, rootCmd *cobra.Command) {
	generateCmd := NewGenerateCommand(factory, flags)
	rootCmd.AddCommand(generateCmd)

	configCmd := NewConfigCommand(factory, flags)
	rootCmd.AddCommand(configCmd)
}

func addFlags(cmd *cobra.Command, flagStore *Openapi2beansFlagStore) {
	cmd.PersistentFlags().StringVarP(&flagStore.apiFilePath, "yaml", "y", "", "Specifies where to pull the openapi yaml from.")
	cmd.PersistentFlags().StringVarP(&flagStore.packageName, "package", "p", "generated", "Specifies what package the Java files belong to. Directories will be generated in accordance.")
	cmd.PersistentFlags().StringVarP(&flagStore.storeFilePath, "output", "o", "generated", "Specifies the file path to store the resulting generated java beans.")
	cmd.PersistentFlags().StringVarP(&flagStore.logFileName, "log", "l", "-", "Specifies the output file for logs.")
	cmd.PersistentFlags().StringVarP(&flagStore.configFileName, "config", "", "c", "Specifies location of config file.")
}