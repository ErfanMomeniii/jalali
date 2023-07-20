![image description](./assets/photos/logo.png)
<p align="center">
<a href="https://pkg.go.dev/github.com/erfanmomeniii/jalali?tab=doc"target="_blank">
    <img src="https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
</a>

<img src="https://img.shields.io/badge/license-MIT-magenta?style=for-the-badge&logo=none" alt="license" />
<img src="https://img.shields.io/badge/Version-v1.0.0-red?style=for-the-badge&logo=none" alt="version" />
</p>

# jalali

The Jalali calendar is a solar calendar that was used in Persia. Variants of this calendar are still used today in Iran, Afghanistan and some other countries. You can find more information on [Wikipedia](https://en.wikipedia.org/wiki/Jalali_calendar) or check out the [Calendar Converter](https://www.fourmilab.ch/documents/calendar).


# Documentation

## Install

```bash
go get github.com/erfanmomeniii/jalali
```   

Next, include it in your application:

```bash
import "github.com/erfanmomeniii/jalali"
``` 

## Quick Start

### Usage
You can see full documentation [Here](https://pkg.go.dev/github.com/erfanmomeniii/jalali).

#### type jalaliDateTime
jalaliDateTime is an instantiation of jalali Date and Time.

####func New
```go
func New(year int, month int, day int, hour int, minute int, second int) *jalaliDateTime
```
New creates a new instance of a jalaliDateTime.

####func (j *jalaliDateTime) SetLocale
```go
func (j *jalaliDateTime) SetLocale(lang Lang)
```
SetLocale sets the locale of the jalaliDateTime.

####func Now
```go
func Now() *jalaliDateTime
```
Now returns the current jalaliDateTime.

####func ConvertGregorianToJalali
```go
func ConvertGregorianToJalali(t time.Time) *jalaliDateTime
```
ConvertGregorianToJalali returns converted date and time on t from gregorian to jalali.

####func ConvertJalaliToGregorian
```go
func ConvertJalaliToGregorian(j *jalaliDateTime) time.Time
```
ConvertJalaliToGregorian returns converted date and time on j from jalali to gregorian.

####func ToGregorian
```go
func ToGregorian(gregorianSeconds int64) time.Time
```
ToGregorian returns time.Time obtained from given seconds.

####func ToJalali
```go
func ToJalali(jalaliSeconds int64) *jalaliDateTime
```
ToJalali returns jalaliDateTime obtained from given seconds.

####func IsLeapYear
```go
func IsLeapYear(year int) int
```
IsLeapYear determines the year is leap or not in jalali date.

####func Yesterday
```go
func Yesterday() *jalaliDateTime
```
Yesterday returns datetime of yesterday.

####func Tomorrow
```go
func Tomorrow() *jalaliDateTime
```
Tomorrow returns datetime of tomorrow.

####func (j *jalaliDateTime) Add
```go
func (j *jalaliDateTime) Add(t jalaliDateTime) *jalaliDateTime
```
Add returns the time j+t.

####func (j *jalaliDateTime) AddDate
```go
func (j *jalaliDateTime) AddDate(year int, month int, day int) *jalaliDateTime
```
AddDate returns the time corresponding to adding the given number of years, months, and days to j

####func (j *jalaliDateTime) Yesterday
```go
func (j *jalaliDateTime) Yesterday() *jalaliDateTime
```
Yesterday returns datetime of yesterday on a given day.

####func (j *jalaliDateTime) Tomorrow
```go
func (j *jalaliDateTime) Tomorrow() *jalaliDateTime
```
Tomorrow returns datetime of tomorrow on a given day.

####func (j *jalaliDateTime) Year
```go
func (j *jalaliDateTime) Year() int
```
Year returns the year in which j occurs.

####func (j *jalaliDateTime) Month
```go
func (j *jalaliDateTime) Month() int
```
Month returns the month of the year specified by j.

####func (j *jalaliDateTime) Day
```go
func (j *jalaliDateTime) Day() int
```
Day returns the day of the month specified by j.

####func (j *jalaliDateTime) Hour
```go
func (j *jalaliDateTime) Hour() int
```
Hour returns the hour within the day specified by j, in the range [0, 23].

####func (j *jalaliDateTime) Minute
```go
func (j *jalaliDateTime) Minute() int
```
Minute returns the minute offset within the hour specified by j, in the range [0, 59].

####func (j *jalaliDateTime) Second
```go
func (j *jalaliDateTime) Second() int
```
Second returns the second offset within the minute specified by j, in the range [0, 59].

####func (j *jalaliDateTime) TimeStamp
```go
func (j *jalaliDateTime) TimeStamp() int64
```
TimeStamp returns the timestamp of the jalaliDateTime.

####func (j *jalaliDateTime) DayOfYear
```go
func (j *jalaliDateTime) DayOfYear() int
```
DayOfYear returns the day of the year for the jalaliDateTime.

####func (j *jalaliDateTime) DayOfMonth
```go
func (j *jalaliDateTime) DayOfMonth() int
```
DayOfMonth returns the day the month for the jalaliDateTime.

####func (j *jalaliDateTime) DayOfWeek
```go
func (j *jalaliDateTime) DayOfWeek() int
```
DayOfWeek returns the day the week for the jalaliDateTime.

####func (j *jalaliDateTime) WeekToString
```go
func (j *jalaliDateTime) WeekToString() string
```
WeekToString returns the localized string representation of the day of the week for the jalaliDateTime.

####func (j *jalaliDateTime) MonthToString
```go
func (j *jalaliDateTime) MonthToString() string
```
MonthToString returns the localized string representation of the month for the jalaliDateTime.

####func (j *jalaliDateTime) Time
```go
func (j *jalaliDateTime) Time() time.Time
```
Time returns the time.Time equivalent of the jalaliDateTime.

####func (j *jalaliDateTime) String
```go
func (j *jalaliDateTime) String() string
```
String returns the string representation of the jalaliDateTime.
