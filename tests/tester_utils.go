package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func GenerateRequest(Json string) *http.Request {
	req, _ := http.NewRequest("Post", "/credit",
		ioutil.NopCloser(bytes.NewBufferString(Json)))

	return req
}
