package valueObjects

import "time"

type PaymentInterface interface {
	getPayment() error
	getType() (string, error)
	getCurrency() string
	getDate() (time.Time, error)
	getState() (string, error)
}
