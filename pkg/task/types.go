package task

import (
	"os"
	"github.com/vinkdong/gox/log"
	"os/exec"
	"time"
)

type TaskD struct {
	Tasks   []Task
	General Task
}

func (t *TaskD) Start() error {

	for _, task := range t.Tasks {
		go task.Run()
	}

	err := t.General.Run()
	if err != nil {
		log.Error(err)
	}
	return nil
}

type Task struct {
	Name string
	Cmd  string
	Args []string
	ExecCmd  *exec.Cmd
}

func (t *Task) Run() error {
	log.Infof("execute %s", t.Name)
	cmd := exec.Command(t.Cmd, t.Args...)
	t.ExecCmd = cmd

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Start(); err != nil {
		return err
	}
	t.listening()
	return cmd.Start()
}

func (t *Task) listening() {
	for{
		select {
		case <-time.Tick(time.Second * 5):

		}
	}
}
