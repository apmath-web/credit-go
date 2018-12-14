package viewModels

import (
	"github.com/apmath-web/credit-go/valueObjects"
	"github.com/franela/goblin"
	"strconv"
	"testing"
)

func TestPersonViewModel(t *testing.T) {
	g := goblin.Goblin(t)
	type TestData struct {
		jsonData    map[string]interface{}
		fnError     int
		lnError     int
		valueObject valueObjects.PersonInterface
	}
	var jsonObjectTest map[string]interface{}
	var valueObjectTest valueObjects.PersonInterface
	var testPersonViewModel *Person
	var firstName, lastName string
	var ln, fn int
	var numOfTestsPos, numOfTestsNeg int
	var negativeTestData, positiveTestData []TestData
	g.Describe("Person view model tests", func() {
		g.Before(func() {
			negativeTestData = []TestData{
				TestData{map[string]interface{}{"firstName": "", "lastName": "lastname"},
					1, 0, nil},
				TestData{map[string]interface{}{"firstName": nil, "lastName": "lastname"},
					1, 0, nil},
				TestData{map[string]interface{}{"firstName": true, "lastName": "lastname"},
					1, 0, nil},
				TestData{map[string]interface{}{"firstName": 1, "lastName": "lastname"},
					1, 0, nil},
				TestData{map[string]interface{}{"firstName": 0, "lastName": "lastname"},
					1, 0, nil},
				TestData{map[string]interface{}{"firstName": "firstname", "lastName": ""},
					0, 1, nil},
				TestData{map[string]interface{}{"firstName": "firstname", "lastName": nil},
					0, 1, nil},
				TestData{map[string]interface{}{"firstName": "firstname", "lastName": true},
					0, 1, nil},
				TestData{map[string]interface{}{"firstName": "firstname", "lastName": 1},
					0, 1, nil},
				TestData{map[string]interface{}{"firstName": "firstname", "lastName": 0},
					0, 1, nil},
			}
			positiveTestData = []TestData{
				TestData{
					map[string]interface{}{"firstName": "firstname", "lastName": "lastname"}, 0, 0,
					valueObjects.GenPerson("firstname", "lastname")},
				TestData{
					map[string]interface{}{"firstName": "FIRSTNAME", "lastName": "LASTNAME"}, 0, 0,
					valueObjects.GenPerson("FIRSTNAME", "LASTNAME")},
				TestData{
					map[string]interface{}{"firstName": "FiRsTnAmE", "lastName": "LaStNaMe"}, 0, 0,
					valueObjects.GenPerson("FiRsTnAmE", "LaStNaMe")},
				TestData{
					map[string]interface{}{"firstName": "first name", "lastName": "first name"}, 0, 0,
					valueObjects.GenPerson("first name", "first name")},
				TestData{
					map[string]interface{}{"firstName": "f i r s t n a m e", "lastName": "l a s t n a m e"}, 0, 0,
					valueObjects.GenPerson("f i r s t n a m e", "l a s t n a m e")},
				TestData{
					map[string]interface{}{"firstName": "O", "lastName": "O"}, 0, 0,
					valueObjects.GenPerson("O", "O")},
				TestData{
					map[string]interface{}{"firstName": "Alexandra", "lastName": "Chernyshova"}, 0, 0,
					valueObjects.GenPerson("Alexandra", "Chernyshova")},
				TestData{
					map[string]interface{}{"firstName": "S-y#m_o!s", "lastName": "(y^mLa^#"}, 0, 0,
					valueObjects.GenPerson("S-y#m_o!s", "(y^mLa^#")}}
			numOfTestsPos = len(positiveTestData)
			numOfTestsNeg = len(negativeTestData)
		})
		g.It("Positive tests", func() {
			for i := 0; i < numOfTestsPos; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						valueObjectTest = positiveTestData[i].valueObject
						jsonObjectTest = positiveTestData[i].jsonData
						ln = positiveTestData[i].lnError
						fn = positiveTestData[i].fnError
					})
					g.Describe("#Person view model from request", func() {
						g.Before(func() {
							firstName = jsonObjectTest["firstName"].(string)
							lastName = jsonObjectTest["lastName"].(string)
							testPersonViewModel = new(Person)
							testPersonViewModel.Fill(jsonObjectTest)

						})
						g.Describe("##Person view model validate", func() {
							g.It("validate first name", func() {
								testPersonViewModel.validateFirstName()
								g.Assert(len(testPersonViewModel.
									validMessages.GetMessages())).Equal(0)
							})
							g.It("validate last name", func() {
								testPersonViewModel.validateLastName()
								g.Assert(len(testPersonViewModel.
									validMessages.GetMessages())).Equal(0)
							})
							g.It("validation is correct", func() {
								res := testPersonViewModel.Validate()
								g.Assert(res).IsTrue()
							})
						})
						g.Describe("##Person view model getters", func() {
							g.It("get first name", func() {
								res := testPersonViewModel.GetFirstName()
								g.Assert(res).Equal(firstName)
							})
							g.It("get last name", func() {
								res := testPersonViewModel.GetLastName()
								g.Assert(res).Equal(lastName)
							})
						})
					})
					g.Describe("#Person view model from model", func() {
						g.Before(func() {
							testPersonViewModel = new(Person)
							firstName = valueObjectTest.GetFirstName()
							lastName = valueObjectTest.GetLastName()
						})
						g.Describe("##Person view model hydrate", func() {
							g.It("correct hydrate from person value object", func() {
								testPersonViewModel.Hydrate(valueObjectTest)
								g.Assert(testPersonViewModel.FirstName).Equal(firstName)
								g.Assert(testPersonViewModel.LastName).Equal(lastName)
							})
						})
						g.Describe("##Person view model fetch", func() {
							g.It("correct fetch to json data", func() {
								res := testPersonViewModel.GetFirstName()
								g.Assert(res).Equal(firstName)
							})
						})
					})

				})
			}
		})
		g.It("Negative tests", func() {
			for i := 0; i < numOfTestsNeg; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						jsonObjectTest = negativeTestData[i].jsonData
						ln = negativeTestData[i].lnError
						fn = negativeTestData[i].fnError
					})
					g.Describe("#Person view model from request", func() {
						g.BeforeEach(func() {
							testPersonViewModel = new(Person)
							testPersonViewModel.Fill(jsonObjectTest)
						})
						g.Describe("##Person view model validate", func() {
							g.It("validate first name", func() {
								testPersonViewModel.validateFirstName()
								g.Assert(len(testPersonViewModel.
									validMessages.GetMessages())).Equal(fn)
							})
							g.It("validate last name", func() {
								testPersonViewModel.validateLastName()
								g.Assert(len(testPersonViewModel.
									validMessages.GetMessages())).Equal(ln)
							})
							g.It("validation is correct", func() {
								res := testPersonViewModel.Validate()
								g.Assert(res).IsFalse()
							})
						})

					})

				})
			}
		})

	})
}
