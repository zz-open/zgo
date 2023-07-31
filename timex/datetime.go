package timex

import "time"

// AddMinute add or sub minute to the time.
func AddMinute(t time.Time, minute int64) time.Time {
	return t.Add(time.Minute * time.Duration(minute))
}

// AddHour add or sub hour to the time.
func AddHour(t time.Time, hour int64) time.Time {
	return t.Add(time.Hour * time.Duration(hour))
}

// AddDay add or sub day to the time.
func AddDay(t time.Time, day int64) time.Time {
	return t.Add(24 * time.Hour * time.Duration(day))
}

// AddYear add or sub year to the time.
func AddYear(t time.Time, year int64) time.Time {
	return t.Add(365 * 24 * time.Hour * time.Duration(year))
}

// BetweenSeconds returns the number of seconds between two times.
func BetweenSeconds(t1 time.Time, t2 time.Time) int64 {
	index := t2.Unix() - t1.Unix()
	return index
}

// IsWeekend checks if passed time is weekend or not.
func IsWeekend(t time.Time) bool {
	return time.Saturday == t.Weekday() || time.Sunday == t.Weekday()
}
