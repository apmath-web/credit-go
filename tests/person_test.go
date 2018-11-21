package tests

import (
	"github.com/apmath-web/credit-go/valueObjects"
	"testing"
)

func TestPerson(t *testing.T) {
	fname := "FName"
	lname := "LName"
	total := valueObjects.GenPerson(fname, lname)
	if total.GetFirstName() != fname {
		t.Errorf("Don't save any firstname. Got: %+v. Want %+v.",
			total.GetFirstName(), fname)
	}

	if total.GetLastName() != lname {
		t.Errorf("Don't append any messages. Got: %+v. Want: %+v.",
			total.GetLastName(), lname)
	}
}
