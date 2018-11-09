package main

import (
	"github.com/apmath-web/credit-go/valueObjects"
	"reflect"
	"testing"
)

func TestValidation(t *testing.T) {
	total := valueObjects.Validation{}
	message1 := valueObjects.Message{}.Message("field1", "text1")
	message2 := valueObjects.Message{}.Message("field2", "text2")
	ans1 := []valueObjects.MessageInterface{message1, message2}
	total.AddMessages(ans1)
	if reflect.DeepEqual(ans1, total.GetMessages()) {
		t.Errorf("Don't save any messages. Got: %+v. Want %+v.", total.GetMessages(), ans1)
	}
	ans2 := []valueObjects.MessageInterface{message1, message2, message1, message2}
	if reflect.DeepEqual(ans2, total.GetMessages()) {
		t.Errorf("Don't append any messages. Got:\n%+v. \nWant\n%+v.", total.GetMessages(), ans2)
	}
}
