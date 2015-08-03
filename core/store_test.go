package core

import (
	"testing"
)

func TestNewTaskStore(t *testing.T) {
	ts := NewTaskStore(1)
	ts.Raw()
}
