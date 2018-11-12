package viewModels

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/models"
	"github.com/apmath-web/credit-go/valueObjects"
)

type CreditInterface interface {
	ViewModelInterface
	Hydrate(credit models.CreditInterface) error
	GetPerson() valueObjects.PersonInterface
	GetAmount() data.Money
	GetAgreementAt() data.Date
	GetCurrency() data.Currency
	GetDuration() int32
	GetPercent() int32
}
