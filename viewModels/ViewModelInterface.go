package viewModels

import "github.com/apmath-web/credit-go/valueObjects"

type ViewModelInterface interface {
	Fill(JsonData interface{}) (bool, error)
	Fetch() (interface{}, error)
	Validate() (bool, error)
	GetValidation() (valueObjects.ValidationInterface, error)
}
