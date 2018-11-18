package tests_test

import (
	"github.com/apmath-web/credit-go/valueObjects"
	"testing"
)

func TestMessage(t *testing.T) {
	total := valueObjects.GenMessage("Field", "Text")
	if total.GetField() != "Field" {
		t.Errorf("Message doesn't save its field. Got: '%s'. Want %s.",
			total.GetField(), "Field")
	}
	if total.GetText() != "Text" {
		t.Errorf("Message doesn't save its text. Got: '%s'. Want %s.",
			total.GetText(), "Text")
	}
}
