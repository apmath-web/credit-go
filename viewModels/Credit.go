package viewModels

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/models"
	"github.com/apmath-web/credit-go/valueObjects"
	"net/http"
	"time"
)

type Credit struct {
	validMessages valueObjects.Validation
}

func (c Credit) Fill(JsonData *http.Request) (bool, error) {
	return false, nil
}
func (c Credit) Fetch() (interface{}, error) {
	return 0, nil
}
func (c Credit) Validate() bool {
	return false
}
func (c Credit) GetValidation() valueObjects.ValidationInterface {
	return nil
}
func (c Credit) Hydrate(credit models.CreditInterface) error {
	return nil
}
func (c Credit) GetPerson() valueObjects.PersonInterface {
	return nil
}
func (c Credit) GetAmount() data.Money {
	return 0
}
func (c Credit) getAgreementAt() data.Date {
	return data.Date(time.Now())
}
func (c Credit) GetCurrency() data.Currency {
	return data.RUR
}
func (c Credit) GetDuration() int32 {
	return 0
}
func (c Credit) GetPercent() int32 {
	return 0
}
