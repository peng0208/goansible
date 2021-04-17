package goansible

type CronModuleArgs struct {
	//create a backup of the crontab before it is modified.
	Backup bool
	//if specified, uses this file instead of an individual user's crontab.
	CronFile string `map:"cron_file"`
	//description of a crontab entry or, if env is set, the name of environment variable.
	Name string
	//minute when the job should run
	Minute string
	//hour when the job should run
	Hour string
	//day of the month the job should run
	Day string
	//month of the year the job the job should run
	Month string
	//day of the week the job should run
	Weekday string
	//special time specification nickname.(choices: annually, daily, hourly, monthly, reboot, weekly, yearly)
	SpecialTime string `map:"special_time"`
	//the specific user whose crontab should be modified.
	User string
	//choices: absent, present
	State string
	//if the job should be disabled (commented out) in the crontab. Only has effect if `state=present'.
	Disabled bool
	//if set, manages a crontab's environment variable
	Env bool
	//if specified, the environment variable will be inserted after the declaration of specified environment variable.
	Insertafter string
	//if specified, the environment variable will be inserted before the declaration of specified environment variable.
	Insertbefore string
	//the command to execute or, if env is set, the value of environment variable.
	Job string
}

func NewCronModule() *ModuleRunner {
	return &ModuleRunner{
		"cron",
		NewAdHocRunner(),
		&CronModuleArgs{},
		&ModuleAdHocOptions{},
		&ConnectionOptions{},
		&PrivilegeOptions{},
	}
}
