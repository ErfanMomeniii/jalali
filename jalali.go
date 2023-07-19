package jalali

import (
	"fmt"
	"time"
)

const DayInSecond uint64 = 60 * 60 * 24
const YearInSecond uint64 = 365 * DayInSecond

type jalaliDateTime struct {
	year   uint64
	month  uint64
	day    uint64
	hour   uint64
	minute uint64
	second uint64
}

func Now() *jalaliDateTime {
	return convertGregorianToJalali(time.Now())
}

func Seconds(t time.Time) uint64 {
	seconds := uint64(0)
	s := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	for s.Before(t) {
		u := s
		s = s.AddDate(1, 0, 0)
		seconds += uint64(s.Sub(u).Seconds())
	}
	if s.Equal(t) {
		return seconds
	}

	seconds -= uint64(s.Sub(t).Seconds())
	return seconds
}

func convertGregorianToJalali(t time.Time) *jalaliDateTime {
	return ToJalali(Seconds(t) -
		Seconds(time.Date(622, 3, 22, 0, 0, 0, 0, time.UTC)) +
		3*60*60 + 30*60)
}

func ToJalali(jalaliSeconds uint64) *jalaliDateTime {
	j := &jalaliDateTime{
		year:   1,
		month:  1,
		day:    1,
		hour:   0,
		minute: 0,
		second: 0,
	}

	for jalaliSeconds >= 60 {
		yearDuration := YearInSecond + uint64(isLeapYear(j.year))*DayInSecond
		if jalaliSeconds >= yearDuration {
			jalaliSeconds -= yearDuration
			j.year++
		} else if jalaliSeconds >= DayInSecond {
			jalaliSeconds -= DayInSecond
			j.day++

			if ShouldUpdateMonth(j.year, j.month, j.day) {
				j.day = 1
				j.month++
			}

			if j.month >= 13 {
				j.month = 1
				j.year++
			}
		} else {
			jalaliSeconds -= 60
			j.minute++

			if j.minute >= 60 {
				j.minute -= 60
				j.hour++
			}

			if j.hour >= 24 {
				j.hour -= 24
				j.day++
			}

			if ShouldUpdateMonth(j.year, j.month, j.day) {
				j.day = 1
				j.month++
			}

			if j.month >= 13 {
				j.month = 1
				j.year++
			}
		}
	}

	j.second = jalaliSeconds

	return j
}
func ShouldUpdateMonth(year uint64, month uint64, day uint64) bool {
	switch month {
	case 1, 2, 3, 4, 5, 6:
		if day > 31 {
			return true
		}
		return false
	case 7, 8, 9, 10, 11:
		if day > 30 {
			return true
		}
		return false
	case 12:
		if isLeapYear(year) == 1 {
			if day > 30 {
				return true
			}
			return false
		} else {
			if day > 29 {
				return true
			}
			return false
		}
	}

	return false
}

func isLeapYear(year uint64) int {
	switch year % 128 {
	case 0, 4, 8, 12, 16, 20, 29, 33, 37, 41, 45, 49, 53, 62, 66, 70, 74, 78, 82, 86, 95, 99, 103, 107, 111, 115, 124:
		return 1
	case 24, 57, 90, 119:
		if year > 473 {
			return 1
		} else {
			return 0
		}
	case 25, 58, 91, 120:
		if year <= 473 {
			return 1
		}
		return 0
	}
	return 0
}

func (j *jalaliDateTime) Year() uint64 {
	return j.year
}

func (j *jalaliDateTime) Month() uint64 {
	return j.month
}

func (j *jalaliDateTime) Day() uint64 {
	return j.day
}

func (j *jalaliDateTime) Hour() uint64 {
	return j.hour
}

func (j *jalaliDateTime) Minute() uint64 {
	return j.minute
}

func (j *jalaliDateTime) Second() uint64 {
	return j.second
}

func (j *jalaliDateTime) String() string {
	return fmt.Sprintf("%04d-%02d-%02d %2d:%2d:%2d", j.year, j.month, j.day, j.hour, j.minute, j.second)
}
