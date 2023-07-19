package jalali_test

import (
	"testing"
	"time"

	"github.com/erfanmomeniii/jalali"

	"github.com/stretchr/testify/assert"
)

func Test_Now(t *testing.T) {
	j := jalali.New(1402, 4, 28, 9, 0, 0)
	j.SetLocale(jalali.Persian)

	assert.Equal(t, j.WeekToString(), "چهارشنبه")

	g := jalali.ConvertJalaliToGregorian(j)

	assert.Equal(t, g.Year(), 2023)
	assert.Equal(t, g.Month(), time.Month(7))
	assert.Equal(t, g.Day(), 19)
}
