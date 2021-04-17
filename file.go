package goansible

type FileModuleArgs struct {
	//filesystem links, if they exist, should be followed.
	Follow bool
	//force the creation of the symlinks
	Force bool
	//name of the user that should own the file/directory, as would be fed to `chown'.
	Owner string
	//name of the group that should own the file/directory, as would be fed to `chown'.
	Group string
	//the permissions the resulting file or directory should have, as would be fed to `chmod'
	Mode string
	//the dst file path
	Path string
	//recursively set the specified file attributes on directory contents.
	Recurse bool
	//the src file path
	Src string
	//the dst file state
	State string
}

func NewFileModule() *ModuleRunner {
	return &ModuleRunner{
		"file",
		NewAdHocRunner(),
		&FileModuleArgs{},
		&ModuleAdHocOptions{},
		&ConnectionOptions{},
		&PrivilegeOptions{},
	}
}
