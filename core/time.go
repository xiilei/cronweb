package core

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
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
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
