package goansible

import (
	"encoding/json"
	"os/exec"
	"sync"
)

type Runner interface {
	//execute ansible task
	Run() (*Result, error)
	//only show ansible task command to run,  don't execute
	CommandString() string
}

type (
	CommonOptions struct {
		//inventory hosts path or comma separated host list
		Inventory string
		//prepend colon-separated path(s) to module library
		ModulePath string
		//don't make any changes
		Check bool
		//set additional variables as k:v
		ExtraVars map[string]string
		//number of parallel processes
		Forks int
	}

	AdHocOptions struct {
		CommonOptions
		//host pattern
		Pattern string
		//module name to execute
		ModuleName string
		//module arguments string, map, or struct
		//if use map, free form parameter is actual parameter named 'free_form'.
		//but use string, free form parameter is no actual parameter
		ModuleArgs interface{}
		//async timeout, seconds
		BackgroundSeconds int
		//async poll interval, seconds
		PollInterval int
	}

	PlaybookOptions struct {
		CommonOptions
		//playbook file path
		Playbook string
		//clear the fact cache for every host in inventory
		FlushCache bool
		//only run plays and tasks tagged with these values
		Tags string
	}

	ConnectionOptions struct {
		//ssh key file, use this file to authenticate the connection
		PrivateKey string
		//connection timeout, seconds
		Timeout int
		//connection type to use (default=smart)
		Connection string
		//connect as this user
		RemoteUser string
	}

	PrivilegeOptions struct {
		//privilege escalation method to use (default=sudo)
		BecomeMethod string
		//run operations as this user (default=root)
		BecomeUser string
		//run operations with become
		Become bool
	}
)

type (
	AdHocRunner struct {
		Bin               string
		Options           *AdHocOptions
		ConnectionOptions *ConnectionOptions
		PrivilegeOptions  *PrivilegeOptions
		mutex             sync.Mutex
	}

	PlaybookRunner struct {
		Bin               string
		Options           *PlaybookOptions
		ConnectionOptions *ConnectionOptions
		PrivilegeOptions  *PrivilegeOptions
		mutex             sync.Mutex
	}
)

type (
	//ansible running results
	Result struct {
		CustomStats       map[string]interface{} `json:"custom_stats"`
		GlobalCustomStats map[string]interface{} `json:"global_custom_stats"`
		Plays             []*Play                `json:"plays"`
		Summary           map[string]interface{} `json:"stats"`
	}

	Play struct {
		Play  map[string]interface{} `json:"play"`
		Tasks []*Task                `json:"tasks"`
	}

	Task struct {
		Hosts map[string]interface{} `json:"hosts"`
		Task  map[string]interface{} `json:"task"`
	}
)

func NewAdHocRunner() *AdHocRunner {
	return &AdHocRunner{
		Options:           &AdHocOptions{},
		ConnectionOptions: &ConnectionOptions{},
		PrivilegeOptions:  &PrivilegeOptions{},
	}
}

func NewPlaybookRunner() *PlaybookRunner {
	return &PlaybookRunner{
		Options:           &PlaybookOptions{},
		ConnectionOptions: &ConnectionOptions{},
		PrivilegeOptions:  &PrivilegeOptions{},
	}
}

func (r *AdHocRunner) Run() (*Result, error) {
	return execute(r.command())
}

func (r *AdHocRunner) CommandString() string {
	return r.command().String()
}

func (r *AdHocRunner) command() *exec.Cmd {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	cmdArgs := genAdHocCommand(r)
	return executor(
		cmdArgs[0],
		cmdArgs[1:len(cmdArgs)]...,
	)
}

func (r *PlaybookRunner) Run() (*Result, error) {
	return execute(r.command())
}

func (r *PlaybookRunner) CommandString() string {
	return r.command().String()
}

func (r *PlaybookRunner) command() *exec.Cmd {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	cmdArgs := genPlaybookCommand(r)
	return executor(
		cmdArgs[0],
		cmdArgs[1:len(cmdArgs)]...,
	)
}

func executor(name string, args ...string) *exec.Cmd {
	return exec.Command(name, args...)
}

func execute(cmd *exec.Cmd) (*Result, error) {
	output, err := cmd.CombinedOutput()
	//returncode
	//rc := comm.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	var data Result
	err = json.Unmarshal(output, &data)
	return &data, err
}

func (r *Result) String() string {
	data, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
