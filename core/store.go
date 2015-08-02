package core

import (
	// "github.com/boltdb/bolt"
	"bufio"
	"log"
	"os/exec"
	"strings"
)

const (
	TSecond = iota
	TMinute
	THour
	TDay
	TMonth
	TWeek
)

type TaskStore struct {
	tasks []Task
}

//NewTaskStore reading tasks from crontab
func NewTaskStore(c int) (ts *TaskStore) {
	ts = &TaskStore{
		tasks: make([]Task, 0, c),
	}
	err := ts.fromCrontab()
	if err != nil {
		log.Panicf("reading crontab error,%s", err.Error())
	}
	return
}

//reading and create tasks from linux crontab
//crontab -l command
func (ts *TaskStore) fromCrontab() (err error) {
	cmd := exec.Command("crontab", "-l")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		s := scanner.Text()
		//ignore comments
		if strings.HasPrefix(s, "#") {
			continue
		}
		task := ParseTask(s)
		ts.tasks = append(ts.tasks, *task)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return
}

//returns raw crontab tasks
func (ts *TaskStore) String() string {
	task_desc := make([]string, len(ts.tasks))
	for i, task := range ts.tasks {
		//Why task.String() work here ? It not pointer of the Task
		task_desc[i] = task.String()
	}
	return strings.Join(task_desc, "\n")
}

// func (ts *TaskStore) Tasks(dt int) []Task {

// }
