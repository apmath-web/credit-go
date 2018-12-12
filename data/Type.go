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
