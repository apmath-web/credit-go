package viewModels

import "credit-go/valueObjects"

type PersonInterface interface {
	ViewModelInterface
	hydrate(person valueObjects.PersonInterface) error
	getFirstName() string
	getLastName() string
}
