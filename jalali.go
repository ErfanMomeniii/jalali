package jalali

import (
	"fmt"
	"strconv"
	"time"
)

type Lang int

const (
	EnglishLanguage Lang = iota
	PersianLanguage

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

// jalaliDateTime is an instantiation of jalali Date and Time.
type jalaliDateTime struct {
	year   int
	month  int
	day    int
	hour   int
	minute int
	second int
	locale Lang
}

// New creates a new instance of a jalaliDateTime.
func New(year int, month int, day int, hour int, minute int, second int) *jalaliDateTime {
	return &jalaliDateTime{
		year:   year,
		month:  month,
		day:    day,
		hour:   hour,
		minute: minute,
		second: second,
	}
}

// SetLocale sets the locale of the jalaliDateTime.
func (j *jalaliDateTime) SetLocale(lang Lang) {
	j.locale = lang

	return
}

// Now returns the current jalaliDateTime.
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
	year := 1

	for year < j.year {
		seconds += YearInSecond + int64(IsLeapYear(year))*DayInSecond
		year++
	}

	for i := 0; i < int(j.month)-1; i++ {
		seconds += monthDay[i] * DayInSecond
	}

	seconds += int64(j.day-1) * DayInSecond
	seconds += int64(j.hour*60*60 + j.minute*60 + j.second)

	return seconds
}

// ConvertGregorianToJalali returns converted date and time on t from gregorian to jalali.
func ConvertGregorianToJalali(t time.Time) *jalaliDateTime {
	return ToJalali(secondsInGregorian(t) -
		secondsInGregorian(time.Date(622, 3, 22, 0, 0, 0, 0, time.UTC)) +
		3*60*60 + 30*60)
}

// ConvertJalaliToGregorian returns converted date and time on j from jalali to gregorian.
func ConvertJalaliToGregorian(j *jalaliDateTime) time.Time {
	return ToGregorian(secondsInJalali(j) +
		secondsInGregorian(time.Date(622, 3, 22, 0, 0, 0, 0, time.UTC)) -
		3*60*60 - 30*60)
}

// ToGregorian returns time.Time obtained from given seconds.
func ToGregorian(gregorianSeconds int64) time.Time {
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

// ToJalali returns jalaliDateTime obtained from given seconds.
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

	j.second = int(jalaliSeconds)

	return j
}

func shouldUpdateMonth(year int, month int, day int) bool {
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

// IsLeapYear determines the year is leap or not in jalali date.
func IsLeapYear(year int) int {
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

// Yesterday returns datetime of yesterday.
func Yesterday() *jalaliDateTime {
	return ConvertGregorianToJalali(time.Now().Add(-24 * time.Hour))
}

// Tomorrow returns datetime of tomorrow.
func Tomorrow() *jalaliDateTime {
	return ConvertGregorianToJalali(time.Now().Add(24 * time.Hour))
}

// Add returns the time j+t.
func (j *jalaliDateTime) Add(t jalaliDateTime) *jalaliDateTime {
	t.year++
	t.month++
	t.day++

	newDate := ToJalali(secondsInJalali(j) + secondsInJalali(&t))
	newDate.SetLocale(j.locale)

	return newDate
}

// AddDate returns the datetime corresponding to adding the given number of years, months, and days to j.
func (j *jalaliDateTime) AddDate(year int, month int, day int) *jalaliDateTime {
	return j.Add(jalaliDateTime{year, month, day, 0, 0, 0, j.locale})
}

// Yesterday returns datetime of yesterday on a given day.
func (j *jalaliDateTime) Yesterday() *jalaliDateTime {
	newDate := &jalaliDateTime{j.year, j.month, j.day, j.hour, j.minute, j.second, j.locale}

	newDate = ConvertGregorianToJalali(ConvertJalaliToGregorian(newDate).Add(-24 * time.Hour))
	newDate.locale = j.locale

	return newDate
}

// Tomorrow returns datetime of tomorrow on a given day.
func (j *jalaliDateTime) Tomorrow() *jalaliDateTime {
	newDate := &jalaliDateTime{j.year, j.month, j.day, j.hour, j.minute, j.second, j.locale}
	return newDate.Add(jalaliDateTime{0, 0, 1, 0, 0, 0, 0})
}

// Year returns the year in which j occurs.
func (j *jalaliDateTime) Year() int {
	return j.year
}

// Month returns the month of the year specified by j.
func (j *jalaliDateTime) Month() int {
	return j.month
}

// Day returns the day of the month specified by j.
func (j *jalaliDateTime) Day() int {
	return j.day
}

// Hour returns the hour within the day specified by j, in the range [0, 23].
func (j *jalaliDateTime) Hour() int {
	return j.hour
}

// Minute returns the minute offset within the hour specified by j, in the range [0, 59].
func (j *jalaliDateTime) Minute() int {
	return j.minute
}

// Second returns the second offset within the minute specified by j, in the range [0, 59].
func (j *jalaliDateTime) Second() int {
	return j.second
}

// TimeStamp returns the timestamp of the j.
func (j *jalaliDateTime) TimeStamp() int64 {
	return secondsInJalali(j)
}

// DayOfYear returns the day of the year for the j.
func (j *jalaliDateTime) DayOfYear() int {
	today := secondsInJalali(j)
	firstDay := &jalaliDateTime{j.year, 1, 1, 0, 0, 0, 0}
	startOfYear := secondsInJalali(firstDay)
	duration := today - startOfYear

	return int(duration/DayInSecond) + 1
}

// DayOfMonth returns the day of the month for the j.
func (j *jalaliDateTime) DayOfMonth() int {
	return j.day
}

// DayOfWeek returns the day of the week for the j.
func (j *jalaliDateTime) DayOfWeek() int {
	duration := secondsInJalali(j)
	duration /= DayInSecond

	return int(duration % 7)
}

// WeekToString returns the localized string representation of the day of the week for the j.
func (j *jalaliDateTime) WeekToString() string {
	return localizeDay(j.DayOfWeek(), j.locale)
}

// MonthToString returns the localized string representation of the month for the j.
func (j *jalaliDateTime) MonthToString() string {
	return localizeMonth(j.month-1, j.locale)
}

// Time returns the time.Time equivalent of the j.
func (j *jalaliDateTime) Time() time.Time {
	return time.Date(int(j.year), time.Month(j.month), int(j.day), int(j.hour), int(j.minute), int(j.second), 0, time.Local)
}

func localizeDay(n int, locale Lang) string {
	if n <= 0 || n >= len(days[locale]) {
		n = 0
	}

	return days[locale][n]
}

func localizeMonth(n int, locale Lang) string {
	if n <= 0 || n >= len(months[locale]) {
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

// String returns the string representation of the jalaliDateTime.
func (j *jalaliDateTime) String() string {
	switch Lang(j.locale) {
	case PersianLanguage:
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
