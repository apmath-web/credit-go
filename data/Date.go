package data

import "time"

type Date time.Time

func (d Date) Date2Str() string {
	t := time.Time(d)
	return t.Format("2006-01-02")
}

func Str2Date(s string) Date {
	date, err := time.Parse("2006-01-02", s)
	if err == nil {
		return Date(date)
	}
	return Str2Date(Date(time.Now()).Date2Str())
}
