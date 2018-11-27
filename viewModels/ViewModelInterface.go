package viewModels

import (
	"github.com/apmath-web/credit-go/valueObjects"
)

type ViewModelInterface interface {
	Fill(jsonData map[string]interface{})
	Fetch() interface{}
	Validate() bool
	GetValidation() valueObjects.ValidationInterface
}
