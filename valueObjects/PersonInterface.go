package valueObjects

type PersonInterface interface {
	Person(firstName string, lastName string)
	GetFirstName() string
	GetLastName() string
}
