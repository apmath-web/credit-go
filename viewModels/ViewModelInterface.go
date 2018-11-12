package viewModels

import (
	"github.com/apmath-web/credit-go/valueObjects"
)

type ViewModelInterface interface {
	Fill(jsonData map[string]interface{}) (bool, valueObjects.ValidationInterface)
	Fetch() (interface{}, error)
	Validate() bool
	GetValidation() valueObjects.ValidationInterface
}
