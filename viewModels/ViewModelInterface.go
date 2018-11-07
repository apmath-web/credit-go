package viewModels

import "github.com/apmath-web/credit-go/valueObjects"

type ViewModelInterface interface {
	Fill(JsonData interface{}) (bool, error)
	Fetch() (interface{}, error)
	Validate() (bool, error)
	getValidation() (valueObjects.ValidationInterface, error)
}
