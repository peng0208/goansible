package goansible

type (
	ModuleAdHocOptions struct {
		CommonOptions
		//host pattern
		Pattern string
		//async timeout, seconds
		BackgroundSeconds int
		//async poll interval, seconds
		PollInterval int
	}

	ModuleRunner struct {
		name   string
		runner *AdHocRunner
		Args   interface{}

		Options           *ModuleAdHocOptions
		ConnectionOptions *ConnectionOptions
		PrivilegeOptions  *PrivilegeOptions
	}
)

func (m *ModuleRunner) Run() (*Result, error) {
	m.copyModuleOptions()
	return m.runner.Run()
}

func (m *ModuleRunner) CommandString() string {
	m.copyModuleOptions()
	return m.runner.CommandString()
}

func (m *ModuleRunner) copyModuleOptions() {
	m.runner.Options.Inventory = m.Options.Inventory
	m.runner.Options.ModulePath = m.Options.ModulePath
	m.runner.Options.Check = m.Options.Check
	m.runner.Options.ExtraVars = m.Options.ExtraVars
	m.runner.Options.Forks = m.Options.Forks
	m.runner.Options.Pattern = m.Options.Pattern
	m.runner.Options.ModuleName = m.name
	m.runner.Options.ModuleArgs = m.Args
	m.runner.Options.BackgroundSeconds = m.Options.BackgroundSeconds
	m.runner.Options.PollInterval = m.Options.PollInterval
	m.runner.ConnectionOptions = m.ConnectionOptions
	m.runner.PrivilegeOptions = m.PrivilegeOptions
}
