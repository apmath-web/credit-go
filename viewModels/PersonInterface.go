package viewModels

import "github.com/apmath-web/credit-go/valueObjects"

type PersonInterface interface {
	ViewModelInterface
	hydrate(person valueObjects.PersonInterface) error
	getFirstName() string
	getLastName() string
}
