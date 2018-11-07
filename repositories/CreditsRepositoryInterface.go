package repositories

import "credit-go/models"

type CreditsRepositoryInterface interface {
	get(id int) models.CreditInterface
	store(credit models.CreditInterface) error
	remove(credit models.CreditInterface) error
}
