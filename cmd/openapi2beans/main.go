package main

import (
	"os"

	"github.com/techcobweb/openapi2beans/pkg/utils"
	"github.com/techcobweb/openapi2beans/pkg/cmd"
)

func main() {
	args := os.Args[1:]
	factory := utils.NewRealFactory()
	
	err := cmd.Execute(factory, args)

	if err != nil {
		panic(err)
	}
}
