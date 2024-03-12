package main

import (
	"os"

	"github.com/galasa-dev/cli/pkg/files"
	"github.com/techcobweb/openapi2beans/pkg/cmd"
)

func main() {
	args := os.Args[1:]
	fs := files.NewOSFileSystem()
	rootCmd := cmd.NewRootCommand(fs, cmd.Openapi2beansFlagStore{})
	rootCmd.SetArgs(args)

	// Execute the command
	err := rootCmd.Execute()

	if err != nil {
		panic(err)
	}
}
