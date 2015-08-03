package core

import (
	"testing"
)

func TestDaysInMonths(t *testing.T) {
	d := DaysInMonth(2, 2015)
	if d != 28 {
		t.Fatalf("wrong days of the month(%d),expect %d,got %d", 2, 28, d)
	}
	d = DaysInMonth(7, 2015)
	if d != 31 {
		t.Fatalf("wrong days of the month(%d),expect %d,got %d", 7, 31, d)
	}
	d = DaysInMonth(2, 2012)
	if d != 29 {
		t.Fatalf("wrong days of the month(%d),expect %d,got %d", 2, 29, d)
	}
}
