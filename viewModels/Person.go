package viewModels

import "C"
import (
	"github.com/apmath-web/credit-go/valueObjects"
)

type Person struct {
	viewModel
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (p *Person) Fetch() interface{} {
	jsonData := make(map[string]string)
	jsonData["firstName"] = p.FirstName
	jsonData["lastName"] = p.LastName
	return jsonData
}

func (p *Person) Validate() bool {
	p.validateFirstName()
	p.validateLastName()
	if len(p.validMessages.GetMessages()) == 0 {
		return true
	}
	return false
}

func (p *Person) validateFirstName() {
	if val := p.check("string", "firstName"); val != nil {
		p.FirstName = val.(string)
		if p.FirstName == "" {
			p.validMessages.AddMessage(valueObjects.GenMessage("Is empty", "firstName"))
		}
	}
}

func (p *Person) validateLastName() {
	if val := p.check("string", "lastName"); val != nil {
		p.LastName = val.(string)
		if p.LastName == "" {
			p.validMessages.AddMessage(valueObjects.GenMessage("Is empty", "lastName"))
		}
	}
}

func (p *Person) Hydrate(person valueObjects.PersonInterface) {
	p.FirstName = person.GetFirstName()
	p.LastName = person.GetLastName()
}

func (p *Person) GetFirstName() string {
	return p.FirstName
}

func (p *Person) GetLastName() string {
	return p.LastName
}
