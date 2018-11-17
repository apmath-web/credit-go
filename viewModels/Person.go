package viewModels

import "C"
import (
	"errors"
	"github.com/apmath-web/credit-go/valueObjects"
	"reflect"
)

type Person struct {
	validMessages valueObjects.Validation
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	JsonData      map[string]interface{}
}

func (p *Person) Fill(jsonData map[string]interface{}) bool {
	p.JsonData = jsonData
	return true
}

func (p *Person) Fetch() (interface{}, error) {
	return 0, errors.New("Not implement\n")
}

func (p *Person) check(type_ string, name string) (bool, interface{}) {
	if val, ok := p.JsonData[name]; ok && val == nil {
		p.validMessages.AddMessage(
			valueObjects.GenMessage(name, "Is empty."))
		return false, nil
	}
	if val, ok := p.JsonData[name]; ok && val != nil && reflect.TypeOf(val).String() == type_ {
		return true, val
	} else {
		if ok {
			p.validMessages.AddMessage(
				valueObjects.GenMessage(name, "Must be "+type_+"."))
		} else {
			p.validMessages.AddMessage(
				valueObjects.GenMessage(name, "No field."))
		}
		return false, nil
	}
}

func (p *Person) Validate() bool {
	if ok, val := p.check("string", "firstName"); ok {
		p.FirstName = val.(string)
	}
	if ok, val := p.check("string", "lastName"); ok {
		p.LastName = val.(string)
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
	return errors.New("Not implement\n")
}

func (p *Person) GetFirstName() string {
	return p.FirstName
}

func (p *Person) GetLastName() string {
	return p.LastName
}
