package models

import (
	"credit-go/valueObjects"
	"time"
)

type CreditInterface interface {
	getId() int
	getPerson() valueObjects.PersonInterface
	getAmount() int64
	getAgrementAt() time.Time
	getCurrency() string
	getDuration() int
	getPercent() int

	getRounding() int
	getRemainAmount()
	getPayments(type_ string, state string) []valueObjects.PaymentInterface

	writeOf(payment valueObjects.PaymentInterface) error
}
