package data

import (
	"time"
)

type Date time.Time

func (d Date) Date2Str() string {
	t := time.Time(d)
	return t.Format("2006-01-02")
}

func (d Date) GetDay() int {
	day := time.Time(d).Day()
	return day
}

func (d Date) SetDay(day int) Date {
	return Date(time.Date(d.GetYear(), d.GetMonth(), d.GetDay(), 0, 0, 0, 0, time.UTC))
}

func (d Date) GetYear() int {
	year := time.Time(d).Year()
	return year
}

func (d Date) GetMonth() time.Month {
	month := time.Time(d).Month()
	return month
}

func (d Date) AddDate(year int, month int, day int) Date {
	return Date(time.Time(d).AddDate(year, month, day))
}

func GenDate(year int, month time.Month, day int) Date {
	return Date(time.Date(year, month, day, 1, 0, 0, 0, time.UTC))
}

func (d Date) GetNumberOfDays() float64 {
	return float64(time.Date(d.GetYear(), 1, 1, 0, 0, 0, 0, time.UTC).Sub(
		time.Date(d.GetYear()+1, 1, 1, 0, 0, 0, 0, time.UTC)).Hours()) / 24
}

func NullDate() Date {
	return Date(time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC))
}

func Str2Date(s string) Date {
	date, err := time.Parse("2006-01-02", s)
	if err == nil && Date(date).Date2Str() == s {
		return Date(date)
	}
	return Str2Date(Date(time.Now()).Date2Str())
}
