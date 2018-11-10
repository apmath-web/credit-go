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
	creditViewModel := new(viewModels.Credit)
	ok, _ := creditViewModel.Fill(request)
	fmt.Println(ok)
	if !ok {
		encoder := json.NewEncoder(response)
		err := encoder.Encode(valueObjects.GenMessageInArray("Json", "Invalid json"))
		if err != nil {
			log.Fatal(err)
		}
	}
	ok = creditViewModel.Validate()
	if !ok {
		encoder := json.NewEncoder(response)
		err := encoder.Encode(creditViewModel.GetValidation().GetMessages())
		if err != nil {
			log.Fatal(err)
		}
	}
	response.Header().Add("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(response, "{\"id\":1}")
}
