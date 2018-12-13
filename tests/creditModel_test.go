package tests

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/models"
	"github.com/apmath-web/credit-go/valueObjects"
	"testing"
)

func TestCreditModelCreation(t *testing.T) {
	person := valueObjects.GenPerson("Alexandra", "Chernyshova")
	_, err := models.GenCredit(person, 100, data.Str2Date("2018-10-10"), data.RUR, 10, 10)
	if err == nil {
		t.Errorf("Logic for small payment does not work.")
		return
	}
	_, err = models.GenCredit(person, 1200, data.Str2Date("2018-10-10"), data.RUR, 1200, 10)
	if err == nil {
		t.Errorf("Logic for small rounding does not work.")
		return
	}
	_, err = models.GenCredit(person, 1200, data.Str2Date("2018-10-10"), data.RUR, 10, 10)
	if err != nil {
		t.Errorf("Starnge error.")
		return
	}
}
