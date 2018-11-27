package actions

import (
	"encoding/json"
	"fmt"
	"github.com/apmath-web/credit-go/viewModels"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Get(response http.ResponseWriter, request *http.Request) {
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
	creditViewModel := new(viewModels.Credit)
	creditViewModel.Hydrate(credit)
	jsonData := creditViewModel.Fetch()
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatal(err.Error())
		errorMessage(err.Error(), 500, response)
		return
	}
	fmt.Fprint(response, string(jsonBytes[:]))
}
