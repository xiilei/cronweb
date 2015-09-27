package cron

import (
	// "github.com/boltdb/bolt"
	"log"
	"strings"
	"time"
)

type TaskStore struct {
	tasks []Task
}

//NewTaskStore management tasks read from the crontab
func NewTaskStore(c int) (ts *TaskStore) {
	tasks, err := fromCrontab(c)
	if err != nil {
		log.Panicf("reading crontab error,%s", err.Error())
	}
	return &TaskStore{
		tasks: tasks,
	}
}

//Raw return raw crontab tasks
func (ts *TaskStore) Raw() string {
	tasks_desc := make([]string, len(ts.tasks))
	for i, task := range ts.tasks {
		tasks_desc[i] = task.Raw()
	}
	return strings.Join(tasks_desc, "\n")
}

//Tasks return tasks by time
func (ts *TaskStore) Tasks(dt TDate, tm time.Time) []Task {
	tasks := make([]Task, 0, 1)
	_, month, day := tm.Date()
	dst_times := [5]int{
		tm.Minute(),
		tm.Hour(),
		day,
		int(month),
		int(tm.Weekday())}
	t_times := dst_times[dt:]
	for _, task := range ts.tasks {
		s_times := task.times[dt:]
		ok, err := checkInCrontabTime(s_times, t_times, tm)
		if err != nil {
			log.Panicf("check cron time failed:%s", err.Error())
		}
		if ok {
			tasks = append(tasks, task)
		}
	}
	return tasks
}
