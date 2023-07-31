package timex

import (
	"fmt"
	"time"
)

const OnlyYear = "2006"
const OnlyMonth = "01"
const OnlyDay = "02"
const OnlyHour = "15"
const OnlyMinute = "04"
const OnlySecond = "05"

var timeFormat map[string]string

func init() {
	timeFormat = map[string]string{
		"YYYY-MM-DD HH:mm:ss": OnlyYear + "-" + OnlyMonth + "-" + OnlyDay + " " + OnlyHour + ":" + OnlyMinute + ":" + OnlySecond,
		"YYYY-MM-DD HH:mm":    OnlyYear + "-" + OnlyMonth + "-" + OnlyDay + " " + OnlyHour + ":" + OnlyMinute,
		"YYYY-MM-DD HH":       OnlyYear + "-" + OnlyMonth + "-" + OnlyDay + " " + OnlyHour,
		"YYYY-MM-DD":          OnlyYear + "-" + OnlyMonth + "-" + OnlyDay,
		"HH:mm:ss":            OnlyHour + ":" + OnlyMinute + ":" + OnlySecond,
		"HH:mm":               OnlyHour + ":" + OnlyMinute,
		"HH":                  OnlyHour,
	}
}

// FormatTimeToStr time.Time to str
func FormatTimeToStr(t time.Time, format string) string {
	if t.IsZero() {
		return ""
	}

	v, ok := timeFormat[format]
	if !ok {
		return ""
	}

	return t.Format(v)
}

func NowDateTimeStr() string {
	return FormatTimeToStr(time.Now(), "YYYY-MM-DD HH:mm:ss")
}

func NowDateStr() string {
	return FormatTimeToStr(time.Now(), "YYYY-MM-DD")
}

func NowTimeStr() string {
	return FormatTimeToStr(time.Now(), "HH:mm:ss")
}

// TodayStartDateTime return the start time of today, format: YYYY-MM-DD 00:00:00.
func TodayStartDateTime() string {
	return NowDateStr() + " 00:00:00"
}

// TodayEndDateTime return the end time of today, format: YYYY-MM-DD 23:59:59.
func TodayEndDateTime() string {
	return NowDateStr() + " 23:59:59"
}

// FormatStrToTime str to time.Time
func FormatStrToTime(str, format string) (time.Time, error) {
	v, ok := timeFormat[format]
	if !ok {
		return time.Time{}, fmt.Errorf("format %s not support", format)
	}

	return time.Parse(v, str)
}
