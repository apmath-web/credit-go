package models

import (
	"errors"
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/valueObjects"
	"math"
)

type Credit struct {
	Id             int
	Person         valueObjects.PersonInterface `json:"person"`
	Amount         data.Money                   `json:"amount"`
	AgreementAt    data.Date                    `json:"agreementAt"`
	Currency       data.Currency                `json:"currency"`
	Duration       int32                        `json:"duration"`
	Percent        int32                        `json:"percent"`
	Rounding       int32
	RemainAmount   data.Money
	regularPayment data.Money
	Payments       []valueObjects.PaymentInterface
}

func GenCredit(person valueObjects.PersonInterface, amount data.Money, agreementAt data.Date,
	currency data.Currency, duration int32, percent int32) (CreditInterface, error) {
	c := new(Credit)
	c.Person = person
	c.Amount = amount
	c.AgreementAt = agreementAt
	c.Currency = currency
	c.Duration = duration
	c.Percent = percent
	annuityPayment := c.getAnnuityPayment()
	rounding, err := c.getRounding(annuityPayment)
	if err != nil {
		return nil, err
	}
	c.Rounding = rounding
	c.regularPayment = c.getRegularPayment(annuityPayment)
	c.Id = -1
	return c, nil
}

func (c *Credit) GetId() int {
	return c.Id
}

func (c *Credit) SetId(id int) {
	if c.Id == -1 {
		c.Id = id
	} else {
		panic("Your set one more ID!")
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

func (c *Credit) IsFinished() bool {
	// todo
	return false
}

func (c *Credit) GetPayments(type_ data.Type, state data.State) []valueObjects.PaymentInterface {
	// todo
	return c.Payments
}

func (c *Credit) WriteOf(payment valueObjects.PaymentInterface) error {
	// todo
	return nil
}

func (c *Credit) getAnnuityPayment() float64 {
	monthPercent := float64(c.Percent) / 12.0 / 100.0
	power := math.Pow(1.0+monthPercent, float64(c.Duration))
	return c.Amount.Mon2Float64() * monthPercent * (power / (power - 1.0))
}

func (c *Credit) getRounding(annuityPayment float64) (int32, error) {
	if annuityPayment < 100 {
		return -1, errors.New("Credit payment is less than 100.")
	}
	for _, round := range []int{100, 10, 1} {
		if (round-(int(annuityPayment)%round))*int(c.Duration) < int(annuityPayment) {

			return int32(round), nil
		}
	}
	return -1, errors.New("Credit amount too small for rounding.")
}

func (c *Credit) getRegularPayment(annuityPayment float64) data.Money {
	return data.Money(math.Ceil(annuityPayment/float64(c.Rounding)) * float64(c.Rounding))
}
