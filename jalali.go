package jalali

import (
	"fmt"
	"strconv"
	"time"
)

type Lang int

const (
	English Lang = iota
	Persian

	DayInSecond  int64 = 60 * 60 * 24
	YearInSecond int64 = 365 * DayInSecond
)

var (
	days = [][]string{
		{"Friday", "Saturady", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday"},
		{"جمعه", "شنبه", "یکشنبه", "دوشنبه", "سه شنبه", "چهارشنبه", "پنجشنبه"},
	}

	months = [][]string{
		{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
		{"فروردین", "اردیبهشت", "خرداد", "تیر", "مرداد", "شهریور", "مهر", "آبان", "آذر", "دی", "بهمن", "اسفند"},
	}

	numbers = []string{
		"0123456789",
		"۰۱۲۳۴۵۶۷۸۹",
	}
)

type jalaliDateTime struct {
	locale Lang
	year   int64
	month  int64
	day    int64
	hour   int64
	minute int64
	second int64
}

func New(year int64, month int64, day int64, hour int64, minute int64, second int64) *jalaliDateTime {
	return &jalaliDateTime{
		year:   year,
		month:  month,
		day:    day,
		hour:   hour,
		minute: minute,
		second: second,
	}
}

func (j *jalaliDateTime) SetLocale(lang Lang) {
	j.locale = lang

	return
}

func Now() *jalaliDateTime {
	return ConvertGregorianToJalali(time.Now())
}

func secondsInGregorian(t time.Time) int64 {
	seconds := int64(0)

	s := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	for s.Before(t) {
		u := s
		s = s.AddDate(1, 0, 0)
		seconds += int64(s.Sub(u).Seconds())
	}

	if s.Equal(t) {
		return seconds
	}

	seconds -= int64(s.Sub(t).Seconds())

	return seconds
}

func secondsInJalali(j *jalaliDateTime) int64 {
	monthDay := []int64{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}
	seconds := int64(0)
	year := int64(1)

	for year < j.year {
		seconds += YearInSecond + int64(IsLeapYear(year))*DayInSecond
		year++
	}

	for i := 0; i < int(j.month)-1; i++ {
		seconds += monthDay[i] * DayInSecond
	}

	seconds += (j.day - 1) * DayInSecond
	seconds += j.hour*60*60 + j.minute*60 + j.second

	return seconds
}

func ConvertGregorianToJalali(t time.Time) *jalaliDateTime {
	return ToJalali(secondsInGregorian(t) -
		secondsInGregorian(time.Date(622, 3, 22, 0, 0, 0, 0, time.UTC)) +
		3*60*60 + 30*60)
}

func ConvertJalaliToGregorian(j *jalaliDateTime) time.Time {
	return ToGregoroian(secondsInJalali(j) +
		secondsInGregorian(time.Date(622, 3, 22, 0, 0, 0, 0, time.UTC)) -
		3*60*60 - 30*60)
}

func ToGregoroian(gregorianSeconds int64) time.Time {
	j := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

	for gregorianSeconds > 0 {
		s := j.AddDate(1, 0, 0)
		yearDuration := int64(s.Sub(j).Seconds())
		if gregorianSeconds >= yearDuration {
			j = s
			gregorianSeconds -= yearDuration
		} else {
			j = j.Add(time.Second * time.Duration(gregorianSeconds))
			gregorianSeconds = 0
		}
	}

	return j
}

func ToJalali(jalaliSeconds int64) *jalaliDateTime {
	j := &jalaliDateTime{
		year:   1,
		month:  1,
		day:    1,
		hour:   0,
		minute: 0,
		second: 0,
	}

	for jalaliSeconds >= 60 {
		yearDuration := YearInSecond + int64(IsLeapYear(j.year))*DayInSecond
		if jalaliSeconds >= yearDuration {
			jalaliSeconds -= yearDuration
			j.year++
		} else if jalaliSeconds >= DayInSecond {
			jalaliSeconds -= DayInSecond
			j.day++

			if shouldUpdateMonth(j.year, j.month, j.day) {
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

			if shouldUpdateMonth(j.year, j.month, j.day) {
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
func shouldUpdateMonth(year int64, month int64, day int64) bool {
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
		if IsLeapYear(year) == 1 {
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

func IsLeapYear(year int64) int {
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

func Yesterday() *jalaliDateTime {
	return ConvertGregorianToJalali(time.Now().Add(-24 * time.Hour))
}

func Tomorrow() *jalaliDateTime {
	return ConvertGregorianToJalali(time.Now().Add(24 * time.Hour))
}

func (j *jalaliDateTime) Add(t jalaliDateTime) *jalaliDateTime {
	newDate := ToJalali(secondsInJalali(j) + secondsInJalali(&t))
	newDate.SetLocale(j.locale)

	return newDate
}

func (j *jalaliDateTime) Yesterday() *jalaliDateTime {
	return ConvertGregorianToJalali(ConvertJalaliToGregorian(j).Add(-24 * time.Hour))
}

func (j *jalaliDateTime) Tomorrow() *jalaliDateTime {
	return j.Add(jalaliDateTime{0, 0, 0, 1, 0, 0, 0})
}

func (j *jalaliDateTime) Year() int64 {
	return j.year
}

func (j *jalaliDateTime) Month() int64 {
	return j.month
}

func (j *jalaliDateTime) Day() int64 {
	return j.day
}

func (j *jalaliDateTime) Hour() int64 {
	return j.hour
}

func (j *jalaliDateTime) Minute() int64 {
	return j.minute
}

func (j *jalaliDateTime) Second() int64 {
	return j.second
}

func (j *jalaliDateTime) TimeStamp() int64 {
	return secondsInJalali(j)
}

func (j *jalaliDateTime) DayOfYear() int64 {
	today := secondsInJalali(j)
	j.month = 1
	j.day = 1
	startOfYear := secondsInJalali(j)
	duration := today - startOfYear

	return duration / DayInSecond
}

func (j *jalaliDateTime) DayOfMonth() int64 {
	return j.day
}

func (j *jalaliDateTime) DayOfWeek() int64 {
	duration := secondsInJalali(j)
	duration /= DayInSecond

	return duration % 7
}

func (j *jalaliDateTime) WeekToString() string {
	return localizeDay(j.DayOfWeek(), j.locale)
}

func (j *jalaliDateTime) MonthToString() string {
	return localizeMonth(j.month-1, j.locale)
}

func (j *jalaliDateTime) Time() time.Time {
	return time.Date(int(j.year), time.Month(j.month), int(j.day), int(j.hour), int(j.minute), int(j.second), 0, time.Local)
}

func localizeDay(n int64, locale Lang) string {
	if n <= 0 || n >= int64(len(days[locale])) {
		n = 0
	}

	return days[locale][n]
}

func localizeMonth(n int64, locale Lang) string {
	if n <= 0 || n >= int64(len(months[locale])) {
		n = 0
	}

	return months[locale][n]
}

func localizeNumber(number string, locale Lang) string {
	answer := make([]rune, len(number))
	nums := []rune(numbers[locale])
	for i, _ := range number {
		j, _ := strconv.Atoi(string(number[i]))
		answer[i] = nums[j]
	}

	return string(answer)
}

func (j *jalaliDateTime) String() string {
	switch Lang(j.locale) {
	case Persian:
		return fmt.Sprintf("%s-%s-%s %s:%s:%s",
			localizeNumber(fmt.Sprintf("%04d", j.year), j.locale),
			localizeNumber(fmt.Sprintf("%02d", j.month), j.locale),
			localizeNumber(fmt.Sprintf("%02d", j.day), j.locale),
			localizeNumber(fmt.Sprintf("%02d", j.hour), j.locale),
			localizeNumber(fmt.Sprintf("%02d", j.minute), j.locale),
			localizeNumber(fmt.Sprintf("%02d", j.second), j.locale))
	default:
		return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", j.year, j.month, j.day, j.hour, j.minute, j.second)
	}
}
