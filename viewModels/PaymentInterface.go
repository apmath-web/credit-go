package viewModels

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/valueObjects"
)

type PaymentInterface interface {
	ViewModelInterface
	Hydrate(payment valueObjects.PaymentInterface)
	GetPayment() data.Money
	GetType() data.Type
	GetCurrency() data.Currency
	GetDate() data.Date
	GetState() data.State
}
