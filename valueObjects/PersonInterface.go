package valueObjects

type PersonInterface interface {
	Person(firstName string, lastName string) error
	getFirstName() string
	getLastName() string
}

/*
Just an example how you should implement an interface
*/

type Person struct {
	firstName, lastName string
}

func (p Person) Person(firstName string, lastName string) {
	p.firstName = firstName
	p.lastName = lastName
}

func (p Person) getFirstName() string {
	return p.firstName
}

func (p Person) getLastName() string {
	return p.lastName
}
