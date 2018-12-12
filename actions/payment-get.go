package actions

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func GetPayments(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type",
		"application/json; charset=utf-8")
	paths := strings.Split(request.URL.Path, "/")
	id, err := strconv.ParseInt(paths[2], 10, 64)
	if err != nil {
		errorMessage("Invalid id format", 400, response)
		return
	}
	type_ := getParam(request, "type")
	if err := validateParam(type_, []string{"regular", "early", "next"}, true); err != nil {
		errorMessage("Type: "+err.Error(), 400, response)
		return
	}
	state := getParam(request, "state")
	if err := validateParam(state, []string{"paid", "upcoming"}, true); err != nil {
		errorMessage("State: "+err.Error(), 400, response)
		return
	}
	repo := Repository
	credit := repo.Get(int(id))
	if credit == nil {
		errorMessage("Credit not found", 404, response)
		return
	}
	fmt.Fprint(response, "{\"payments\":[{\"type\":\"regular\",\"state\":\"paid\",\"date\":\"2018-10-08\","+
		"\"payment\":22300,\"percent\":10000,\"body\":12299,\"remainCreditBody\":907704,\"fullEarlyRepayment\":908704}]}")

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
