package viewModels

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/models"
	"github.com/apmath-web/credit-go/valueObjects"
)

type CreditInterface interface {
	ViewModelInterface
	hydrate(credit models.CreditInterface) error
	getPerson() valueObjects.PersonInterface
	getAmount() data.Money
	getAgreementAt() data.Date
	getCurrency() data.Currency
	getDuration() int32
	getPercent() int32
}
