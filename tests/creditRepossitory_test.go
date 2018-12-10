package tests

import (
	"github.com/apmath-web/credit-go/actions"
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/models"
	"github.com/apmath-web/credit-go/valueObjects"
	"reflect"
	"testing"
)

func TestCreditRepositoryAddCredit(t *testing.T) {
	repo := actions.Repository
	person := valueObjects.GenPerson("Alexandra", "Chernyshova")
	credit, err := models.GenCredit(person, 10000, data.Str2Date("2018-10-10"), data.RUR, 10, 10)
	if err != nil {
		t.Errorf("Error in credit: %v", err.Error())
	}
	repo.Store(credit)
	if credit.GetId() != 1 {
		t.Errorf("Not right id: %v", credit.GetId())
	}
	creditFromRepo := repo.Get(1)
	if !reflect.DeepEqual(credit, creditFromRepo) {
		t.Errorf("Not right credit: %v", creditFromRepo)
	}
}
