package valueObjects

import "github.com/apmath-web/credit-go/data"

type Payment struct {
	Currency           data.Currency
	Payment            data.Money
	Type               data.Type
	Date               data.Date
	State              data.State
	Percent            data.Money
	Body               data.Money
	FullEarlyRePayment data.Money
	RemainCreditBody   data.Money
}

func GenPayment(payment data.Money, type_ data.Type, currency data.Currency, date data.Date, state data.State,
	percent data.Money, body data.Money, remainCreditBody data.Money, fullEarlyRePayment data.Money) PaymentInterface {
	p := new(Payment)
	p.Payment = payment
	p.Type = type_
	p.Currency = currency
	p.Date = date
	p.State = state
	p.Percent = percent
	p.Body = body
	p.RemainCreditBody = remainCreditBody
	p.FullEarlyRePayment = fullEarlyRePayment
	return p
}

func GenRequestPayment(payment data.Money, type_ data.Type, currency data.Currency, date data.Date) PaymentInterface {
	p := new(Payment)
	p.Payment = payment
	p.Type = type_
	p.Currency = currency
	p.Date = date
	return p
}

func (p *Payment) GetPayment() data.Money {
	return p.Payment
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
func (p *Payment) GetPercent() data.Money {
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
