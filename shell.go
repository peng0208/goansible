package goansible

type ShellModuleArgs struct {
	//remote shell command to run
	Cmd string `map:"free_form"`
	//change into this directory before running the command
	Chdir string
	//the shell used to execute the command
	Executable string
}

func NewShellModule() *ModuleRunner {
	return &ModuleRunner{
		"shell",
		NewAdHocRunner(),
		&ShellModuleArgs{},
		&ModuleAdHocOptions{},
		&ConnectionOptions{},
		&PrivilegeOptions{},
	}
}
