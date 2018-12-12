package data

type State string

const (
	Paid     State = "paid"
	Upcoming State = "upcoming"
)

func (c State) State2Str() string {
	if c == Paid {
		return "paid"
	}
	if c == Upcoming {
		return "upcoming"
	}
	return ""
}
