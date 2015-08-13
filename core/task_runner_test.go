package core

import (
	"os"
	"testing"
)

func TestRunTask(t *testing.T) {
	times := TTimes{"*", "*", "*", "*", "*"}
	abs, _ := os.Getwd()
	task := NewTask("python3", "", times, []string{abs + "/../ever.py"})
	_, err := RunTask(task)
	if err != nil {
		t.Errorf("run error:%s\n", err.Error())
	}
}
