package viewModels

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/valueObjects"
	"reflect"
	"time"
)

type Payment struct {
	validMessages   valueObjects.Validation
	Type            string `json:"type"`
	State           string `json:"state"`
	Date            string `json:"date"`
	AmountOfPayment int64  `json:"payment"`
	Currency        string `json:"currency"`
	JsonData        map[string]interface{}
}

func (p *Payment) Fill(jsonData map[string]interface{}) bool {
	p.JsonData = jsonData
	return true
}

func (p *Payment) Fetch() (interface{}, error) {
	return 0, nil
}

func (p *Payment) Validate() bool {
	p.validateType()
	p.validatePayment()
	p.validateCurrency()
	p.validateDate()
	p.validateState()
	if len(p.validMessages.GetMessages()) == 0 {
		return true
	}
	return false
}

func (p *Payment) validateDate() {
	if val, ok := p.JsonData["date"]; (ok && val == nil) || !ok {
		p.Date = data.Date(time.Now()).Date2Str()
		p.JsonData["date"] = p.Date
	}
	if val := p.check("string", "date"); val != nil {
		p.Date = val.(string)
		if _, err := time.Parse("2006-01-02", p.Date); err != nil {
			p.validMessages.AddMessage(
				valueObjects.GenMessage("date", "Is wrong format of date."))
		}
	}
}

func (p *Payment) check(type_ string, name string) interface{} {
	if val, ok := p.JsonData[name]; ok && val == nil {
		p.validMessages.AddMessage(
			valueObjects.GenMessage(name, "Is empty."))
		return nil
	}
	if val, ok := p.JsonData[name]; ok && val != nil && reflect.TypeOf(val).String() == type_ {
		return val
	} else {
		if ok {
			if type_ == "float64" {
				type_ = "integer number"
			}
			p.validMessages.AddMessage(
				valueObjects.GenMessage(name, "Must be "+type_+"."))
		} else {
			p.validMessages.AddMessage(
				valueObjects.GenMessage(name, "No field."))
		}
		return nil
	}
}

func (p *Payment) validatePayment() {
	if val := p.check("float64", "amount"); val != nil {
		p.AmountOfPayment = int64(val.(float64))
		if p.GetPayment() <= 0 && p.GetPayment() > 3000000000000000 {
			p.validMessages.AddMessage(
				valueObjects.GenMessage("payment", "Must be between 1 and 3000000000000000"))
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

func (p *Payment) validateState() {
	if val := p.check("string", "state"); val != nil {
		p.State = val.(string)
		if p.GetState() == "" {
			p.validMessages.AddMessage(
				valueObjects.GenMessage("state", "Is unknown state."))
		}
	}
}

func (p *Payment) GetValidation() valueObjects.ValidationInterface {
	return &p.validMessages
}

func (p *Payment) Hydrate(payment valueObjects.PaymentInterface) error {
	return nil
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
