package repositories

import "github.com/apmath-web/credit-go/models"

type CreditsRepositoryInterface interface {
	Get(id int) models.CreditInterface
	Store(credit models.CreditInterface) error
	Remove(credit models.CreditInterface) error
}
