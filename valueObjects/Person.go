package valueObjects

type Person struct {
	FirstName, LastName string
}

func GenPerson(firstName string, lastName string) PersonInterface {
	p := new(Person)
	p.FirstName = firstName
	p.LastName = lastName
	return p
}

func (p *Person) GetFirstName() string {
	return p.FirstName
}

func (p *Person) GetLastName() string {
	return p.LastName
}
