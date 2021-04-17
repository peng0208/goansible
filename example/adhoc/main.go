package main

import (
	"fmt"
	"github.com/peng0208/goansible"
)

func main() {
	//ansible adhoc模式用法

	runner := goansible.NewAdHocRunner()
	runner.Options.Inventory = "127.0.0.1,"
	runner.Options.Pattern = "all"
	runner.Options.ModuleName = "shell"
	//module args map
	runner.Options.ModuleArgs = map[string]interface{}{
		"chdir":     "/tmp",
		"free_form": "pwd",
	}
	//also use string
	//runner.Options.ModuleArgs = "chdir=/tmp pwd"

	//run only locally
	runner.ConnectionOptions.Connection = "local"

	output, err := runner.Run()
	fmt.Printf("output: %v\nerror: %v\n", output.String(), err)

}
