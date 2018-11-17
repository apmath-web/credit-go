package actions

import (
	"encoding/json"
	"fmt"
	"github.com/apmath-web/credit-go/valueObjects"
	"github.com/apmath-web/credit-go/viewModels"
	"net/http"
)

func Create(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type",
		"application/json; charset=utf-8")
	jsonData, ok := ToJson(response, request)
	if !ok {
		return
	}
	creditViewModel := new(viewModels.Credit)
	creditViewModel.Fill(jsonData)
	ok = creditViewModel.Validate()
	if !ok {
		jsonData := PtrMessagesToJsonErrMessage("Validation error",
			creditViewModel.GetValidation().GetMessages())
		response.WriteHeader(400)
		fmt.Fprint(response, jsonData)
		return
	}
	fmt.Fprintf(response, "{\"id\":1}")
}

func ToJson(response http.ResponseWriter, request *http.Request) (map[string]interface{}, bool) {
	decoder := json.NewDecoder(request.Body)
	var jsonData map[string]interface{}
	err := decoder.Decode(&jsonData)
	if err != nil {
		jsonData := PtrMessagesToJsonErrMessage("Validation error",
			[]valueObjects.MessageInterface{
				valueObjects.GenMessage("package", err.Error())})
		response.WriteHeader(400)
		fmt.Fprint(response, jsonData)
		return nil, false
	}
	return jsonData, true
}

func PtrMessagesToJsonErrMessage(message string,
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
