package models

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/valueObjects"
)

type CreditInterface interface {
	Credit(person valueObjects.PersonInterface, amount int64, agreementAt data.Date,
		currency data.Currency, duration int, percent int)
	getId() int
	getPerson() valueObjects.PersonInterface
	getAmount() int64
	getAgrementAt() data.Date
	getCurrency() data.Currency
	getDuration() int
	getPercent() int

	getRounding() int
	getRemainAmount() data.Money
	getPayments(type_ data.Type, state data.State) []valueObjects.PaymentInterface

	writeOf(payment valueObjects.PaymentInterface) error
}
