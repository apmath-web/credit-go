package viewModels

import (
	"github.com/apmath-web/credit-go/valueObjects"
	"github.com/franela/goblin"
	"strconv"
	"testing"
)

func TestViewModel(t *testing.T) {
	g := goblin.Goblin(t)
	type TestData struct {
		jsonData     map[string]interface{}
		validMessage *valueObjects.Validation
	}
	var jsonObjectTest map[string]interface{}
	var validationObjectTest *valueObjects.Validation
	var testViewModel *viewModel
	var negativeTestData, positiveTestData []TestData
	var numOfTestsPos, numOfTestsNeg int
	g.Describe("View model tests", func() {
		g.Before(func() {
			negativeTestData = []TestData{}
			positiveTestData = []TestData{
				{map[string]interface{}{"number": 0.2, "bool": false, "string": "str"},
					&valueObjects.Validation{}},
				{map[string]interface{}{"number": 1.0, "bool": true, "string": "str many words"},
					&valueObjects.Validation{}},
				{map[string]interface{}{"number": 23.0, "bool": true, "string": ""},
					&valueObjects.Validation{}},
				{map[string]interface{}{"number": -12.0, "bool": false, "string": "true"},
					&valueObjects.Validation{}},
				{map[string]interface{}{"number": -100000000.2, "bool": true, "string": "435"},
					&valueObjects.Validation{}},
				{map[string]interface{}{"number": 0.0, "bool": false, "string": "s!@326854asjdshb"},
					&valueObjects.Validation{}}}
			numOfTestsNeg = len(negativeTestData)
			numOfTestsPos = len(positiveTestData)
		})
		g.It("Positive tests", func() {
			for i := 0; i < numOfTestsPos; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						jsonObjectTest = positiveTestData[i].jsonData
						validationObjectTest = positiveTestData[i].validMessage
						testViewModel = new(viewModel)
					})
					g.It("View model fill", func() {
						testViewModel.Fill(jsonObjectTest)
						g.Assert(testViewModel.JsonData).Equal(jsonObjectTest)
					})
					g.It("View model check number", func() {
						exp := jsonObjectTest["number"]
						res := testViewModel.check("float64", "number")
						g.Assert(res).Equal(exp)

					})
					g.It("View model check bool", func() {
						exp := jsonObjectTest["bool"].(bool)
						res := testViewModel.check("bool", "bool")
						g.Assert(res).Equal(exp)
					})
					g.It("View model check string", func() {
						exp := jsonObjectTest["string"]
						res := testViewModel.check("string", "string")
						g.Assert(res).Equal(exp)
					})
					g.It("View model get validation", func() {
						g.Assert(validationObjectTest).Equal(testViewModel.GetValidation())
					})
				})
			}
		})
		g.It("Negative tests", func() {
			for i := 0; i < numOfTestsNeg; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {

				})
			}
		})
	})
}
