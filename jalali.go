package jalali

import (
	"fmt"
	"time"
)

const DifferenceInSecond uint64 = 19604073600
const DayInSecond uint64 = 60 * 60 * 24
const YearInSecond uint64 = DayInSecond * 365

type jalaliDate struct {
	Year   uint64
	Month  uint64
	Day    uint64
	Hour   uint64
	Minute uint64
	Second uint64
}

func Now() *jalaliDate {
	return convertGregorianToJalali(uint64(time.Now().UTC().Sub(time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC))))
}

func convertGregorianToJalali(gregorianSeconds uint64) *jalaliDate {
	return ToJalai(gregorianSeconds - DifferenceInSecond)
}

func ToJalai(jalaliSeconds uint64) *jalaliDate {
	j := &jalaliDate{
		Year:   1,
		Month:  1,
		Day:    1,
		Hour:   0,
		Minute: 0,
		Second: 0,
	}

	for jalaliSeconds >= 60 {
		yearDuration := YearInSecond + uint64(isLeapYear(j.Year))*DayInSecond
		if jalaliSeconds >= yearDuration {
			jalaliSeconds -= yearDuration
			j.Year++
		} else if jalaliSeconds >= DayInSecond {
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
		} else {
			jalaliSeconds -= 60
			j.Minute++
			if j.Minute >= 60 {
				j.Minute -= 60
				j.Hour++
			}

			if j.Hour >= 24 {
				j.Hour -= 24
				j.Day++
			}

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

	j.Second = jalaliSeconds

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
	if year%4 == 0 {
		return 1
	}

	return 0
}

func (j *jalaliDate) ToString() string {
	return fmt.Sprintf("%d - %d - %d  %d : %d : %d", j.Year, j.Month, j.Day, j.Hour, j.Minute, j.Second)
}
