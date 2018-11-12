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
	Duration     int                          `json:"duration"`
	Percent      int                          `json:"percent"`
	Rounding     int
	RemainAmount data.Money
	Payments     []valueObjects.PaymentInterface
}

func (c *Credit) Credit(person valueObjects.PersonInterface, amount data.Money, agreementAt data.Date,
	currency data.Currency, duration int, percent int, rounding int) {
	c.Person = person
	c.Amount = amount
	c.AgreementAt = agreementAt
	c.Currency = currency
	c.Duration = duration
	c.Percent = percent
	c.Rounding = rounding
}

func (c *Credit) GetId() int {
	return c.Id
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

func (c *Credit) GetDuration() int {
	return c.Duration
}

func (c *Credit) GetPercent() int {
	return c.Percent
}

func (c *Credit) GetRounding() int {
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
