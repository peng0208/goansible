package main

import (
	"fmt"
	"github.com/peng0208/goansible"
)

func main() {
	//ansible模块用法，此例为shell模块
	shell := goansible.NewShellModule()
	shell.Options.Inventory = "127.0.0.1,"
	shell.Options.Pattern = "all"
	shell.ConnectionOptions.Connection = "local"

	//module args struct, every module has a struct "XXXModuleArgs"
	shell.Args = goansible.ShellModuleArgs{
		Cmd:   "pwd",
		Chdir: "/tmp",
	}

	//alse use map
	//shell.Args = map[string]interface{}{
	//	"chdir":     "/tmp",
	//	"free_form": "pwd",
	//}
	//also use string
	//shell.Args = "chdir=/tmp pwd"

	fmt.Println(shell.CommandString())
}
