package cmd

import (
	"github.com/galasa-dev/cli/pkg/files"
	"github.com/spf13/cobra"
	"github.com/techcobweb/openapi2beans/pkg/generator"
)

type Openapi2beansFlagStore struct {
	apiFilePath   string
	packageNane   string
	storeFilePath string
}

func NewRootCommand(fs files.FileSystem, flags Openapi2beansFlagStore) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "openapi2beans",
		Short: "CLI for openapi2beans",
		Long:  "A tool for generating java beans from an openapi yaml file.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return generator.GenerateFiles(fs, flags.storeFilePath, flags.apiFilePath, flags.packageNane)
		},
	}

	cmd.Flags().StringVarP(&flags.apiFilePath, "yaml", "y", "", "Specifies where to pull the openapi yaml from.")
	cmd.Flags().StringVarP(&flags.packageNane, "package", "p", "generated", "Specifies what package the Java files belong to.")
	cmd.Flags().StringVarP(&flags.storeFilePath, "store", "s", "generated", "Specifies the file path to store the resulting generated java beans.")

	cmd.MarkFlagRequired("yaml")
	cmd.MarkFlagRequired("package")
	cmd.MarkFlagRequired("store")
	return cmd
}