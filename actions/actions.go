package actions

import (
	"net/http"
)

func Handle(response http.ResponseWriter, request *http.Request) {
	url := request.URL
	if url.Path == "/credit" && request.Method == "POST" {
		Create(response, request)
	}
	// fetch and display errors here
}
