package goansible

type ScriptModuleArgs struct {
	//remote shell script to run
	Cmd string `map:"free_form"`
	//change into this directory before running the script
	Chdir string
	//the executable name or path to execute the script
	Executable string
}

func NewScriptModule() *ModuleRunner {
	return &ModuleRunner{
		"script",
		NewAdHocRunner(),
		&ScriptModuleArgs{},
		&ModuleAdHocOptions{},
		&ConnectionOptions{},
		&PrivilegeOptions{},
	}
}
