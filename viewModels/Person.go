package viewModels

import "C"
import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apmath-web/credit-go/valueObjects"
	"net/http"
)

type Person struct {
	validMessages valueObjects.Validation
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
}

func (p *Person) Fill(JsonData *http.Request) (bool, error) {
	body := JsonData.Body
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(p); err != nil {
		fmt.Println(body, p)
		return false, err
	}
	return true, nil
}

func (p *Person) Fetch() (interface{}, error) {
	return 0, errors.New("Not implement\n")
}

func (p *Person) Validate() bool {
	if p.FirstName == "" {
		p.validMessages.AddMessages(
			valueObjects.GenMessageInArray("FistName", "Is empty."))
	}
	if p.LastName == "" {
		p.validMessages.AddMessages(
			valueObjects.GenMessageInArray("LastName", "Is empty."))
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
