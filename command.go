package goansible

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	//Default Value
	DefaultAnsibleBin         = "ansible"
	DefaultAnsiblePlaybookBin = "ansible-playbook"
	DefaultPattern            = "all"
	DefaultForks              = 10
	DefaultPollInterval       = 5
	EmptyValue                = ""
	ZeroValue                 = 0
	FreeFormValue             = "free_form"

	//Optional Arguments Flag
	Inventory         = "-i"
	Check             = "-C"
	ExtraVars         = "-e"
	Forks             = "-f"
	ModulePath        = "-M"
	ModuleName        = "-m"
	ModuleArgs        = "-a"
	FlushCache        = "--flush-cache"
	Tags              = "-t"
	BackgroundSeconds = "-B"
	PollInterval      = "-P"

	//Connection Options Flag
	PrivateKey = "--private-key"
	Timeout    = "-T"
	Connection = "-c"
	RemoteUser = "-u"

	//Privilege Escalation Options Flag
	BecomeMethod = "--become-method"
	BecomeUser   = "--become-user"
	Become       = "-b"
)

func genAdHocCommand(r *AdHocRunner) []string {
	argv := make([]string, 0)
	if r.Bin == EmptyValue {
		r.Bin = DefaultAnsibleBin
	}
	argv = append(argv, r.Bin)

	if r.Options.Pattern == EmptyValue {
		r.Options.Pattern = DefaultPattern
	}
	argv = append(argv, r.Options.Pattern)

	if r.Options.Inventory != EmptyValue {
		argv = append(argv, Inventory, r.Options.Inventory)
	}

	if r.Options.Forks == ZeroValue {
		r.Options.Forks = DefaultForks
	}
	argv = append(argv, Forks, strconv.Itoa(r.Options.Forks))

	if r.Options.BackgroundSeconds != ZeroValue {
		r.Options.PollInterval = DefaultPollInterval
		argv = append(argv, BackgroundSeconds, strconv.Itoa(r.Options.BackgroundSeconds),
			PollInterval, strconv.Itoa(r.Options.PollInterval))
	}

	if r.Options.ModulePath != EmptyValue {
		argv = append(argv, ModulePath, r.Options.ModulePath)
	}

	if r.Options.ModuleName != EmptyValue {
		argv = append(argv, ModuleName, r.Options.ModuleName)
	}

	if r.Options.ModuleArgs != EmptyValue {
		moduleArgs := genModuleArgsOptions(r.Options.ModuleArgs)
		argv = append(argv, ModuleArgs, moduleArgs)
	}

	if r.Options.Check {
		argv = append(argv, Check)
	}

	if r.Options.ExtraVars != nil {
		extraVars := genExtraVarsOptions(r.Options.ExtraVars)
		argv = append(argv, ExtraVars, extraVars)
	}

	argv = genConnectionOptions(argv, r.ConnectionOptions)
	argv = genPrivilegeOptions(argv, r.PrivilegeOptions)

	return argv
}

func genPlaybookCommand(r *PlaybookRunner) []string {
	argv := make([]string, 0)
	if r.Bin == EmptyValue {
		r.Bin = DefaultAnsiblePlaybookBin
	}
	argv = append(argv, r.Bin)

	if r.Options.Inventory != EmptyValue {
		argv = append(argv, Inventory, r.Options.Inventory)
	}

	if r.Options.Playbook != EmptyValue {
		argv = append(argv, r.Options.Playbook)
	}

	if r.Options.Forks == ZeroValue {
		r.Options.Forks = DefaultForks
	}
	argv = append(argv, Forks, strconv.Itoa(r.Options.Forks))

	if r.Options.ModulePath != EmptyValue {
		argv = append(argv, ModulePath, r.Options.ModulePath)
	}

	if r.Options.Check {
		argv = append(argv, Check)
	}

	if r.Options.FlushCache {
		argv = append(argv, FlushCache)
	}

	if r.Options.Tags != EmptyValue {
		argv = append(argv, Tags, r.Options.Tags)
	}

	if r.Options.ExtraVars != nil {
		extraVars := genExtraVarsOptions(r.Options.ExtraVars)
		argv = append(argv, ExtraVars, extraVars)
	}

	argv = genConnectionOptions(argv, r.ConnectionOptions)
	argv = genPrivilegeOptions(argv, r.PrivilegeOptions)

	return argv
}

func genConnectionOptions(argv []string, o *ConnectionOptions) []string {
	if o.PrivateKey != EmptyValue {
		argv = append(argv, PrivateKey, o.PrivateKey)
	}

	if o.Timeout != ZeroValue {
		argv = append(argv, Timeout, strconv.Itoa(o.Timeout))
	}

	if o.Connection != EmptyValue {
		argv = append(argv, Connection, o.Connection)
	}

	if o.RemoteUser != EmptyValue {
		argv = append(argv, RemoteUser, o.RemoteUser)
	}

	return argv
}

func genPrivilegeOptions(argv []string, o *PrivilegeOptions) []string {
	if !o.Become {
		return argv
	}
	argv = append(argv, Become)

	if o.BecomeMethod != EmptyValue {
		argv = append(argv, BecomeMethod, o.BecomeMethod)
	}

	if o.BecomeUser != EmptyValue {
		argv = append(argv, BecomeUser, o.BecomeUser)
	}

	return argv
}

func genModuleArgsOptions(args interface{}) string {
	args = convertModuleArgs(args)
	switch args.(type) {
	case string:
		return args.(string)
	case map[string]interface{}:
		argsMap := args.(map[string]interface{})

		var _args []string

		//module has free_form option
		if v, ok := argsMap[FreeFormValue]; ok {
			_args = append(_args, v.(string))
			delete(argsMap, FreeFormValue)
		}
		//convert bool to string, true: "yes", false:"no".
		for k, v := range argsMap {
			switch v.(type) {
			case bool:
				argsMap[k] = boolString(v.(bool))
			}
		}

		for k, v := range argsMap {
			_args = append(_args, fmt.Sprintf("%s=%s", k, v))
		}

		return strings.Join(_args, " ")
	default:
		return ""
	}
}

func genExtraVarsOptions(args map[string]string) string {
	extraVars, _ := json.Marshal(args)
	return string(extraVars)
}

func convertModuleArgs(args interface{}) interface{} {
	ref := reflect.TypeOf(args)
	switch ref.Kind() {
	case reflect.String:
		return args
	case reflect.Map:
		return args
	case reflect.Struct:
		return moduleArgsMap(args)
	case reflect.Ptr:
		return moduleArgsMap(args)
	default:
		return ""
	}
}

func moduleArgsMap(s interface{}) map[string]interface{} {
	tag := "map"
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	var data = make(map[string]interface{})
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		tagName := t.Field(i).Tag.Get(tag)
		if tagName == "" || tagName == "-" {
			tagName = strings.ToLower(t.Field(i).Name)
		}
		//if v.Field(i).String() != "" {
		//	data[tagName] = v.Field(i).String()
		//}
		switch v.Field(i).Kind() {
		case reflect.String:
			if v.Field(i).String() != "" {
				data[tagName] = v.Field(i).String()
			}
		case reflect.Bool:
			data[tagName] = v.Field(i).Bool()
		}

	}

	return data
}

//convert bool to string, true: "yes", false:"no".
func boolString(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
