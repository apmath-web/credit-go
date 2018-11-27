package actions

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Get(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type",
		"application/json; charset=utf-8")
	paths := strings.Split(request.URL.Path, "/credit/")
	fmt.Println(paths)
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
	fmt.Fprintf(response, "{\"id\": %d }", credit.GetId())
}
