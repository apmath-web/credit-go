package valueObjects

import "github.com/apmath-web/credit-go/data"

type Payment struct {
	Currency           data.Currency
	Payment            int
	Type               data.Type
	Date               data.Date
	State              data.State
	Percent            int32
	Body               data.Money
	FullEarlyRePayment data.Money
	RemainCreditBody   data.Money
}

func GenPayment(payment int, type_ data.Type, currency data.Currency, date data.Date) {
	p := new(Payment)
	p.Payment = payment
	p.Type = type_
	p.Currency = currency
	p.Date = date
}

func (p *Payment) GetPayment() data.Money {
	return data.Money(p.Payment)
}
func (p *Payment) GetType() data.Type {
	return p.Type
}
func (p *Payment) GetCurrency() data.Currency {
	return p.Currency
}
func (p *Payment) GetDate() data.Date {
	return p.Date
}
func (p *Payment) GetState() data.State {
	return p.State
}
func (p *Payment) GetPercent() int32 {
	return p.Percent
}
func (p *Payment) GetBody() data.Money {
	return p.Body
}
func (p *Payment) GetRemainCreditBody() data.Money {
	return p.RemainCreditBody
}
func (p *Payment) GetFullEarlyRepayment() data.Money {
	return p.FullEarlyRePayment
}
