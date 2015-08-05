package core

import (
	"strconv"
	"testing"
	"time"
)

func TestCheckCrontabDate(t *testing.T) {
	dt := 1
	tm := time.Now()
	_, month, day := tm.Date()
	dst_times := [5]int{
		tm.Minute(),
		tm.Hour(),
		day,
		int(month),
		int(tm.Weekday())}
	t_times := dst_times[dt:]
	src_times := [5]string{"0", strconv.Itoa(tm.Hour()), "*", "*", "*"}
	s_times := src_times[dt:]
	// days := DaysInMonth(int(month), year)
	if !checkInCrontabTime(s_times, t_times, tm) {
		t.Errorf("check failed where %d", dt)
	}
}

func TestResolveCrontabTimeAtom(t *testing.T) {
	if ok, err := resolveCrontabTimeAtom(TDay, 2, "1-7", time.Now()); !ok {
		t.Error("resolve faild", err)
	}
	if ok, err := resolveCrontabTimeAtom(TDay, 28, "23-7", time.Now()); !ok {
		t.Error("resolve faild", err)
	}
	if ok, err := resolveCrontabTimeAtom(TDay, 28, "23-7", time.Now()); !ok {
		t.Error("resolve faild", err)
	}
	if ok, err := resolveCrontabTimeAtom(TDay, 25, "23,25,7", time.Now()); !ok {
		t.Error("resolve faild", err)
	}
}
