package valueObjects

import (
	"github.com/apmath-web/credit-go/data"
)

type PaymentInterface interface {
	GetPayment() data.Money
	GetType() data.Type
	GetCurrency() data.Currency
	GetDate() data.Date
	GetState() data.State
	GetPercent() data.Money
	GetBody() data.Money
	GetRemainCreditBody() data.Money
	GetFullEarlyRepayment() data.Money
}
