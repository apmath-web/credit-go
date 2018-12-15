package viewModels

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/models"
	"github.com/apmath-web/credit-go/valueObjects"
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
	var valueObjectTest models.CreditInterface
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
			negativeTestData = []TestData{
				{map[string]interface{}{
					"person":      map[string]interface{}{"firstName": "FName", "lastName": "dfs"},
					"amount":      0.0,
					"agreementAt": "2017-12-10",
					"currency":    "RUR",
					"duration":    5.0,
					"percent":     0.0,
				},
					0, 1, 1, 0, 1, 0,
					nil,
				},
				{map[string]interface{}{
					"person":      map[string]interface{}{"firstName": "FName", "lastName": "dfs"},
					"amount":      3000000000000001.0,
					"agreementAt": "2017-12-10",
					"currency":    "RUR",
					"duration":    1201.0,
					"percent":     301.0,
				},
					0, 1, 1, 0, 1, 0,
					nil,
				},
				{map[string]interface{}{
					"person":      map[string]interface{}{"firstName": nil, "lastName": nil},
					"agreementAt": "2017-12-10",
					"currency":    "RUR",
					"duration":    12.0,
					"percent":     10.0,
				},
					0, 0, 0, 2, 1, 0,
					nil,
				},
				{map[string]interface{}{
					"person":      map[string]interface{}{"firstName": "FName", "lastName": "dfs"},
					"amount":      2000.0,
					"agreementAt": "2017-12-10",
					"duration":    12.0,
					"percent":     10.0,
				},
					1, 0, 0, 0, 0, 0,
					nil,
				},
				{map[string]interface{}{
					"person":      map[string]interface{}{"firstName": "FName", "lastName": "dfs"},
					"amount":      2000.0,
					"agreementAt": "2017-12-10",
					"currency":    "RUR",
				},
					0, 1, 1, 0, 0, 0,
					nil,
				},
				{map[string]interface{}{
					"person":      map[string]interface{}{"firstName": "FName", "lastName": "dfs"},
					"amount":      -2000.0,
					"agreementAt": "2017-12-10",
					"currency":    "RUR",
					"duration":    -12.0,
					"percent":     -10.0,
				},

					0, 1, 1, 0, 1, 0,
					nil,
				},
				{map[string]interface{}{
					"person":      map[string]interface{}{"firstName": "FName", "lastName": "dfs"},
					"amount":      -2000.0,
					"agreementAt": "2017-12-10",
					"currency":    "RUR",
					"percent":     -10.0,
				},

					0, 1, 1, 0, 1, 0,
					nil,
				},
				{map[string]interface{}{
					"person":      map[string]interface{}{"firstName": "FName", "lastName": "dfs"},
					"amount":      -2000.0,
					"agreementAt": "2017-12-10",
					"currency":    "RUR",
					"duration":    -12.0,
				},

					0, 1, 1, 0, 1, 0,
					nil,
				},
				{map[string]interface{}{
					"person":      map[string]interface{}{"firstName": "FName", "lastName": "dfs"},
					"amount":      2000.0,
					"agreementAt": "2009/21/02",
					"currency":    "RUR",
					"duration":    12.0,
					"percent":     10.0,
				},
					0, 0, 0, 0, 0, 1,
					nil,
				},
			}
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
					"amount":   600.0,
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
			positiveTestData[0].valueObject, _ = models.GenCredit(valueObjects.GenPerson("FName", "dfs"),
				2000, data.Str2Date("2017-12-10"), "RUR", 12, 10)
			positiveTestData[1].valueObject, _ = models.GenCredit(valueObjects.GenPerson("FName", "dfs"),
				600, data.Str2Date(data.Date(time.Now()).Date2Str()), "EUR", 6, 1)
			positiveTestData[2].valueObject, _ = models.GenCredit(valueObjects.GenPerson("FName", "dfs"),
				3000000000000000, data.Str2Date(data.Date(time.Now()).Date2Str()), "USD", 1200, 300)
			numOfTestsNeg = len(negativeTestData)
			numOfTestsPos = len(positiveTestData)
		})
		g.It("Positive tests", func() {
			for i := 0; i < numOfTestsPos; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						jsonObjectTest = positiveTestData[i].jsonData
						valueObjectTest = positiveTestData[i].valueObject
						c = positiveTestData[i].currencyError
						d = positiveTestData[i].durationError
						am = positiveTestData[i].amountError
						ag = positiveTestData[i].agrementError
						ps = positiveTestData[i].personError
						pc = positiveTestData[i].percentError
					})
					g.Describe("#Credit view model from request", func() {
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
						g.Describe("##Credit view model validate", func() {
							g.It("validate agreement at", func() {
								testCreditViewModel.validateAgreementAt()
								g.Assert(len(testCreditViewModel.
									validMessages.GetMessages())).Equal(ag)
							})
							g.It("validate amount", func() {
								testCreditViewModel.validateAmount()
								g.Assert(len(testCreditViewModel.
									validMessages.GetMessages())).Equal(am)
							})
							g.It("validate currency", func() {
								testCreditViewModel.validateCurrency()
								g.Assert(len(testCreditViewModel.
									validMessages.GetMessages())).Equal(c)
							})
							g.It("validate duration", func() {
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
					g.Describe("#Credit view model from value object", func() {
						g.Before(func() {
							testCreditViewModel = new(Credit)
							currency = valueObjectTest.GetCurrency().Cur2Str()
							amount = valueObjectTest.GetAmount().Mon2Int64()
							agreementAt = valueObjectTest.GetAgreementAt().Date2Str()
							percent = valueObjectTest.GetPercent()
							duration = valueObjectTest.GetDuration()
							person = new(Person)
							person.Hydrate(valueObjectTest.GetPerson())
						})
						g.Describe("##Credit view model hydrate", func() {
							g.It("correct hydrate from person value object", func() {
								testCreditViewModel.Hydrate(valueObjectTest)
								g.Assert(testCreditViewModel.GetPerson()).Equal(person)
								g.Assert(testCreditViewModel.Amount).Equal(amount)
								g.Assert(testCreditViewModel.AgreementAt).Equal(agreementAt)
								g.Assert(testCreditViewModel.Currency).Equal(currency)
								g.Assert(testCreditViewModel.Duration).Equal(duration)
								g.Assert(testCreditViewModel.Percent).Equal(percent)
							})
						})
						g.Describe("##Crdit view model fetch", func() {
							g.It("correct fetch to json data", func() {
								res := testCreditViewModel.Fetch().(map[string]interface{})
								res["amount"] = float64(res["amount"].(int64))
								res["duration"] = float64(res["duration"].(int32))
								res["percent"] = float64(res["percent"].(int32))
								g.Assert(res).Equal(jsonObjectTest)
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
						c = negativeTestData[i].currencyError
						d = negativeTestData[i].durationError
						am = negativeTestData[i].amountError
						ag = negativeTestData[i].agrementError
						ps = negativeTestData[i].personError
						pc = negativeTestData[i].percentError
					})
					g.BeforeEach(func() {
						testCreditViewModel = new(Credit)
						testCreditViewModel.Fill(jsonObjectTest)
					})
					g.Describe("#Credit view model validate", func() {
						g.It("validate agreement at", func() {
							testCreditViewModel.validateAgreementAt()
							g.Assert(len(testCreditViewModel.
								validMessages.GetMessages())).Equal(ag)
						})
						g.It("validate amount", func() {
							testCreditViewModel.validateAmount()
							g.Assert(len(testCreditViewModel.
								validMessages.GetMessages())).Equal(am)
						})
						g.It("validate currency", func() {
							testCreditViewModel.validateCurrency()
							g.Assert(len(testCreditViewModel.
								validMessages.GetMessages())).Equal(c)
						})
						g.It("validate duration", func() {
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
