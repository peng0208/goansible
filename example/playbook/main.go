package main

import (
	"fmt"
	"github.com/peng0208/goansible"
)

func main() {
	//ansible playbook模式用法

	runner := goansible.NewPlaybookRunner()
	runner.Options.Inventory = "127.0.0.1,"
	runner.Options.Playbook = "/root/workspace/goansible/example/playbook/play.yml"
	runner.ConnectionOptions.Connection = "local"
	fmt.Println(runner.CommandString())
	output, err := runner.Run()
	fmt.Printf("output: %v\nerror: %v\n", output.String(), err)

}
