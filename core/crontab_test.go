package core

import (
	"testing"
	"time"
)

func TestCheckCrontabDate(t *testing.T) {
	dt := 1
	tm := time.Now()
	dst_times := [5]int{tm.Minute(), 2, tm.Day(), int(tm.Month()), int(tm.Weekday())}
	t_times := dst_times[dt:]
	src_times := []string{"0", "2", "*", "*", "*"}
	s_times := src_times[dt:]
	if !checkInCrontabTime(s_times, t_times) {
		t.Errorf("check failed where %d", dt)
	}
}
