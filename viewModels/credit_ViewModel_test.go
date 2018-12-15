package viewModels

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/models"
	"github.com/franela/goblin"
	"strconv"
	"testing"
	"time"
)

func TestCreditViewModel(t *testing.T) {
	g := goblin.Goblin(t)
	type TestData struct {
		jsonData      map[string]interface{}
		currencyError int
		durationError int
		percentError  int
		personError   int
		amountError   int
		agrementError int
		valueObject   models.CreditInterface
	}
	var jsonObjectTest map[string]interface{}
	//var valueObjectTest models.CreditInterface
	var testCreditViewModel *Credit
	var negativeTestData, positiveTestData []TestData
	var numOfTestsPos, numOfTestsNeg int
	var agreementAt, currency string
	var amount int64
	var duration, percent int32
	var person *Person
	var c, d, am, ag, ps, pc int
	g.Describe("View model tests", func() {
		g.Before(func() {
			negativeTestData = []TestData{}
			positiveTestData = []TestData{
				{map[string]interface{}{
					"person":      map[string]interface{}{"firstName": "FName", "lastName": "dfs"},
					"amount":      2000.0,
					"agreementAt": "2017-12-10",
					"currency":    "RUR",
					"duration":    12.0,
					"percent":     10.0,
				},
					0, 0, 0, 0, 0, 0,
					nil,
				},
				{map[string]interface{}{
					"person":   map[string]interface{}{"firstName": "FName", "lastName": "dfs"},
					"amount":   1.0,
					"currency": "EUR",
					"duration": 6.0,
					"percent":  1.0,
				},
					0, 0, 0, 0, 0, 0,
					nil,
				},
				{map[string]interface{}{
					"person":      map[string]interface{}{"firstName": "FName", "lastName": "dfs"},
					"amount":      3000000000000000.0,
					"agreementAt": "",
					"currency":    "USD",
					"duration":    1200.0,
					"percent":     300.0,
				},
					0, 0, 0, 0, 0, 0,
					nil,
				},
			}
			numOfTestsNeg = len(negativeTestData)
			numOfTestsPos = len(positiveTestData)
		})
		g.It("Positive tests", func() {
			for i := 0; i < numOfTestsPos; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						jsonObjectTest = positiveTestData[i].jsonData
						//valueObjectTest = positiveTestData[i].valueObject
						c = positiveTestData[i].currencyError
						d = positiveTestData[i].durationError
						am = positiveTestData[i].amountError
						ag = positiveTestData[i].agrementError
						ps = positiveTestData[i].personError
						pc = positiveTestData[i].percentError
					})
					g.Describe("#Payment view model from request", func() {
						g.Before(func() {
							currency = jsonObjectTest["currency"].(string)
							date_tmp, ok := jsonObjectTest["agreementAt"].(string)
							agreementAt = date_tmp
							if agreementAt == "" || !ok {
								agreementAt = data.Date(time.Now()).Date2Str()
							}
							amount = int64(jsonObjectTest["amount"].(float64))
							percent = int32(jsonObjectTest["percent"].(float64))
							duration = int32(jsonObjectTest["duration"].(float64))
							person = new(Person)
							person.Fill(jsonObjectTest["person"].(map[string]interface{}))
							person.Validate()
							testCreditViewModel = new(Credit)
							testCreditViewModel.Fill(jsonObjectTest)

						})
						g.Describe("##Payment view model validate", func() {
							g.It("validate agreement at", func() {
								testCreditViewModel.validateAgreementAt()
								g.Assert(len(testCreditViewModel.
									validMessages.GetMessages())).Equal(ag)
							})
							g.It("validate payment", func() {
								testCreditViewModel.validateAmount()
								g.Assert(len(testCreditViewModel.
									validMessages.GetMessages())).Equal(am)
							})
							g.It("validate currency", func() {
								testCreditViewModel.validateCurrency()
								g.Assert(len(testCreditViewModel.
									validMessages.GetMessages())).Equal(c)
							})
							g.It("validate type", func() {
								testCreditViewModel.validateDuration()
								g.Assert(len(testCreditViewModel.
									validMessages.GetMessages())).Equal(d)
							})
							g.It("validate person", func() {
								testCreditViewModel.validatePerson()
								g.Assert(len(testCreditViewModel.
									validMessages.GetMessages())).Equal(ps)
							})
							g.It("validate percent", func() {
								testCreditViewModel.validatePercent()
								g.Assert(len(testCreditViewModel.
									validMessages.GetMessages())).Equal(pc)
							})
							g.It("validation is correct", func() {
								res := testCreditViewModel.Validate()
								g.Assert(res).IsTrue()
							})
						})
						g.Describe("##Payment view model getters", func() {
							g.It("get currency", func() {
								res := testCreditViewModel.GetCurrency()
								g.Assert(res).Equal(data.Str2Cur(currency))
							})
							g.It("get amount", func() {
								res := testCreditViewModel.GetAmount()
								g.Assert(res).Equal(data.Money(amount))
							})
							g.It("get percent", func() {
								res := testCreditViewModel.GetPercent()
								g.Assert(res).Equal(percent)
							})
							g.It("get duration", func() {
								res := testCreditViewModel.GetDuration()
								g.Assert(res).Equal(duration)
							})
							g.It("get agreement at", func() {
								res := testCreditViewModel.GetAgreementAt()
								g.Assert(res).Equal(data.Str2Date(agreementAt))
							})
							g.It("get person", func() {
								res := testCreditViewModel.GetPerson()
								g.Assert(res).Equal(person)
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
						c = positiveTestData[i].currencyError
						d = positiveTestData[i].durationError
						am = positiveTestData[i].amountError
						ag = positiveTestData[i].agrementError
						ps = positiveTestData[i].personError
						pc = positiveTestData[i].percentError
					})
					g.BeforeEach(func() {
						testCreditViewModel = new(Credit)
						testCreditViewModel.Fill(jsonObjectTest)
					})
					g.Describe("##Payment view model validate", func() {
						g.It("validate agreement at", func() {
							testCreditViewModel.validateAgreementAt()
							g.Assert(len(testCreditViewModel.
								validMessages.GetMessages())).Equal(ag)
						})
						g.It("validate payment", func() {
							testCreditViewModel.validateAmount()
							g.Assert(len(testCreditViewModel.
								validMessages.GetMessages())).Equal(am)
						})
						g.It("validate currency", func() {
							testCreditViewModel.validateCurrency()
							g.Assert(len(testCreditViewModel.
								validMessages.GetMessages())).Equal(c)
						})
						g.It("validate type", func() {
							testCreditViewModel.validateDuration()
							g.Assert(len(testCreditViewModel.
								validMessages.GetMessages())).Equal(d)
						})
						g.It("validate person", func() {
							testCreditViewModel.validatePerson()
							g.Assert(len(testCreditViewModel.
								validMessages.GetMessages())).Equal(ps)
						})
						g.It("validate percent", func() {
							testCreditViewModel.validatePercent()
							g.Assert(len(testCreditViewModel.
								validMessages.GetMessages())).Equal(pc)
						})
						g.It("validation is correct", func() {
							res := testCreditViewModel.Validate()
							g.Assert(res).IsFalse()
						})
					})
				})
			}
		})
	})
}
