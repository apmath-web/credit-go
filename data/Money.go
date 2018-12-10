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

func (m Money) Mon2Int64() int64 {
	i := int64(m)
	return i
}

func (m Money) Mon2Float64() float64 {
	i := float64(m)
	return i
}

/*
make this type more useful, not only int64
*/
