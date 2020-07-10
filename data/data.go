package data

import "fmt"

type Time struct {
	Hour     int
	Minute   int
	Second   int
	Duration int
}

func NewTime(hour, minute, second, duration int) *Time {
	return &Time{Hour: hour, Minute: minute, Second: second, Duration: duration}
}

func (t *Time) convertString(val int) string {

	if val < 10 {
		return fmt.Sprintf("0%d", val)
	}

	return fmt.Sprintf("%d", val)
}

func (t *Time) String() string {

	hourStr := t.convertString(t.Hour)
	minuteStr := t.convertString(t.Minute)
	secondStr := t.convertString(t.Second)

	return fmt.Sprintf("%s:%s:%s", hourStr, minuteStr, secondStr)
}
