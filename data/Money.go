package data

import "strconv"

type Money int64

func Str2Mon(s string) Money {
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return Money(i)
	}
	return -1
}

func (m Money) Mon2Str() string {
	s := strconv.FormatInt(int64(m), 10)
	return s
}

/*
make this type more useful, not only int64
*/
