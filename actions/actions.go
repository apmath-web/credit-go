package actions

import (
	"encoding/json"
	"fmt"
	"github.com/apmath-web/credit-go/repositories"
	"github.com/apmath-web/credit-go/valueObjects"
	"log"
	"net/http"
	"regexp"
)

func Handle(response http.ResponseWriter, request *http.Request) {
	url := request.URL
	path := []byte(url.Path)
	if validCredit.Match(path) && request.Method == "POST" {
		Create(response, request)
		return
	}
	if validCreditId.Match(path) && request.Method == "GET" {
		Get(response, request)
		return
	}
	if validCreditId.Match(path) && request.Method == "PUT" {
		PaymentWriteOf(response, request)
		return
	}
	if validPayments.Match(path) && request.Method == "GET" {
		fmt.Println(request.URL.Path)
		GetPayments(response, request)
		return
	}
	errorMessage("Page not found.", 404, response)
	// Todo add some header and more information about 404 error
	// fetch and display errors here
}

func toJson(response http.ResponseWriter, request *http.Request) map[string]interface{} {
	decoder := json.NewDecoder(request.Body)
	var jsonData map[string]interface{}
	err := decoder.Decode(&jsonData)
	if err != nil {
		errorMessage("Validation error", 400, response)
		return nil
	}
	return jsonData
}

func ptrMessagesToJsonErrMessage(message string,
	descriptionPtr []valueObjects.MessageInterface) string {
	if descriptionPtr == nil {
		return "{\"message\":\"" + message + "\"}"
	}
	var description map[string]interface{}
	description = make(map[string]interface{})
	for _, item := range descriptionPtr {
		description[item.GetField()] = item.GetText()
	}
	jsonData := map[string]interface{}{
		"message":     message,
		"description": description,
	}
	res, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatal(err.Error())
		return "{\"message\":\"some error on server\"}"
	}
	return string(res)
}

func errorMessage(message string, code int, response http.ResponseWriter) {
	response.WriteHeader(code)
	fmt.Fprintf(response, "{\"message\":\"%s\"}", message)
}

var Repository = repositories.GenRepository()

var validCredit = regexp.MustCompile("^/credit$")
var validCreditId = regexp.MustCompile("^/credit/[0-9]+$")
var validPayments = regexp.MustCompile("^/credit/[0-9]+/payments$")
