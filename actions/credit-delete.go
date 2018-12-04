package actions

import (
	"net/http"
	"strconv"
	"strings"
)

func Delete(response http.ResponseWriter, request *http.Request) {
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
	if credit.GetRemainAmount().Mon2Int64() != 0 {
		errorMessage("Credit not paid in full", 400, response)
		return
	}
	repo.Remove(credit)
	response.WriteHeader(204)
}
