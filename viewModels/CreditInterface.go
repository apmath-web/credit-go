package viewModels

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/models"
)

type CreditInterface interface {
	ViewModelInterface
	Hydrate(credit models.CreditInterface) error
	GetPerson() PersonInterface
	GetAmount() data.Money
	GetAgreementAt() data.Date
	GetCurrency() data.Currency
	GetDuration() int32
	GetPercent() int32
}
