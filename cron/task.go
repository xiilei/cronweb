package cron

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

type Task struct {
	//the task short description.
	title string
	//the executable's name.
	name string
	//the crontab times ,like * * * * *
	times TTimes
	//
	args []string
	//id
	id string
}

func NewTask(name, title string, times TTimes, args []string) *Task {
	return &Task{
		name:  name,
		times: times,
		args:  args,
		title: title,
	}
}

//the stored id of task
//@todo add machine unique id
func (t *Task) Id() string {
	if t.id != "" {
		return t.id
	}
	h := md5.New()
	io.WriteString(h, t.name)
	io.WriteString(h, strings.Join(t.args, ""))
	t.id = fmt.Sprintf("%x", h.Sum(nil))
	return t.id
}

func (t *Task) SetTitle(title string) {
	t.title = title
}

//return the character for crontab
func (t *Task) Raw() string {
	return fmt.Sprintf("%s %s %s",
		strings.Join(t.times[:], " "), t.name, strings.Join(t.args, " "))
}

//resolve crontab's task
func ResolveTask(desc string) *Task {
	var name string
	var times TTimes
	args := make([]string, 0, 1)
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
	return NewTask(name, "", times, args)
}
