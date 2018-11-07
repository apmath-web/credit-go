package viewModels

import (
	"time"
)

type CreditInterface interface {
	ViewModelInterface
	hydrate(credit interface{}) error
	getPerson() (interface{}, error)
	getAmount() (int64, error)
	getAgreementAt() (time.Time, error)
	getCurrency() (string, error)
	getDuration() (int32, error)
	getPercent() (int32, error)
}
