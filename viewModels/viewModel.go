package viewModels

import (
	"github.com/apmath-web/credit-go/valueObjects"
	"reflect"
)

type viewModel struct {
	JsonData      map[string]interface{}
	validMessages valueObjects.Validation
}

func (v *viewModel) check(type_ string, name string) interface{} {
	if val, ok := v.JsonData[name]; ok && val == nil {
		v.validMessages.AddMessage(
			valueObjects.GenMessage(name, "Is empty."))
		return nil
	}
	if val, ok := v.JsonData[name]; ok && val != nil && reflect.TypeOf(val).String() == type_ {
		return val
	} else {
		if ok {
			if type_ == "float64" {
				type_ = "integer number"
			}
			v.validMessages.AddMessage(
				valueObjects.GenMessage(name, "Must be "+type_+"."))
		} else {
			v.validMessages.AddMessage(
				valueObjects.GenMessage(name, "No field."))
		}
		return nil
	}
}

func (v *viewModel) Fill(jsonData map[string]interface{}) {
	v.JsonData = jsonData
}

func (v *viewModel) GetValidation() valueObjects.ValidationInterface {
	return &v.validMessages
}
