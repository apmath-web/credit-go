package main

import (
	"github.com/apmath-web/credit-go/actions"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", actions.Handle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
