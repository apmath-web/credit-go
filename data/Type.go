package data

type Type string

const (
	Regular Type = "regular"
	Early   Type = "early"
)

func (c Type) Type2Str() string {
	if c == Regular {
		return "regular"
	}
	if c == Early {
		return "early"
	}
	return ""
}
