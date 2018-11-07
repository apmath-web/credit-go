package valueObjects

type PersonInterface interface {
	getFirstName() (string, error)
	getLastName() (string, error)
}

/*
Just an example how you should implement an interface
*/

type Person struct {
	firstName, lastName string
}

func (p Person) getFirstName() string {
	return p.firstName
}

func (p Person) getLastName() string {
	return p.lastName
}
