package actions

import (
	"fmt"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "{id:1}")
}