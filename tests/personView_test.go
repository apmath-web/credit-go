package main

import (
	"bytes"
	"encoding/json"
	"github.com/apmath-web/credit-go/valueObjects"
	"github.com/apmath-web/credit-go/viewModels"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestPersonViewCreation(t *testing.T) {
	req := ioutil.NopCloser(bytes.NewBufferString("{\"firstName\":\"aaa\"}"))
	decoder := json.NewDecoder(req)
	a := new(viewModels.Person)
	_ = decoder.Decode(a)
	if !a.Validate() {
		validator := a.GetValidation()
		messages := validator.GetMessages()
		message := messages[0]
		messageEx := new(valueObjects.Message)
		messageEx.Message("LastName", "Is empty.")
		if !reflect.DeepEqual(message, messageEx) {
			t.Errorf("Don't give any messages. Got: %+v. Want: %+v.", message, messageEx)
		}
	} else {
		t.Errorf("Isn't good validation")
	}

}
