package viewModels

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/valueObjects"
)

type PaymentInterface interface {
	ViewModelInterface
	hydrate(payment valueObjects.PaymentInterface) error
	getPayment() data.Money
	getType() data.Type
	getCurrency() data.Currency
	getDate() data.Date
	getState() data.State
}
