package library

type Person struct {
	Name string
	Age  int
}

const (
	secondsInMinute int = 60
	minutesInHour   int = 60
	HourInADay      int = 24
)

func daysToHour(day int) int {
	return day * HourInADay
}

func DaysToMinute(day int) int {
	return day * HourInADay * minutesInHour
}
