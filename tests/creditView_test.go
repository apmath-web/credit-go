package tests_test

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/tests"
	"github.com/apmath-web/credit-go/valueObjects"
	"github.com/apmath-web/credit-go/viewModels"
	"reflect"
	"testing"
	"time"
)

func TestCreditViewCreation(t *testing.T) {
	req := tests.GenerateRequest(
		"{\"person\":{\"firstName\":\"FName\",\"lastName\":\"LName\"}," +
			"\"amount\":2000, \"agreementAt\":\"2018-10-10\", \"currency\":\"RUR\"," +
			"\"duration\":6, \"percent\":10}")
	date, _ := time.Parse("2006-01-02", "2018-10-10")
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
	if a.GetAmount() != data.Money(2000) {
		t.Errorf("Don't fill Amount. Got: %+v. "+
			"Want: %+v.", a.GetAmount(), 2000)
	}
	if !reflect.DeepEqual(a.GetAgreementAt(), data.Date(date)) {
		t.Errorf("Don't fill AgreementAt. Got: %+v. "+
			"Want: %+v.", a.GetAgreementAt().Date2Str(), "2018-10-10")
	}
	if a.GetCurrency() != data.RUR {
		t.Errorf("Don't fill Curency. Got: %+v. "+
			"Want: %+v.", a.GetCurrency(), "RUR")
	}
	if a.GetDuration() != 6 {
		t.Errorf("Don't fill Duration. Got: %+v. "+
			"Want: %+v.", a.GetDuration(), 6)
	}
	if a.GetPercent() != 10 {
		t.Errorf("Don't fill Percent. Got: %+v. "+
			"Want: %+v.", a.GetPercent(), 10)
	}
}

func TestCreditViewValidationPos(t *testing.T) {
	req := tests.GenerateRequest(
		"{\"person\":{\"firstName\":\"FName\",\"lastName\":\"LName\"}," +
			"\"amount\":2000, \"agreementAt\":\"2018-10-10\", \"currency\":\"RUR\"," +
			"\"duration\":6, \"percent\":10}")
	a := new(viewModels.Credit)
	if ok, err := a.Fill(req); !ok {
		t.Errorf("Can't parse. Error %+v", err)
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

func TestCreditViewValidationNeg(t *testing.T) {
	req := tests.GenerateRequest(
		"{\"person\":{\"firstName\":\"FName\",\"lastName\":\"LName\"}," +
			"\"amount\":200, \"agreementAt\":\"2018-10-10\", \"currency\":\"RUR\"," +
			"\"duration\":6, \"percent\":301}")
	a := new(viewModels.Credit)
	if _, err := a.Fill(req); err != nil {
		t.Errorf("Can't parse. Error %v", err)
	}
	if a.Validate() {
		t.Errorf("Wrong validation.")
	}
	validator := a.GetValidation()
	messages := validator.GetMessages()
	if len(messages) != 1 {
		t.Errorf("Error in parsing. Got: %+v", messages)
	} else {
		total := valueObjects.GenMessageInArray("Percent",
			"Is wrong value. Minimum 1%, maximum 300%.")
		if !reflect.DeepEqual(total, messages) {
			t.Errorf("Wrong message. Got: %+v. Want: %+v.", messages, total)
		}
	}
}
