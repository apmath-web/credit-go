package models

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/valueObjects"
)

type Credit struct {
	Id           int
	Person       valueObjects.PersonInterface `json:"person"`
	Amount       data.Money                   `json:"amount"`
	AgreementAt  data.Date                    `json:"agreementAt"`
	Currency     data.Currency                `json:"currency"`
	Duration     int32                        `json:"duration"`
	Percent      int32                        `json:"percent"`
	Rounding     int32
	RemainAmount data.Money
	Payments     []valueObjects.PaymentInterface
}

func GenCredit(person valueObjects.PersonInterface, amount data.Money, agreementAt data.Date,
	currency data.Currency, duration int32, percent int32) CreditInterface {
	c := new(Credit)
	c.Person = person
	c.Amount = amount
	c.AgreementAt = agreementAt
	c.Currency = currency
	c.Duration = duration
	c.Percent = percent
	c.Rounding = 0 //TODO: Write a magic function for calculate rounding
	c.Id = -1
	return c
}

func (c *Credit) GetId() int {
	return c.Id
}

func (c *Credit) SetId(id int) {
	if c.Id == -1 {
		c.Id = id
	}
}

func (c *Credit) GetPerson() valueObjects.PersonInterface {
	return c.Person
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

func (c *Credit) GetRemainAmount() data.Money {
	return c.RemainAmount
}

func (c *Credit) GetPayments(type_ data.Type, state data.State) []valueObjects.PaymentInterface {
	return c.Payments
}

func (c *Credit) WriteOf(payment valueObjects.PaymentInterface) error {
	return nil
}
