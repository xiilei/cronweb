package core

import (
	"strconv"
	"testing"
	"time"
)

func TestCheckCrontabDate(t *testing.T) {
	dt := 1
	tm := time.Now()
	year, month, day := tm.Date()
	dst_times := [5]int{
		tm.Minute(),
		tm.Hour(),
		day,
		int(month),
		int(tm.Weekday())}
	t_times := dst_times[dt:]
	src_times := [5]string{"0", strconv.Itoa(tm.Hour()), "*", "*", "*"}
	s_times := src_times[dt:]
	days := DaysInMonth(int(month), year)
	if !checkInCrontabTime(s_times, t_times, days) {
		t.Errorf("check failed where %d", dt)
	}
}
