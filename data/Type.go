package data

type Type string

const (
	Regular Type = "regular"
	Early   Type = "early"
	Next    Type = "next"
)

func (c Type) Type2Str() string {
	if c == Regular {
		return "regular"
	}
	if c == Early {
		return "early"
	}
	if c == Next {
		return "next"
	}
	return ""
}

func Str2Type(s string) Type {
	if s == "regular" {
		return Regular
	}
	if s == "early" {
		return Early
	}
	if s == "next" {
		return Next
	}
	return ""
}
