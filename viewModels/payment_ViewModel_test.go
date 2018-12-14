package viewModels

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/valueObjects"
	"github.com/franela/goblin"
	"strconv"
	"testing"
	"time"
)

func TestPaymentViewModel(t *testing.T) {
	g := goblin.Goblin(t)
	type TestData struct {
		jsonData    map[string]interface{}
		cError      int
		dError      int
		tError      int
		pError      int
		valueObject valueObjects.PaymentInterface
	}
	var jsonObjectTest map[string]interface{}
	//var valueObjectTest valueObjects.PaymentInterface
	var testPaymentViewModel *Payment
	var negativeTestData, positiveTestData []TestData
	var numOfTestsPos, numOfTestsNeg int
	var date, currency, type_ string
	var paymentAmount int64
	var c, d, ty, p int
	g.Describe("View model tests", func() {
		g.Before(func() {
			negativeTestData = []TestData{
				{map[string]interface{}{"payment": 0.0,
					"currency": "RUR",
					"date":     "2018-01-01",
					"type":     "regular"},
					0, 0, 0, 1,
					nil},
				{map[string]interface{}{
					"currency": "RUB",
					"date":     "none",
					"type":     "regular"},
					1, 1, 0, 1,
					nil},
				{map[string]interface{}{"payment": -1324.0,
					"currency": "RUR",
					"date":     "2018/01/01",
					"type":     "regular"},
					0, 1, 0, 1,
					nil},
				{map[string]interface{}{"payment": 3750000000000001.0,
					"currency": "",
					"date":     "2018-54-01",
					"type":     "regular"},
					1, 1, 0, 1,
					nil},
				{map[string]interface{}{"payment": 3750000000000000.0,
					"currency": "RUR",
					"date":     "000-01-01",
					"type":     "kerbfjd"},
					0, 1, 1, 0,
					nil},
				{map[string]interface{}{"payment": 1.0,
					"currency": "non",
					"date":     "2018-01-01",
					"type":     2113178.5},
					1, 0, 1, 0,
					nil},
				{map[string]interface{}{"payment": 23532.0,
					"currency": "EUR",
					"date":     "01.02.2029",
					"type":     "regular"},
					0, 1, 0, 0,
					nil},
				{map[string]interface{}{"payment": 5686.0,
					"currency": "RUR",
					"date":     true,
					"type":     "next"},
					0, 1, 1, 0,
					nil},
			}
			positiveTestData = []TestData{
				{map[string]interface{}{"payment": 2955.0,
					"currency": "RUR",
					"date":     "2018-01-01",
					"type":     "regular"},
					0, 0, 0, 0,
					nil},
				{map[string]interface{}{"payment": 29324225.0,
					"currency": "USD",
					"date":     "2018-10-21",
					"type":     "early"},
					0, 0, 0, 0,
					nil},
				{map[string]interface{}{"payment": 1.0,
					"currency": "RUR",
					"date":     "",
					"type":     "regular"},
					0, 0, 0, 0,
					nil},
				{map[string]interface{}{"payment": 214.0,
					"currency": "EUR",
					"date":     "2006-01-02",
					"type":     ""},
					0, 0, 0, 0,
					nil},
				{map[string]interface{}{"payment": 234454532.0,
					"currency": "RUR",
					"date":     "",
					"type":     ""},
					0, 0, 0, 0,
					nil},
				{map[string]interface{}{"payment": 23512.0,
					"currency": "USD",
					"date":     "1970-01-01",
					"type":     "early"},
					0, 0, 0, 0,
					nil},
				{map[string]interface{}{"payment": 8436296234.02,
					"currency": "USD",
					"date":     "",
					"type":     "early"},
					0, 0, 0, 0,
					nil},
				{map[string]interface{}{"payment": 2955.0,
					"currency": "RUR",
					"date":     "2018-01-01"},
					0, 0, 0, 0,
					nil},
				{map[string]interface{}{"payment": 2955.0,
					"currency": "RUR",
					"type":     "regular"},
					0, 0, 0, 0,
					nil},
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
						c = positiveTestData[i].cError
						d = positiveTestData[i].dError
						ty = positiveTestData[i].tError
						p = positiveTestData[i].pError
					})
					g.Describe("#Payment view model from request", func() {
						g.Before(func() {
							currency = jsonObjectTest["currency"].(string)
							date_tmp, ok := jsonObjectTest["date"].(string)
							date = date_tmp
							if date == "" || !ok {
								date = data.Date(time.Now()).Date2Str()
							}
							type_, ok = jsonObjectTest["type"].(string)
							if !ok {
								type_ = ""
							}
							paymentAmount = int64(jsonObjectTest["payment"].(float64))
							testPaymentViewModel = new(Payment)
							testPaymentViewModel.Fill(jsonObjectTest)

						})
						g.Describe("##Payment view model validate", func() {
							g.It("validate date", func() {
								testPaymentViewModel.validateDate()
								g.Assert(len(testPaymentViewModel.
									validMessages.GetMessages())).Equal(d)
							})
							g.It("validate payment", func() {
								testPaymentViewModel.validatePayment()
								g.Assert(len(testPaymentViewModel.
									validMessages.GetMessages())).Equal(p)
							})
							g.It("validate currency", func() {
								testPaymentViewModel.validateCurrency()
								g.Assert(len(testPaymentViewModel.
									validMessages.GetMessages())).Equal(c)
							})
							g.It("validate type", func() {
								testPaymentViewModel.validateType()
								g.Assert(len(testPaymentViewModel.
									validMessages.GetMessages())).Equal(ty)
							})
							g.It("validation is correct", func() {
								res := testPaymentViewModel.Validate()
								g.Assert(res).IsTrue()
							})
						})
						g.Describe("##Payment view model getters", func() {
							g.It("get currency", func() {
								res := testPaymentViewModel.GetCurrency()
								g.Assert(res).Equal(data.Str2Cur(currency))
							})
							g.It("get type", func() {
								res := testPaymentViewModel.GetType()
								g.Assert(res).Equal(data.Str2Type(type_))
							})
							g.It("get payment", func() {
								res := testPaymentViewModel.GetPayment()
								g.Assert(res).Equal(data.Money(paymentAmount))
							})
							g.It("get date", func() {
								res := testPaymentViewModel.GetDate()
								g.Assert(res).Equal(data.Str2Date(date))
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
						//valueObjectTest = negativeTestData[i].valueObject
						c = negativeTestData[i].cError
						d = negativeTestData[i].dError
						ty = negativeTestData[i].tError
						p = negativeTestData[i].pError
					})
					g.BeforeEach(func() {
						testPaymentViewModel = new(Payment)
						testPaymentViewModel.Fill(jsonObjectTest)
					})
					g.Describe("##Payment view model validate", func() {
						g.It("validate date", func() {
							testPaymentViewModel.validateDate()
							g.Assert(len(testPaymentViewModel.
								validMessages.GetMessages())).Equal(d)
						})
						g.It("validate payment", func() {
							testPaymentViewModel.validatePayment()
							g.Assert(len(testPaymentViewModel.
								validMessages.GetMessages())).Equal(p)
						})
						g.It("validate currency", func() {
							testPaymentViewModel.validateCurrency()
							g.Assert(len(testPaymentViewModel.
								validMessages.GetMessages())).Equal(c)
						})
						g.It("validate type", func() {
							testPaymentViewModel.validateType()
							g.Assert(len(testPaymentViewModel.
								validMessages.GetMessages())).Equal(ty)
						})
						g.It("validation is correct", func() {
							res := testPaymentViewModel.Validate()
							g.Assert(res).IsFalse()
						})
					})
				})
			}
		})
	})
}
