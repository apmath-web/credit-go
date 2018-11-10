package data

type Currency string

const (
	RUR Currency = "RUR"
	EUR Currency = "EUR"
	USD Currency = "USD"
)

func Str2Cur(s string) Currency {
	if s == "RUR" {
		return RUR
	}
	if s == "EUR" {
		return EUR
	}
	if s == "USD" {
		return USD
	}
	return ""
}

func (c Currency) Cur2Str() string {
	if c == RUR {
		return "RUR"
	}
	if c == EUR {
		return "EUR"
	}
	if c == USD {
		return "USD"
	}
	return ""
}
