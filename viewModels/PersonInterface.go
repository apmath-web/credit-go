package viewModels

type PersonInterface interface {
	ViewModelInterface
	hydrate(person interface{}) error
	getFirstName() (string, error)
	getLastName() (string, error)
}
