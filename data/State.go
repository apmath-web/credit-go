package data

type State string

const (
	Paid     State = "paid"
	Upcoming State = "upcoming"
	Next     State = "next"
)

func (c State) State2Str() string {
	if c == Paid {
		return "paid"
	}
	if c == Upcoming {
		return "upcoming"
	}
	if c == Next {
		return "next"
	}
	return ""
}
