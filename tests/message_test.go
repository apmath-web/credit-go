package main

import (
	"github.com/apmath-web/credit-go/valueObjects"
	"testing"
)

func TestMessage(t *testing.T) {
	total := new(valueObjects.Message)
	total.Message("Field", "Text")
	if total.GetField() != "Field" {
		t.Errorf("Message doesn't save its field. Got: '%s'. Want %s.", total.GetField(), "Field")
	}
	if total.GetText() != "Text" {
		t.Errorf("Message doesn't save its text. Got: '%s'. Want %s.", total.GetText(), "Text")
	}
}
