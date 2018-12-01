package viewModels

import "github.com/apmath-web/credit-go/valueObjects"

type PersonInterface interface {
	ViewModelInterface
	Hydrate(person valueObjects.PersonInterface)
	GetFirstName() string
	GetLastName() string
}
