package cron

import (
	"os/exec"
	"syscall"
	"time"
)

type TaskStat struct {
	Memory int64
	Cost   time.Duration
	// Output io.Reader
	// Errors []byte
}

func newTaskStat() *TaskStat {
	return &TaskStat{}
}

func RunTask(task *Task) (stat *TaskStat, err error) {
	cmd := exec.Command(task.name, task.args...)
	var usage *syscall.Rusage
	start := time.Now()
	// _, err = cmd.StdoutPipe()
	if err = cmd.Start(); err != nil {
		return nil, err
	}
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}
	stat = newTaskStat()
	stat.Cost = time.Now().Sub(start)
	usage = cmd.ProcessState.SysUsage().(*syscall.Rusage)
	stat.Memory = usage.Maxrss
	return stat, err
}
