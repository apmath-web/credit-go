package tests

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/models"
	"github.com/apmath-web/credit-go/repositories"
	"github.com/apmath-web/credit-go/valueObjects"
	"reflect"
	"testing"
)

func TestCreditRepositoryAddCredit(t *testing.T) {
	repo := repositories.Repository
	person := valueObjects.GenPerson("Alexandra", "Chernyshova")
	credit := models.GenCredit(person, 1000, data.Str2Date("2018-10-10"), data.RUR, 10, 10)
	repo.Store(credit)
	if credit.GetId() != 1 {
		t.Errorf("Not right id: %v", credit.GetId())
	}
	creditFromRepo := repo.Get(1)
	if !reflect.DeepEqual(credit, creditFromRepo) {
		t.Errorf("Not right credit: %v", creditFromRepo)
	}
}
