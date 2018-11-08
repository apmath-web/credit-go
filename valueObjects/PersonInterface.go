package valueObjects

type PersonInterface interface {
	Person(firstName string, lastName string) error
	GetFirstName() string
	GetLastName() string
}
