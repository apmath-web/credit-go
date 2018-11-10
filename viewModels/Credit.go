package viewModels

import (
	"encoding/json"
	"fmt"
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/models"
	"github.com/apmath-web/credit-go/valueObjects"
	"net/http"
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
	Rounding      int32  `json:"rounding"`
}

func (c *Credit) Fill(JsonData *http.Request) (bool, error) {
	body := JsonData.Body
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(c); err != nil {
		fmt.Println(body, c)
		return false, err
	}
	if c.AgreementAt == "" {
		c.AgreementAt = data.Date(time.Now()).Date2Str()
	}
	if _, err := time.Parse("2006-01-02", c.AgreementAt); err != nil {
		return false, err
	}
	return true, nil
}

func (c *Credit) Fetch() (interface{}, error) {
	return 0, nil
}

func (c *Credit) Validate() bool {
	if !c.Person.Validate() {
		c.validMessages.AddMessages(c.Person.GetValidation().GetMessages())
	}
	if c.GetAmount() <= 0 {
		c.validMessages.AddMessages(
			valueObjects.GenMessageInArray("Amount", "Wrong amount value"))
	}
	if _, err := time.Parse("2006-01-02", c.AgreementAt); err != nil {
		c.validMessages.AddMessages(
			valueObjects.GenMessageInArray("AgreementAt", "Is wrong format of date."))
	}
	if c.GetCurrency() == "" {
		c.validMessages.AddMessages(
			valueObjects.GenMessageInArray("Currency", "Is unknown currency."))
	}
	if c.GetDuration() < 6 || c.GetDuration() > 1200 {
		c.validMessages.AddMessages(
			valueObjects.GenMessageInArray("Duration",
				"Is wrong value. Minimum 6 months, maximum 1200."))
	}
	if c.GetPercent() < 1 || c.GetPercent() > 300 {
		c.validMessages.AddMessages(
			valueObjects.GenMessageInArray("Percent",
				"Is wrong value. Minimum 1%, maximum 300%."))
	}
	if c.GetRounding() != 1 && c.GetRounding() != 10 && c.GetRounding() != 100 {
		c.validMessages.AddMessages(
			valueObjects.GenMessageInArray("Rounding",
				"Is wrong value. Only 1, 10, 100."))
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

func (c *Credit) GetRounding() int32 {
	return c.Rounding
}
