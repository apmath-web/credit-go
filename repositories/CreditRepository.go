package repositories

import (
	"errors"
	"github.com/apmath-web/credit-go/models"
)

type CreditRepository struct {
	credits           []models.CreditInterface
	number_of_credits int
}

func (r *CreditRepository) Get(id int) models.CreditInterface {
	if id < r.number_of_credits && id >= 0 {
		return r.credits[id]
	}
	return nil
}
func (r *CreditRepository) Store(credit models.CreditInterface) error {
	if credit.GetId() == r.number_of_credits {
		r.number_of_credits++
		r.credits = append(r.credits, credit)
	}
}
func (r *CreditRepository) Remove(credit models.CreditInterface) error {
	return errors.New("Not implement")
}
func (r *CreditRepository) GenId() int {
	return r.number_of_credits
}
