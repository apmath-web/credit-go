package actions

import (
	"github.com/apmath-web/credit-go/repositories"
	"net/http"
	"regexp"
)

func Handle(response http.ResponseWriter, request *http.Request) {
	url := request.URL
	path := []byte(url.Path)
	if request.Method == "POST" && validCredit.Match(path) {
		Create(response, request)
		return
	}
	if request.Method == "GET" && validCreditId.Match(path) {
		Get(response, request)
		return
	}
	if request.Method == "PUT" && validCreditId.Match(path) {
		PaymentWriteOf(response, request)
		return
	}
	if request.Method == "GET" && validPayments.Match(path) {
		GetPayments(response, request)
		return
	}
	if request.Method == "DELETE" && validCreditId.Match(path) {
		Delete(response, request)
		return
	}
	errorMessage("Page not found.", 404, response)
	// Todo add some header and more information about 404 error
	// fetch and display errors here
}

var Repository = repositories.GenRepository()

var validCredit = regexp.MustCompile("^/credit$")
var validCreditId = regexp.MustCompile("^/credit/[0-9]+$")
var validPayments = regexp.MustCompile("^/credit/[0-9]+/payments$")
