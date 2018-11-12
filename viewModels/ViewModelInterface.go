package viewModels

import (
	"github.com/apmath-web/credit-go/valueObjects"
	"net/http"
)

type ViewModelInterface interface {
	Fill(JsonData *http.Request) (bool, valueObjects.ValidationInterface)
	Fetch() (interface{}, error)
	Validate() bool
	GetValidation() valueObjects.ValidationInterface
}
