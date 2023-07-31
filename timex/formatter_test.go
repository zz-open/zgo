package timex

import (
	"testing"
	"time"

	"github.com/zzopen/go-saber/internal"
)

func TestFormatTimeToStr(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFormatTimeToStr")

	type Case = struct {
		time      time.Time
		formatter string
		expected  string
	}

	datetime, _ := time.Parse(time.DateTime, "2023-07-31 16:58:30")
	startDateTime, _ := time.Parse(time.DateTime, "1970-01-01 00:00:00")
	zeroDateTime, _ := time.Parse(time.DateTime, "0000-00-00 00:00:00")
	zeroOtherDateTime, _ := time.Parse(time.DateTime, "0001-01-01 00:00:00")

	times := []time.Time{datetime, startDateTime, zeroDateTime, zeroOtherDateTime}
	var cases []Case
	for _, v := range times {
		if v == datetime {
			cases = append(cases,
				Case{time: v, formatter: "YYYY-MM-DD HH:mm:ss", expected: "2023-07-31 16:58:30"},
				Case{time: v, formatter: "YYYY-MM-DD", expected: "2023-07-31"},
				Case{time: v, formatter: "HH:mm:ss", expected: "16:58:30"},
			)
		}

		if v == startDateTime {
			cases = append(cases,
				Case{time: v, formatter: "YYYY-MM-DD HH:mm:ss", expected: "1970-01-01 00:00:00"},
				Case{time: v, formatter: "YYYY-MM-DD", expected: "1970-01-01"},
				Case{time: v, formatter: "HH:mm:ss", expected: "00:00:00"},
			)
		}

		if v == zeroOtherDateTime || v == zeroDateTime {
			cases = append(cases,
				Case{time: v, formatter: "YYYY-MM-DD HH:mm:ss", expected: ""},
				Case{time: v, formatter: "YYYY-MM-DD", expected: ""},
				Case{time: v, formatter: "HH:mm:ss", expected: ""},
			)
		}
	}

	for i := 0; i < len(cases); i++ {
		actual := FormatTimeToStr(cases[i].time, cases[i].formatter)
		assert.Equal(cases[i].expected, actual)
	}
}

func TestNowDateTimeStr(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestNowDateTimeStr")

	actual := NowDateTimeStr()
	assert.Equal(time.Now().Format(time.DateTime), actual)
}

func TestNowDateStr(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestNowDateStr")

	actual := NowDateStr()
	assert.Equal(time.Now().Format(time.DateOnly), actual)
}

func TestNowTimeStr(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestNowTimeStr")

	actual := NowTimeStr()
	assert.Equal(time.Now().Format(time.TimeOnly), actual)
}

func TestFormatStrToTime(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFormatTimeToStr")

	type Case = struct {
		time      string
		formatter string
		layout    string
	}

	times := [][]string{
		{"2021-01-02 16:04:08", "2021-01-02 16:04", "2021-01-02 16", "2021-01-02", "16:04:08", "16:04", "16"},
		{"0000-00-00 00:00:00", "0000-00-00 00:00", "0000-00-00 00", "0000-00-00", "00:00:00", "00:00", "00"},
		{"0001-01-01 00:00:00", "0001-01-01 00:00", "0001-01-01 00", "0001-01-01", "00:00:00", "00:00", "00"},
	}

	var cases []Case
	for _, v := range times {
		cases = append(cases,
			Case{time: v[0], formatter: "YYYY-MM-DD HH:mm:ss", layout: "2006-01-02 15:04:05"},
			Case{time: v[1], formatter: "YYYY-MM-DD HH:mm", layout: "2006-01-02 15:04"},
			Case{time: v[2], formatter: "YYYY-MM-DD HH", layout: "2006-01-02 15"},
			Case{time: v[3], formatter: "YYYY-MM-DD", layout: "2006-01-02"},
			Case{time: v[4], formatter: "HH:mm:ss", layout: "15:04:05"},
			Case{time: v[5], formatter: "HH:mm", layout: "15:04"},
			Case{time: v[6], formatter: "HH", layout: "15"},
		)
	}

	for i := 0; i < len(cases); i++ {
		actual, err := FormatStrToTime(cases[i].time, cases[i].formatter)
		if err != nil && !actual.IsZero() {
			t.Fatal(err, i, cases[i].formatter)
		}
		expected, _ := time.Parse(cases[i].layout, cases[i].time)
		assert.Equal(expected, actual)
	}
}
