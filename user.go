package goansible

type UserModuleArgs struct {
	//name of the user to create, remove or modify.
	Name string
	//append user to groups,if false,remove them from all other groups.
	Append bool
	//description of user account
	Comment string
	//create user home directory
	CreateHome bool `map:"create_home"`
	Force      bool
	//user's primary group
	Group string
	//list of groups user will be added to,separate by commaparated
	Groups string
	//user home directory
	Home string
	//remove directories, same as `userdel --remove'
	Remove bool
	//user's shell
	Shell string
	State string
	//UID
	Uid int
}

func NewUserModule() *ModuleRunner {
	return &ModuleRunner{
		"user",
		NewAdHocRunner(),
		&UserModuleArgs{},
		&ModuleAdHocOptions{},
		&ConnectionOptions{},
		&PrivilegeOptions{},
	}
}
