package core

import (
	"os/exec"
)

type TaskStat struct {
	Memory   uint
	TimeCost uint
	Output   []byte
	Errors   []byte
}

func newTaskStat() *TaskStat {
	return &TaskStat{}
}

func RunTask(task *Task) error {
	cmd := exec.Command(task.name, task.args...)
	_, err := cmd.StdoutPipe()
	if err = cmd.Start(); err != nil {
		return err
	}
	//...
	err = cmd.Wait()
	return err
}
