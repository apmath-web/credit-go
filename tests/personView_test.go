package tests_test

import (
	"github.com/apmath-web/credit-go/tests"
	"github.com/apmath-web/credit-go/valueObjects"
	"github.com/apmath-web/credit-go/viewModels"
	"reflect"
	"testing"
)

func TestPersonViewCreation(t *testing.T) {
	req := tests.GenerateRequest(
		"{\"firstName\":\"FName\",\"lastName\":\"LName\"}")
	a := new(viewModels.Person)
	if _, err := a.Fill(req); err != nil {
		t.Errorf("Can't parse. Error %v", err)
	}
	if a.GetLastName() != "LName" {
		t.Errorf("Don't fill FirstName. "+
			"Got: %+v. Want: %+v.", a.GetLastName(), "Lname")
	}
	if a.GetFirstName() != "FName" {
		t.Errorf("Don't fill LastName. "+
			"Got: %+v. Want: %+v.", a.GetFirstName(), "Fname")
	}
}

func TestPersonViewValidationPos(t *testing.T) {
	req := tests.GenerateRequest(
		"{\"firstName\":\"FName\",\"lastName\":\"LName\"}")
	a := new(viewModels.Person)
	if _, err := a.Fill(req); err != nil {
		t.Errorf("Can't parse. Error %v", err)
	}
	if !a.Validate() {
		t.Errorf("Wrong validation.")
	}
	validator := a.GetValidation()
	messages := validator.GetMessages()
	if len(messages) != 0 {
		t.Errorf("Error in parsing. Error %v", messages)
	}
}

func TestPersonViewValidationNeg(t *testing.T) {
	req := tests.GenerateRequest(
		"{\"firstName\":\"FName\"}")
	a := new(viewModels.Person)
	if _, err := a.Fill(req); err != nil {
		t.Errorf("Can't parse. Error %v", err)
	}
	a.Validate()
	validator := a.GetValidation()
	messages := validator.GetMessages()
	if len(messages) != 1 {
		t.Errorf("Error in parsing. Got: %+v", messages)
	} else {
		total := new(valueObjects.Message)
		total.Message("LastName", "Is empty.")
		if !reflect.DeepEqual(total, messages[0]) {
			t.Errorf("Wrong message. Got: %+v. Want: %+v.", messages[0], total)
		}
	}
}
