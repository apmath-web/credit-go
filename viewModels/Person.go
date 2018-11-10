package viewModels

import (
	"github.com/apmath-web/credit-go/valueObjects"
)

type SimplePerson struct {
	FirstName string
	LastName  string
}

type Person struct {
	validMessages valueObjects.Validation
	FirstName     string
	LastName      string
}

func (p *Person) Fill(JsonData interface{}) (bool, error) {
	return false, nil
}
func (p *Person) Fetch() (interface{}, error) {
	return 0, nil
}
func (p *Person) Validate() bool {
	if p.FirstName == "" {
		message := new(valueObjects.Message)
		message.Message("FistName", "Is empty.")
		p.validMessages.AddMessages([]valueObjects.MessageInterface{message})
	}
	if p.LastName == "" {
		message := new(valueObjects.Message)
		message.Message("LastName", "Is empty.")
		p.validMessages.AddMessages([]valueObjects.MessageInterface{message})
	}
	if len(p.validMessages.GetMessages()) == 0 {
		return true
	}
	return false
}
func (p *Person) GetValidation() valueObjects.ValidationInterface {
	return &p.validMessages
}
func (p *Person) Hydrate(person valueObjects.PersonInterface) error {
	return nil
}
func (p *Person) GetFirstName() string {
	return p.FirstName
}
func (p *Person) GetLastName() string {
	return p.LastName
}
