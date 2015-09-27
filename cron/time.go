package cron

import (
	"time"
)

const (
	TMinute TDate = iota
	THour
	TDay
	TMonth
	TWeek
)

//the crontab times
type TTimes [5]string

//the task time flag
type TDate int

var Months = [...]string{
	"jan",
	"feb",
	"mar",
	"apr",
	"may",
	"jun",
	"jul",
	"aug",
	"sep",
	"oct",
	"nov",
	"dec",
}

var td_limits = map[TDate]int{
	TMinute: 60,
	THour:   60,
	TMonth:  12,
	TWeek:   7,
}

var month_days = [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

//DaysInMonth return total days of the month in year
func DaysInMonth(month, year int) int {
	month = month - 1
	if month == 1 && ((year%4 == 0 && year%100 != 0) || year%400 == 0) {
		return 29
	}
	return month_days[month]
}

func TDateMax(dt TDate, t time.Time) int {
	if dt == TDay {
		year, month, _ := t.Date()
		return DaysInMonth(int(month), year)
	}
	return td_limits[dt]
}
