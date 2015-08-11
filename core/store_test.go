package core

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTaskStore(t *testing.T) {
	ts := NewTaskStore(1)
	ts.Raw()
}

func TestTasks(t *testing.T) {
	ts := NewTaskStore(1)
	tm := time.Now()
	fmt.Printf("Today,total tasks:%d\n", len(ts.Tasks(TDay, tm)))
}
