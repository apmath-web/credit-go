package valueObjects

import (
	"github.com/apmath-web/credit-go/data"
)

type PaymentInterface interface {
	Payment(payment int, type_ data.Type,
		currency data.Currency, date data.Date) error
	GetPayment() data.Money
	GetType() data.Type
	GetCurrency() data.Currency
	GetDate() data.Date
	GetState() data.State
	GetPercent() int32
	GetBody() data.Money
	GetRemainCreditBody() data.Money
	GetFullEarlyRepayment() data.Money
}
