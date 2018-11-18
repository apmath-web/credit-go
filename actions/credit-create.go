package actions

import (
	"fmt"
	"github.com/apmath-web/credit-go/viewModels"
	"net/http"
)

func Create(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type",
		"application/json; charset=utf-8")
	jsonData := toJson(response, request)
	if jsonData != nil {
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
	fmt.Fprintf(response, "{\"id\":1}")
}
