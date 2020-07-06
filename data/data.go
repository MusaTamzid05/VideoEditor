package data

import "fmt"

type Time struct {
	hour   int
	minute int
	second int
}

func NewTime(hour, minute, second int) *Time {
	return &Time{hour: hour, minute: minute, second: second}
}

func (t *Time) convertString(val int) string {

	if val < 10 {
		return fmt.Sprintf("0%d", val)
	}

	return fmt.Sprintf("%d", val)
}

func (t *Time) String() string {

	hourStr := t.convertString(t.hour)
	minuteStr := t.convertString(t.minute)
	secondStr := t.convertString(t.second)

	return fmt.Sprintf("%s:%s:%s", hourStr, minuteStr, secondStr)
}
