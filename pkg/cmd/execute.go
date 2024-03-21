package cmd

import "github.com/techcobweb/openapi2beans/pkg/utils"


func Execute(factory utils.Factory, args []string) error {
	rootCmd := NewRootCommand(factory, Openapi2beansFlagStore{})
	rootCmd.SetArgs(args)

	// Execute the command
	err := rootCmd.Execute()

	return err
}
