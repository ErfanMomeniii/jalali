package main

import (
	"github.com/erfanmomeniii/jalali"
	"time"
)

func main() {
	// Return datetime of today
	jalali.Now() //1402-04-29 14:12:00

	// Return datetime of yesterday
	jalali.Yesterday() // 1402-04-28 14:12:59

	// Return datetime of tomorrow
	jalali.Tomorrow() // 1402-04-30 14:13:47

	// Convert gregorian datetime to jalali
	jalali.ConvertGregorianToJalali(time.Date(
		2023, 7, 20, 10, 45, 25, 0, time.UTC)) // 1402-04-29 14:15:25

	// Convert jalali datetime to gregorian
	jalali.ConvertJalaliToGregorian(jalali.New(
		1402, 4, 29, 14, 15, 25)) // 2023-07-20 10:45:25 +0000 UTC

	// Create instance of jalali datetime
	j := jalali.New(1402, 4, 29, 14, 15, 25)

	// Set locale for localizing the output
	j.SetLocale(jalali.PersianLanguage)
	j.String() // ۱۴۰۲-۰۴-۲۹ ۱۴:۱۵:۲۵

	// Return the localized string representation of the day of the week and month of the year for j
	j.MonthToString() // تیر
	j.WeekToString()  // پنجشنبه

	// Return the day of the week, month and year for the j
	j.DayOfWeek()  // 6
	j.DayOfMonth() // 29
	j.DayOfYear()  // 122

	// Return yesterday and tomorrow of j
	j.Yesterday() // ۱۴۰۲-۰۴-۲۸ ۱۴:۱۵:۲۵
	j.Tomorrow()  // ۱۴۰۲-۰۴-۳۰ ۱۴:۱۵:۲۵

	j.SetLocale(jalali.EnglishLanguage)

	// Returns the datetime corresponding to adding the given number of years, months, and days to j
	j.AddDate(0, 0, 1) // 1402-04-30 14:15:25

	// Return year, month, day, hour, minute and second of j
	j.Year()   // 1402
	j.Month()  // 4
	j.Day()    // 29
	j.Hour()   // 14
	j.Minute() // 15
	j.Second() // 25
}
