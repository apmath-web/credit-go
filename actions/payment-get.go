package actions

import (
	"encoding/json"
	"fmt"
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/viewModels"
	"log"
	"net/http"
)

func GetPayments(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type",
		"application/json; charset=utf-8")
	id, err := getId(request.URL.Path)
	if err != nil {
		errorMessage("Invalid id format", 400, response)
		return
	}
	type_ := getParam(request, "type")
	if err := validateParam(type_, []string{"regular", "early", "next"}, true); err != nil {
		errorMessage("Type: "+err.Error(), 400, response)
		return
	}
	state := getParam(request, "state")
	if err := validateParam(state, []string{"paid", "upcoming"}, true); err != nil {
		errorMessage("State: "+err.Error(), 400, response)
		return
	}
	repo := Repository
	credit := repo.Get(id)
	if credit == nil {
		errorMessage("Credit not found", 404, response)
		return
	}
	payments := credit.GetPayments(data.Type(type_), data.State(state))
	var answer []interface{}
	for _, payment := range payments {
		viewModelPayment := new(viewModels.Payment)
		viewModelPayment.Hydrate(payment)
		jsonPayment := viewModelPayment.Fetch()
		answer = append(answer, jsonPayment)
	}
	jsonData := make(map[string][]interface{})
	jsonData["payments"] = answer
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatal(err.Error())
		errorMessage(err.Error(), 500, response)
		return
	}
	fmt.Println(string(jsonBytes[:]))
	fmt.Fprint(response, "{\"payments\":[{\"type\":\"regular\",\"state\":\"paid\",\"date\":\"2018-10-08\","+
		"\"payment\":22300,\"percent\":10000,\"body\":12299,\"remainCreditBody\":907704,\"fullEarlyRepayment\":908704}]}")
}
