package valueObjects

type Person struct {
	FirstName, LastName string
}

func (p *Person) Person(firstName string, lastName string) {
	p.FirstName = firstName
	p.LastName = lastName
}

func (p *Person) GetFirstName() string {
	return p.FirstName
}

func (p *Person) GetLastName() string {
	return p.LastName
}
