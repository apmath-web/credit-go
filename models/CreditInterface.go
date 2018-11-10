package models

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/valueObjects"
)

type CreditInterface interface {
	Credit(person valueObjects.PersonInterface, amount int64, agreementAt data.Date,
		currency data.Currency, duration int, percent int)
	GetId() int
	GetPerson() valueObjects.PersonInterface
	GetAmount() data.Money
	GetAgreementAt() data.Date
	GetCurrency() data.Currency
	GetDuration() int
	GetPercent() int

	GetRounding() int
	GetRemainAmount() data.Money
	GetPayments(type_ data.Type, state data.State) []valueObjects.PaymentInterface

	WriteOf(payment valueObjects.PaymentInterface) error
}
