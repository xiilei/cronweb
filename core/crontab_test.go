package core

import (
	"testing"
	"time"
)

func TestCheckCrontabDate(t *testing.T) {
	dt := 0
	now := time.Date(2015, time.August, 3, 5, 2, 0, 0, time.Local)
	dst_times := [5]int{2, 5, 31, 8, 5}
	t_times := dst_times[dt:]
	src_times := [5]string{"1-2/1", "2,3,5", "30-5", "5,8,1", "*"}
	s_times := src_times[dt:]
	if ok, err := checkInCrontabTime(s_times, t_times, now); !ok {
		t.Errorf("check failed with %d,error:%v", dt, err)
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
	if ok, err := resolveCrontabTimeAtom(TMonth, 8, "aug,25,7", time.Now()); !ok {
		t.Error("resolve faild", err)
	}
}
