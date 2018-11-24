package repositories

import (
	"errors"
	"github.com/apmath-web/credit-go/models"
)

type CreditRepository struct {
	credits         map[int]models.CreditInterface
	numberOfCredits int
}

func GenRepository() CreditsRepositoryInterface {
	repo := &CreditRepository{make(map[int]models.CreditInterface), 0}
	return repo
}

func (r *CreditRepository) Get(id int) models.CreditInterface {
	credit, ok := r.credits[id]
	if ok {
		return credit
	}
	return nil
}
func (r *CreditRepository) Store(credit models.CreditInterface) {
	r.numberOfCredits++
	credit.SetId(r.numberOfCredits)
	r.credits[r.numberOfCredits] = credit
}

func (r *CreditRepository) Remove(credit models.CreditInterface) error {
	return errors.New("Not implement")
}

var Repository = GenRepository()
