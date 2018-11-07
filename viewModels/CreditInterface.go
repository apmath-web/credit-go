package viewModels

import (
	"credit-go/models"
	"credit-go/valueObjects"
	"time"
)

type CreditInterface interface {
	ViewModelInterface
	hydrate(credit models.CreditInterface) error
	getPerson() valueObjects.PersonInterface
	getAmount() int64
	getAgreementAt() time.Time
	getCurrency() string
	getDuration() int32
	getPercent() int32
}
