package actions

import (
	"net/http"
)

func Delete(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type",
		"application/json; charset=utf-8")
	id, err := getId(request.URL.Path)
	if err != nil {
		errorMessage("Invalid id format", 400, response)
		return
	}
	repo := Repository
	credit := repo.Get(id)
	if credit == nil {
		errorMessage("Credit not found", 404, response)
		return
	}
	if !credit.IsFinished() {
		errorMessage("Credit not paid in full", 400, response)
		return
	}
	repo.Remove(credit)
	response.WriteHeader(204)
}
