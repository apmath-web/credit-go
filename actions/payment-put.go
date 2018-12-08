package actions

import (
	"fmt"
	"github.com/apmath-web/credit-go/valueObjects"
	"github.com/apmath-web/credit-go/viewModels"
	"net/http"
	"strconv"
	"strings"
)

func PaymentWriteOf(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type",
		"application/json; charset=utf-8")
	paths := strings.Split(request.URL.Path, "/credit/")
	id, err := strconv.ParseInt(paths[1], 10, 64)
	if err != nil {
		errorMessage("Invalid id format", 400, response)
		return
	}
	repo := Repository
	credit := repo.Get(int(id))
	if credit == nil {
		errorMessage("Credit not found", 404, response)
		return
	}
	jsonData := toJson(response, request)
	if jsonData == nil {
		return
	}
	if val, ok := jsonData["currency"]; (ok && val == nil) || !ok {
		jsonData["currency"] = credit.GetCurrency().Cur2Str()
	} else {
		if jsonData["currency"] != credit.GetCurrency().Cur2Str() {
			jsonData := ptrMessagesToJsonErrMessage("Validation error",
				[]valueObjects.MessageInterface{valueObjects.
					GenMessage("currency", "Not same for credit currency")})
			response.WriteHeader(400)
			fmt.Fprint(response, jsonData)
			return
		}
	}
	paymentViewModel := new(viewModels.Payment)
	paymentViewModel.Fill(jsonData)
	ok := paymentViewModel.Validate()
	if !ok {
		jsonData := ptrMessagesToJsonErrMessage("Validation error",
			paymentViewModel.GetValidation().GetMessages())
		response.WriteHeader(400)
		fmt.Fprint(response, jsonData)
		return
	}
	fmt.Fprintf(response, "{\"paymentExecutedAt\":\"%s\"}", jsonData["date"])
}
