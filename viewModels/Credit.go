package viewModels

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/models"
	"github.com/apmath-web/credit-go/valueObjects"
	"reflect"
	"time"
)

type Credit struct {
	validMessages valueObjects.Validation
	Person        Person `json:"person"`
	Amount        int64  `json:"amount"`
	AgreementAt   string `json:"agreementAt"`
	Currency      string `json:"currency"`
	Duration      int32  `json:"duration"`
	Percent       int32  `json:"percent"`
	JsonData      map[string]interface{}
}

func (c *Credit) Fill(jsonData map[string]interface{}) bool {
	c.JsonData = jsonData
	if ok, val := c.check("map[string]interface {}", "person"); ok {
		c.Person.Fill(val.(map[string]interface{}))
		return true
	}
	return false
}

func (c *Credit) Fetch() (interface{}, error) {
	return 0, nil
}

func (c *Credit) check(type_ string, name string) (bool, interface{}) {

	if val, ok := c.JsonData[name]; ok && reflect.TypeOf(val).String() == type_ {
		if val == nil {
			c.validMessages.AddMessages(
				valueObjects.GenMessageInArray(name, "Is empty."))
			return false, nil
		}
		return true, val
	} else {
		if ok {
			c.validMessages.AddMessages(
				valueObjects.GenMessageInArray(name, "Must be "+type_+"."))
		} else {
			c.validMessages.AddMessages(
				valueObjects.GenMessageInArray(name, "No field."))
		}
		return false, nil
	}
}

func (c *Credit) Validate() bool {
	if !c.Person.Validate() {
		c.validMessages.AddMessages(c.Person.GetValidation().GetMessages())
	}
	if ok, val := c.check("int", "amount"); ok {
		c.Amount = int64(val.(int))
		if c.GetAmount() <= 0 {
			c.validMessages.AddMessages(
				valueObjects.GenMessageInArray("Amount", "Wrong amount value"))
		}
	}
	if val, ok := c.JsonData["agreementAt"]; ok && reflect.TypeOf(val).String() == "string" {
		if val == nil {
			c.AgreementAt = data.Date(time.Now()).Date2Str()
		}
	}
	if ok, val := c.check("string", "agreementAt"); ok {
		c.AgreementAt = val.(string)
		if _, err := time.Parse("2006-01-02", c.AgreementAt); err != nil {
			c.validMessages.AddMessages(
				valueObjects.GenMessageInArray("AgreementAt", "Is wrong format of date."))
		}
	}
	if ok, val := c.check("string", "currency"); ok {
		c.Currency = val.(string)
		if c.GetCurrency() == "" {
			c.validMessages.AddMessages(
				valueObjects.GenMessageInArray("Currency", "Is unknown currency."))
		}
	}
	if ok, val := c.check("int", "duration"); ok {
		c.Duration = int32(val.(int))
		if c.GetDuration() < 6 || c.GetDuration() > 1200 {
			c.validMessages.AddMessages(
				valueObjects.GenMessageInArray("Duration",
					"Is wrong value. Minimum 6 months, maximum 1200."))
		}
	}
	if ok, val := c.check("int", "percent"); ok {
		c.Percent = int32(val.(int))
		if c.GetPercent() < 1 || c.GetPercent() > 300 {
			c.validMessages.AddMessages(
				valueObjects.GenMessageInArray("Percent",
					"Is wrong value. Minimum 1%, maximum 300%."))
		}
	}
	if len(c.validMessages.GetMessages()) == 0 {
		return true
	}
	return false
}

func (c *Credit) GetValidation() valueObjects.ValidationInterface {
	return &c.validMessages
}

func (c *Credit) Hydrate(credit models.CreditInterface) error {
	return nil
}

func (c *Credit) GetPerson() PersonInterface {
	return &c.Person
}

func (c *Credit) GetAmount() data.Money {
	return data.Money(c.Amount)
}

func (c *Credit) GetAgreementAt() data.Date {
	return data.Str2Date(c.AgreementAt)
}

func (c *Credit) GetCurrency() data.Currency {
	return data.Str2Cur(c.Currency)
}

func (c *Credit) GetDuration() int32 {
	return c.Duration
}

func (c *Credit) GetPercent() int32 {
	return c.Percent
}
