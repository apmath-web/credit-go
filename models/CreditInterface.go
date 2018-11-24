package models

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/valueObjects"
)

type CreditInterface interface {
	GetId() int
	SetId(id int)
	GetPerson() valueObjects.PersonInterface
	GetAmount() data.Money
	GetAgreementAt() data.Date
	GetCurrency() data.Currency
	GetDuration() int32
	GetPercent() int32

	GetRounding() int32
	GetRemainAmount() data.Money
	GetPayments(type_ data.Type, state data.State) []valueObjects.PaymentInterface

	WriteOf(payment valueObjects.PaymentInterface) error
}
