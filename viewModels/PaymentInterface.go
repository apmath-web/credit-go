package viewModels

import "time"

type PaymentInterface interface {
	ViewModelInterface
	hydrate(payment interface{}) error
	getPayment() error
	getType() (string, error)
	getCurrency() string
	getDate() (time.Time, error)
	getState() (string, error)
}
