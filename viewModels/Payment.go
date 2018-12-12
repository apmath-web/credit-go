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

	percent          int32
	body             int64
	remainCreditBody int64
	fullEarlyPayment int64
}

func (p *Payment) Fetch() interface{} {
	jsonData := make(map[string]interface{})
	jsonData["type"] = p.Type
	jsonData["state"] = p.State
	jsonData["date"] = p.Date
	jsonData["payment"] = p.AmountOfPayment
	jsonData["percent"] = p.percent
	jsonData["body"] = p.body
	jsonData["remainCreditBody"] = p.remainCreditBody
	if p.State == data.Upcoming.State2Str() || p.Type == data.Next.Type2Str() {
		jsonData["fullEarlyRepayment"] = p.fullEarlyPayment
	}
	return jsonData
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
	if val, ok := p.JsonData["date"]; (ok && val == nil) || !ok {
		p.JsonData["date"] = data.Date(time.Now()).Date2Str()
	}
	if val := p.check("string", "date"); val != nil {
		p.Date = val.(string)
		if _, err := time.Parse("2006-01-02", p.Date); err != nil {
			p.validMessages.AddMessage(
				valueObjects.GenMessage("date", "Is wrong format of date."))
		}
	}
}

func (p *Payment) validatePayment() {
	if val := p.check("float64", "payment"); val != nil {
		p.AmountOfPayment = int64(val.(float64))
		if p.GetPayment() < 1 && p.GetPayment() > 3750000000000000 {
			p.validMessages.AddMessage(
				valueObjects.GenMessage("payment", "Must be between 1 and 3750000000000000"))
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
	if val, ok := p.JsonData["type"]; (ok && val == nil) || !ok {
		return
	}
	if val := p.check("string", "type"); val != nil {
		p.Type = val.(string)
		if p.GetType() == "" {
			p.validMessages.AddMessage(
				valueObjects.GenMessage("type", "Is unknown type."))
		}
	}
}

func (p *Payment) Hydrate(payment valueObjects.PaymentInterface) {
	p.Currency = payment.GetCurrency().Cur2Str()
	p.Type = payment.GetType().Type2Str()
	p.Date = payment.GetDate().Date2Str()
	p.State = payment.GetState().State2Str()
	p.percent = payment.GetPercent()
	p.AmountOfPayment = payment.GetPayment().Mon2Int64()
	p.body = payment.GetBody().Mon2Int64()
	p.remainCreditBody = payment.GetRemainCreditBody().Mon2Int64()
	if payment.GetState() == data.Upcoming || payment.GetType() == data.Next {
		p.fullEarlyPayment = payment.GetFullEarlyRepayment().Mon2Int64()
	}

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
