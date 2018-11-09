package actions

import (
	"fmt"
	"net/http"
)

func Create(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(response, "{\"id\":1}")
}
