package actions

import (
	"fmt"
	"github.com/apmath-web/credit-go/models"
	"github.com/apmath-web/credit-go/valueObjects"
	"github.com/apmath-web/credit-go/viewModels"
	"net/http"
)

func Create(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type",
		"application/json; charset=utf-8")
	jsonData := toJson(response, request)
	if jsonData == nil {
		return
	}
	creditViewModel := new(viewModels.Credit)
	creditViewModel.Fill(jsonData)
	ok := creditViewModel.Validate()
	if !ok {
		jsonData := ptrMessagesToJsonErrMessage("Validation error",
			creditViewModel.GetValidation().GetMessages())
		response.WriteHeader(400)
		fmt.Fprint(response, jsonData)
		return
	}
	personViewModel := creditViewModel.GetPerson()
	person := valueObjects.GenPerson(
		personViewModel.GetFirstName(), personViewModel.GetLastName())
	credit, err := models.GenCredit(person, creditViewModel.GetAmount(),
		creditViewModel.GetAgreementAt(), creditViewModel.GetCurrency(),
		creditViewModel.GetDuration(), creditViewModel.GetPercent())
	if err != nil {
		errorMessage(err.Error(), 400, response)
		return
	}
	repo := Repository
	repo.Store(credit)
	fmt.Fprintf(response, "{\"id\": %d }", credit.GetId())
}
