package main

import (
	"os/exec"
	"flag"
	"io/ioutil"
	"os"
	"github.com/vinkdong/gox/log"
	"gopkg.in/yaml.v2"
	"github.com/vinkdong/vtask/common"
)

var(
	config = flag.String("conf","config.yml","config file of taskd")
)

func main() {

	flag.Parse()
	tasks := getTasks()
	for _, task := range tasks.Tasks {
		go func() {
			log.Infof("execute %s",task.Name)
			cmd := exec.Command(task.Cmd, task.Args...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Run()
		}()
	}
	general := tasks.General
	log.Infof("execute %s",general.Name)
	cmd := exec.Command(general.Cmd, general.Args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getTasks() *common.TaskD {
	if *config == "" {
		log.Error("need config file")
		os.Exit(127)
	}
	r, err := ioutil.ReadFile(*config)
	if err != nil {
		log.Error(err)
	}
	tasks := &common.TaskD{}
	yaml.Unmarshal(r, tasks)
	return tasks
}
