package models

import (
	"github.com/apmath-web/credit-go/valueObjects"
	"time"
)

type CreditInterface interface {
	Credit(person valueObjects.PersonInterface, amount int64, agreementAt time.Time,
		currency string, duration int, percent int)
	getId() int
	getPerson() valueObjects.PersonInterface
	getAmount() int64
	getAgrementAt() time.Time
	getCurrency() string
	getDuration() int
	getPercent() int

	getRounding() int
	getRemainAmount() data.Money
	getPayments(type_ data.Type, state data.State) []valueObjects.PaymentInterface

	writeOf(payment valueObjects.PaymentInterface) error
}
