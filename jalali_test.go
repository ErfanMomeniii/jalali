package jalali_test

import (
	"fmt"
	"github.com/erfanmomeniii/jalali"
	"testing"
	"time"
)

func Test_Now(t *testing.T) {
	fmt.Println(jalali.Now())
}

func Test_Seconds(t *testing.T) {
	fmt.Println(jalali.Now())
	s := jalali.Seconds(time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC))
	fmt.Println(s)

	s = jalali.Seconds(time.Date(1, 1, 1, 0, 0, 2, 0, time.UTC))
	fmt.Println(s)
}
