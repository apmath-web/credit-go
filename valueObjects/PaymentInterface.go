package valueObjects

import "time"

type PaymentInterface interface {
	Payment(payment int, type_ string, currency string, date time.Time) error
	getPayment()
	getType() string
	getCurrency() string
	getDate() time.Time
	getState() string
}
