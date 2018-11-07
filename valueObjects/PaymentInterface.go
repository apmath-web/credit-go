package valueObjects

import "time"

type PaymentInterface interface {
	Payment(payment int, type_ data.Type,
		currency data.Currency, date data.Date) error
	getPayment() data.Money
	getType() data.Type
	getCurrency() data.Currency
	getDate() data.Date
	getState() data.State
	getPercent() data.Money
	getBody() data.Money
	getRemainCreditBody() data.Money
	getFullEarlyRepayment() data.Money
}
