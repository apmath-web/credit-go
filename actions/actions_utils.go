package actions

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apmath-web/credit-go/valueObjects"
	"log"
	"net/http"
	"strconv"
	"strings"
)

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

func getParam(request *http.Request, param string) (string, bool) {
	var key string
	keys, ok := request.URL.Query()[param]
	if !ok || len(keys[0]) < 1 {
		key = ""
	} else {
		key = keys[0]
	}
	return key, ok
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

func validateParam(param string, values []string, isnull bool) error {
	if param == "" {
		if isnull {
			return nil
		}
		return errors.New("is empty")
	} else {
		for _, val := range values {
			if val == param {
				return nil
			}
		}
		return errors.New(param + " is unknown value.")
	}
}

func getId(s string) (int, error) {
	path := strings.Split(s, "/")
	return strconv.Atoi(path[2])
}
