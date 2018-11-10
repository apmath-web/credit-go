package viewModels

import (
	"encoding/json"
	"fmt"
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/models"
	"github.com/apmath-web/credit-go/valueObjects"
	"net/http"
)

type Credit struct {
	validMessages valueObjects.Validation
	Person        Person        `json:"person"`
	Amount        data.Money    `json:"amount"`
	AgreementAt   data.Date     `json:"agreementAt"`
	Currency      data.Currency `json:"currency"`
	Duration      int32         `json:"duration"`
	Percent       int32         `json:"percent"`
	Rounding      int32         `json:"rounding"`
}

func (c *Credit) Fill(JsonData *http.Request) (bool, error) {
	body := JsonData.Body
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(c); err != nil {
		fmt.Println(body, c)
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
	return c.Amount
}

func (c *Credit) GetAgreementAt() data.Date {
	return c.AgreementAt
}

func (c *Credit) GetCurrency() data.Currency {
	return c.Currency
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
