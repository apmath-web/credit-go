package actions

import (
	"encoding/json"
	"fmt"
	"github.com/apmath-web/credit-go/repositories"
	"github.com/apmath-web/credit-go/valueObjects"
	"net/http"
)

func Handle(response http.ResponseWriter, request *http.Request) {
	url := request.URL
	if url.Path == "/credit" && request.Method == "POST" {
		Create(response, request)
	}
	// fetch and display errors here
}

func toJson(response http.ResponseWriter, request *http.Request) map[string]interface{} {
	decoder := json.NewDecoder(request.Body)
	var jsonData map[string]interface{}
	err := decoder.Decode(&jsonData)
	if err != nil {
		jsonData := ptrMessagesToJsonErrMessage("Validation error",
			[]valueObjects.MessageInterface{
				valueObjects.GenMessage("package", err.Error())})
		response.WriteHeader(400)
		fmt.Fprint(response, jsonData)
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
		return "{\"message\":\"some error on server\"}"
	}
	return string(res)
}

var Repository = repositories.GenRepository()
