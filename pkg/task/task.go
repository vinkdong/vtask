package task

import (
	"os"
	"io/ioutil"
	"github.com/vinkdong/gox/log"
	"github.com/ghodss/yaml"
)

func GetTasks(config string) *TaskD {
	if config == "" {
		log.Error("need config file")
		os.Exit(127)
	}
	r, err := ioutil.ReadFile(config)
	if err != nil {
		log.Error(err)
	}
	tasks := &TaskD{}
	yaml.Unmarshal(r, tasks)
	return tasks
}

