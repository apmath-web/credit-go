package tests_test

import (
	"github.com/apmath-web/credit-go/valueObjects"
	"github.com/apmath-web/credit-go/viewModels"
	"reflect"
	"testing"
)

func TestPersonViewCreation(t *testing.T) {
	req := map[string]interface{}{
		"firstName": "Fname",
		"lastName":  "Lname",
	}
	a := new(viewModels.Person)
	a.Fill(req)
	a.Validate()
	if a.GetLastName() != "Lname" {
		t.Errorf("Don't fill firstName. "+
			"Got: %+v. Want: %+v.", a.GetLastName(), "Lname")
	}
	if a.GetFirstName() != "Fname" {
		t.Errorf("Don't fill lastName. "+
			"Got: %+v. Want: %+v.", a.GetFirstName(), "Fname")
	}
}

func TestPersonViewValidationPos(t *testing.T) {
	req := map[string]interface{}{
		"firstName": "Fname",
		"lastName":  "Lname",
	}
	a := new(viewModels.Person)
	a.Fill(req)
	if !a.Validate() {
		t.Errorf("Wrong validation.")
	}
	validator := a.GetValidation()
	messages := validator.GetMessages()
	if len(messages) != 0 {
		t.Errorf("Error in parsing. Got: %v", messages)
	}
}

func TestPersonViewValidationNeg(t *testing.T) {
	req := map[string]interface{}{
		"firstName": "Fname",
	}
	a := new(viewModels.Person)
	a.Fill(req)
	if a.Validate() {
		t.Errorf("Wrong validation.")
	}
	validator := a.GetValidation()
	messages := validator.GetMessages()
	if len(messages) != 1 {
		t.Errorf("Error in parsing. Got: %+v", messages)
	} else {
		total := valueObjects.GenMessage("lastName", "No field.")
		if !reflect.DeepEqual(total, messages[0]) {
			t.Errorf("Wrong message. Got: %+v. Want: %+v.",
				messages[0], total)
		}
	}
}
