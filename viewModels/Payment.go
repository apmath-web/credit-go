package viewModels

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/valueObjects"
	"time"
)

type Payment struct {
	viewModel
	Type            string `json:"type"`
	State           string `json:"state"`
	Date            string `json:"date"`
	AmountOfPayment int64  `json:"payment"`
	Currency        string `json:"currency"`
}

func (p *Payment) Fetch() interface{} { //TODO
	return 0
}

func (p *Payment) Validate() bool {
	p.validateType()
	p.validatePayment()
	p.validateCurrency()
	p.validateDate()
	if len(p.validMessages.GetMessages()) == 0 {
		return true
	}
	return false
}

func (p *Payment) validateDate() {
	if val := p.check("string", "date"); val != nil {
		p.Date = val.(string)
		if _, err := time.Parse("2006-01-02", p.Date); err != nil {
			p.validMessages.AddMessage(
				valueObjects.GenMessage("date", "Is wrong format of date."))
		}
	}
}

func (p *Payment) validatePayment() {
	if val := p.check("float64", "amount"); val != nil {
		p.AmountOfPayment = int64(val.(float64))
		if p.GetPayment() <= 100 && p.GetPayment() > 3000000000000000 {
			p.validMessages.AddMessage(
				valueObjects.GenMessage("payment", "Must be between 100 and 3000000000	000000"))
		}
	}
}

func (p *Payment) validateCurrency() {
	if val := p.check("string", "currency"); val != nil {
		p.Currency = val.(string)
		if p.GetCurrency() == "" {
			p.validMessages.AddMessage(
				valueObjects.GenMessage("currency", "Is unknown currency."))
		}
	}
}

func (p *Payment) validateType() {
	if val := p.check("string", "type"); val != nil {
		p.Type = val.(string)
		if p.GetType() == "" {
			p.validMessages.AddMessage(
				valueObjects.GenMessage("type", "Is unknown type."))
		}
	}
}

func (p *Payment) Hydrate(payment valueObjects.PaymentInterface) { //TODO
	return
}

func (p *Payment) GetPayment() data.Money {
	return data.Money(p.AmountOfPayment)
}

func (p *Payment) GetCurrency() data.Currency {
	return data.Str2Cur(p.Currency)
}

func (p *Payment) GetDate() data.Date {
	return data.Str2Date(p.Date)
}

func (p *Payment) GetType() data.Type {
	return data.Type(p.Type)
}

func (p *Payment) GetState() data.State {
	return data.State(p.State)
}
