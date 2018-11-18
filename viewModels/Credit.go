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
	Person        valueObjects.PersonInterface `json:"person"`
	Amount        int64                        `json:"amount"`
	AgreementAt   string                       `json:"agreementAt"`
	Currency      string                       `json:"currency"`
	Duration      int32                        `json:"duration"`
	Percent       int32                        `json:"percent"`
	JsonData      map[string]interface{}
}

func (c *Credit) Fill(jsonData map[string]interface{}) bool {
	c.JsonData = jsonData

	return true
}

func (c *Credit) Fetch() (interface{}, error) {
	return 0, nil
}

func (c *Credit) check(type_ string, name string) interface{} {
	if val, ok := c.JsonData[name]; ok && val == nil {
		c.validMessages.AddMessage(
			valueObjects.GenMessage(name, "Is empty."))
		return nil
	}
	if val, ok := c.JsonData[name]; ok && val != nil && reflect.TypeOf(val).String() == type_ {
		return val
	} else {
		if ok {
			if type_ == "float64" {
				type_ = "number"
			}
			c.validMessages.AddMessage(
				valueObjects.GenMessage(name, "Must be "+type_+"."))
		} else {
			c.validMessages.AddMessage(
				valueObjects.GenMessage(name, "No field."))
		}
		return nil
	}
}

func (c *Credit) Validate() bool {
	c.validatePerson()
	c.validateAmount()
	c.validateAgreementAt()
	c.validateCurrency()
	c.validateDuration()
	c.validatePercent()
	if len(c.validMessages.GetMessages()) == 0 {
		return true
	}
	return false
}

func (c *Credit) validatePerson() {
	viewPerson := new(Person)
	if val := c.check("map[string]interface {}", "person"); val != nil {
		viewPerson.Fill(val.(map[string]interface{}))
	}
	if !viewPerson.Validate() {
		personValidationMessages := viewPerson.GetValidation().GetMessages()
		for _, message := range personValidationMessages {
			messageValidation := valueObjects.GenMessage("person."+message.GetField(), message.GetText())
			c.validMessages.AddMessage(messageValidation)
		}
	} else {
		c.Person = valueObjects.GenPerson(viewPerson.GetFirstName(), viewPerson.GetLastName())
	}
}

func (c *Credit) validateAgreementAt() {
	if val, ok := c.JsonData["agreementAt"]; (ok && val == nil) || !ok {
		// only for areementAt cause it isn't required
		c.AgreementAt = data.Date(time.Now()).Date2Str()
		c.JsonData["agreementAt"] = c.AgreementAt
	}
	if val := c.check("string", "agreementAt"); val != nil {
		c.AgreementAt = val.(string)
		if _, err := time.Parse("2006-01-02", c.AgreementAt); err != nil {
			c.validMessages.AddMessage(
				valueObjects.GenMessage("agreementAt", "Is wrong format of date."))
		}
	}
}

func (c *Credit) validateAmount() {
	if val := c.check("float64", "amount"); val != nil {
		c.Amount = int64(val.(float64))
		if c.GetAmount() <= 0 {
			c.validMessages.AddMessage(
				valueObjects.GenMessage("amount", "Wrong amount value"))
		}
	}
}

func (c *Credit) validateCurrency() {
	if val := c.check("string", "currency"); val != nil {
		c.Currency = val.(string)
		if c.GetCurrency() == "" {
			c.validMessages.AddMessage(
				valueObjects.GenMessage("currency", "Is unknown currency."))
		}
	}
}

func (c *Credit) validateDuration() {
	if val := c.check("float64", "duration"); val != nil {
		c.Duration = int32(val.(float64))
		if c.GetDuration() < 6 || c.GetDuration() > 1200 {
			c.validMessages.AddMessage(
				valueObjects.GenMessage("duration",
					"Is wrong value. Minimum 6 months, maximum 1200."))
		}
	}
}

func (c *Credit) validatePercent() {
	if val := c.check("float64", "percent"); val != nil {
		c.Percent = int32(val.(float64))
		if c.GetPercent() < 1 || c.GetPercent() > 300 {
			c.validMessages.AddMessage(
				valueObjects.GenMessage("percent",
					"Is wrong value. Minimum 1%, maximum 300%."))
		}
	}
}

func (c *Credit) GetValidation() valueObjects.ValidationInterface {
	return &c.validMessages
}

func (c *Credit) Hydrate(credit models.CreditInterface) error {
	return nil
}

func (c *Credit) GetPerson() valueObjects.PersonInterface {
	return c.Person
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
