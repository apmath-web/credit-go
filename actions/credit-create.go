package actions

import (
	"encoding/json"
	"fmt"
	"github.com/apmath-web/credit-go/valueObjects"
	"github.com/apmath-web/credit-go/viewModels"
	"log"
	"net/http"
)

func Create(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json; charset=utf-8")
	creditViewModel := new(viewModels.Credit)
	ok, _ := creditViewModel.Fill(request)
	if !ok {
		fmt.Fprintf(response, "messages:[{\"Package\":\"Invalid json format\"}]")
		return
	}
	ok = creditViewModel.Validate()
	if !ok {
		jsonData := PtrMessagesToJson(creditViewModel.GetValidation().GetMessages())
		fmt.Fprint(response, jsonData)
		return
	}
	response.Header().Add("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(response, "{\"id\":1}")
}

func PtrMessagesToJson(messagesPtr []valueObjects.MessageInterface) string {
	type message struct {
		Field string `json:"field"`
		Text  string `json:"text"`
	}
	messages := []message{}
	for _, b := range messagesPtr {
		messages = append(messages, message{Field: b.GetField(), Text: b.GetText()})
	}
	jsonData, err := json.Marshal(&messages)
	if err != nil {
		log.Fatalf("%+v", err)
		return "{}"
	}
	return string(jsonData)
}
