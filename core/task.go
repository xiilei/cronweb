package core

import (
	"fmt"
	"io"
	"strings"
)

type Task struct {
	// Name is the the executable's name.
	name string
	// the crontab times ,like * * * * *
	times []string
	//
	args []string
	//output
	output io.Writer
}

func NewTask(name string, times, args []string) *Task {
	return &Task{
		name:  name,
		times: times,
		args:  args,
	}
}

func (t *Task) Run() (err error) {
	//todo
	return nil
}

func (t *Task) Out() io.Writer {
	//todo
	return t.output
}

//return the character for crontab
func (t *Task) String() string {
	return fmt.Sprintf("%s %s %s",
		strings.Join(t.times, " "), t.name, strings.Join(t.args, " "))
}

//parse crontab's line task
func ParseTask(desc string) *Task {
	var name string
	times := make([]string, 5, 5)
	args := make([]string, 1)
	for i, s := range strings.Fields(desc) {
		if i < 5 {
			times[i] = s
			continue
		}
		if i == 5 {
			name = s
			continue
		}
		args = append(args, s)
	}
	return NewTask(name, times, args)
}