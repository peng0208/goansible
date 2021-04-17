package goansible

type SynchronizeModuleArgs struct {
	//mirrors the rsync archive flag, enables recursive, links, perms, times, owner, group flags and -D.
	Archive  bool
	Checksum bool
	//compress file data during the transfer.
	Compress bool
	//copy symlinks as the item that they point to (the referent) is copied, rather than the symlink.
	CopyLinks bool `map:"copy_links"`
	//delete files that don't exist.
	Delete bool
	//path on the source host that will be synchronized to the destination.
	Src string
	//path on the destination host that will be synchronized from the source.
	Dest string
	//port number for ssh on the destination host.
	DestPort int `map:"dest_port"`
	//transfer directories without recursing.
	Dirs bool
	//copy symlinks as symlinks.
	Link bool
	//specify the direction of the synchronization. (choices: pull, push)
	Mode string
	//preserve group.
	Group bool
	//preserve owner.
	Owner bool
	//preserve perms.
	Perms bool
	//specify the private key to use for SSH-based rsync connections (e.g. `~/.ssh/id_rsa').
	PrivateKey string `map:"private_key"`
	//recurse into directories.
	Recursive bool
	//specify a `--timeout' for the rsync command in seconds.
	RsyncTimeout int `map:"rsync_timeout"`
	//preserve modification times.
	Times bool
	//put user@ for the remote paths.
	SetRemoteUser bool `map:"set_remote_user"`
	//use the ssh_args specified in ansible.cfg.
	UseSSHArgs bool `map:"use_ssh_args"`
	//verify destination host key.
	VerifyHost bool `map:"verify_host"`
}

func NewSynchronizeModule() *ModuleRunner {
	return &ModuleRunner{
		"synchronize",
		NewAdHocRunner(),
		&SynchronizeModuleArgs{},
		&ModuleAdHocOptions{},
		&ConnectionOptions{},
		&PrivilegeOptions{},
	}
}
