package jalali

import (
	"fmt"
	"time"
)

const DayInSecond uint64 = 60 * 60 * 24
const YearInSecond uint64 = DayInSecond * 365

type jalaliDate struct {
	Year  uint64
	Month uint64
	Day   uint64
}

func Now() *jalaliDate {
	return convertGregorianToJalali(time.Now())
}

func Seconds(t time.Time) uint64 {
	seconds := uint64(0)
	y, m, d := t.Date()
	t = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
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

func convertGregorianToJalali(t time.Time) *jalaliDate {
	return ToJalali(Seconds(t) - Seconds(time.Date(622, 3, 22, 0, 0, 0, 0, time.UTC)))
}

func ToJalali(jalaliSeconds uint64) *jalaliDate {
	j := &jalaliDate{
		Year:  1,
		Month: 1,
		Day:   1,
	}

	for jalaliSeconds >= DayInSecond {
		yearDuration := YearInSecond + uint64(isLeapYear(j.Year))*DayInSecond

		if jalaliSeconds >= yearDuration {
			jalaliSeconds -= yearDuration
			j.Year++
		} else {
			jalaliSeconds -= DayInSecond
			j.Day++

			if ShouldUpdateMonth(j.Year, j.Month, j.Day) {
				j.Day = 1
				j.Month++
			}

			if j.Month >= 13 {
				j.Month = 1
				j.Year++
			}
		}
	}

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

func (j *jalaliDate) ToString() string {
	return fmt.Sprintf("%d - %d - %d", j.Year, j.Month, j.Day)
}
