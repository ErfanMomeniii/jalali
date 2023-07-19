package jalali_test

import (
	"testing"
	"time"

	"github.com/erfanmomeniii/jalali"

	"github.com/stretchr/testify/assert"
)

func Test_Everything(t *testing.T) {
	// check jalali instantiation and localization
	j := jalali.New(1402, 4, 28, 9, 0, 0)
	j.SetLocale(jalali.PersianLanguage)

	assert.Equal(t, j.WeekToString(), "چهارشنبه")
	assert.Equal(t, j.MonthToString(), "تیر")

	j.SetLocale(jalali.EnglishLanguage)

	assert.Equal(t, j.WeekToString(), "Wednesday")
	assert.Equal(t, j.MonthToString(), "April")

	// check convert jalali to gregorian
	g := jalali.ConvertJalaliToGregorian(j)

	assert.Equal(t, g.Year(), 2023)
	assert.Equal(t, g.Month(), time.Month(7))
	assert.Equal(t, g.Day(), 19)

	// check yesterday and tomorrow
	assert.Equal(t, j.Yesterday().Day(), int64(27))
	assert.Equal(t, j.Yesterday().Month(), int64(4))
	assert.Equal(t, j.Yesterday().Year(), int64(1402))

	assert.Equal(t, j.Tomorrow().Day(), int64(29))
	assert.Equal(t, j.Tomorrow().Month(), int64(4))
	assert.Equal(t, j.Tomorrow().Year(), int64(1402))

	// check day of week,month and year
	assert.Equal(t, j.DayOfWeek(), 5)
	assert.Equal(t, j.DayOfMonth(), 28)
	assert.Equal(t, j.DayOfYear(), 121)
}
