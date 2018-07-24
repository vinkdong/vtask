package common

type TaskD struct {
	Tasks []Task
	General Task
}

type Task struct {
	Name string
	Cmd  string
	Args []string
}
