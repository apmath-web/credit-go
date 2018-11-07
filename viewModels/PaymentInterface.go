package viewModels

import (
	"credit-go/valueObjects"
	"time"
)

type PaymentInterface interface {
	ViewModelInterface
	hydrate(payment valueObjects.PaymentInterface) error
	getPayment() error
	getType() string
	getCurrency() string
	getDate() time.Time
	getState() string
}
