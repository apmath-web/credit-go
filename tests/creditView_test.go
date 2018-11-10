package tests_test

import (
	"github.com/apmath-web/credit-go/tests"
	"github.com/apmath-web/credit-go/viewModels"
	"testing"
)

func TestCreditViewCreation(t *testing.T) {
	req := tests.GenerateRequest(
		"{\"person\":{\"firstName\":\"FName\",\"lastName\":\"LName\"}," +
			"\"amount\":2000, \"agreementAt\":\"2018-10-10\", \"currency\":\"RUB\"," +
			"\"duration\":6, \"percent\":10, \"rounding\":10}")
	a := new(viewModels.Credit)
	if ok, err := a.Fill(req); !ok {
		t.Errorf("Can't parse. Error %v", err)
	}
	if a.GetPerson().GetLastName() != "LName" {
		t.Errorf("Don't fill FirstName. Got: %+v. "+
			"Want: %+v.", a.GetPerson().GetLastName(), "Lname")
	}
	if a.GetPerson().GetFirstName() != "FName" {
		t.Errorf("Don't fill LastName. Got: %+v. "+
			"Want: %+v.", a.GetPerson().GetFirstName(), "Fname")
	}
	if a.GetAmount() != 2000 {
		t.Errorf("Don't fill Amount. Got: %+v. "+
			"Want: %+v.", a.GetAmount(), 2000)
	}
	if a.GetAgreementAt() != "2018-10-10" {
		t.Errorf("Don't fill AgreementAt. Got: %+v. "+
			"Want: %+v.", a.GetAgreementAt(), "2018-10-10")
	}
	if a.GetCurrency() != "RUB" {
		t.Errorf("Don't fill Curency. Got: %+v. "+
			"Want: %+v.", a.GetCurrency(), "RUB")
	}
	if a.GetDuration() != 6 {
		t.Errorf("Don't fill Duration. Got: %+v. "+
			"Want: %+v.", a.GetDuration(), 6)
	}
	if a.GetPercent() != 10 {
		t.Errorf("Don't fill Percent. Got: %+v. "+
			"Want: %+v.", a.GetPercent(), 10)
	}
	if a.GetRounding() != 10 {
		t.Errorf("Don't fill Rounding. Got: %+v. "+
			"Want: %+v.", a.GetRounding(), 10)
	}
}

/*func TestCreditViewValidationPos(t *testing.T) {
	req := tests.GenerateRequest("{\"FirstName\":\"FName\",\"LastName\":\"LName\"}")
	a := new(viewModels.Person)
	if _, err := a.Fill(req); err != nil {
		t.Errorf("Can't parse. Error %v", err)
	}
	a.Validate()
	validator := a.GetValidation()
	messages := validator.GetMessages()
	if len(messages) != 0 {
		t.Errorf("Error in parsing. Error %v", messages)
	}
}

func TestCreditViewValidationNeg(t *testing.T) {
	req := tests.GenerateRequest("{\"FirstName\":\"FName\"}")
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
}*/
