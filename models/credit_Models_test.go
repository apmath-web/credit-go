package models

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/valueObjects"
	"github.com/franela/goblin"
	"strconv"
	"testing"
)

func TestCreditModel(t *testing.T) {
	g := goblin.Goblin(t)
	type TestData struct {
		person         valueObjects.PersonInterface
		currency       data.Currency
		duration       int32
		percent        int32
		amount         data.Money
		agreementAt    data.Date
		rounding       int32
		regularPayment data.Money
		annuityPayment float64
		id             int
	}
	var testObject TestData
	var testCreditModel *Credit
	var negativeTestData, positiveTestData []TestData
	var numOfTestsPos, numOfTestsNeg int
	g.Describe("View model tests", func() {
		g.Before(func() {
			negativeTestData = []TestData{
				{valueObjects.GenPerson("Fname", "Lname"),
					data.Currency("RUR"), 1200, 1,
					data.Money(12001), data.Str2Date("2018-01-02"),
					100, data.Money(0), 15.824921440605804, -1},
				{valueObjects.GenPerson("Fname", "Lname"),
					data.Currency("RUR"), 120, 1,
					data.Money(12001), data.Str2Date("2018-01-02"),
					1, data.Money(0), 105.13370605633396, -1},
			}
			positiveTestData = []TestData{
				{valueObjects.GenPerson("Fname", "Lname"),
					data.Currency("RUR"), 6, 1,
					data.Money(120321), data.Str2Date("2018-01-02"),
					100, data.Money(20200), 20112.02997569273, -1},
				{valueObjects.GenPerson("Fname", "Lname"),
					data.Currency("RUR"), 200, 5,
					data.Money(35478), data.Str2Date("2018-01-02"),
					1, data.Money(262), 261.799938969009, -1},
				{valueObjects.GenPerson("Fname", "Lname"),
					data.Currency("RUR"), 15, 10,
					data.Money(136234), data.Str2Date("2018-01-02"),
					100, data.Money(9700), 9699.47259479209, -1},
				{valueObjects.GenPerson("Fname", "Lname"),
					data.Currency("RUR"), 6, 300,
					data.Money(13420321), data.Str2Date("2018-01-02"),
					100, data.Money(4547100), 4.54706643301674e+06, -1},
			}
			numOfTestsNeg = len(negativeTestData)
			numOfTestsPos = len(positiveTestData)
		})
		g.It("Positive tests", func() {
			for i := 0; i < numOfTestsPos; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						testObject = positiveTestData[i]
						testCreditModel = &Credit{testObject.id, testObject.person,
							testObject.amount, testObject.agreementAt,
							testObject.currency, testObject.duration,
							testObject.percent, testObject.rounding,
							0, testObject.regularPayment,
							[]valueObjects.PaymentInterface(nil)}
					})
					g.Describe("#Credit  model creation", func() {
						g.It("Create credit", func() {
							res, _ := GenCredit(testObject.person, testObject.amount, testObject.agreementAt,
								testObject.currency, testObject.duration, testObject.percent)
							g.Assert(res).Equal(testCreditModel)
						})

					})
					g.Describe("#Credit  model getters", func() {
						g.It("get currency", func() {
							res := testCreditModel.GetCurrency()
							g.Assert(res).Equal(testObject.currency)
						})
						g.It("get percent", func() {
							res := testCreditModel.GetPercent()
							g.Assert(res).Equal(testObject.percent)
						})
						g.It("get person", func() {
							res := testCreditModel.GetPerson()
							g.Assert(res).Equal(testObject.person)
						})
						g.It("get agreement at", func() {
							res := testCreditModel.GetAgreementAt()
							g.Assert(res).Equal(testObject.agreementAt)
						})
						g.It("get amount", func() {
							res := testCreditModel.GetAmount()
							g.Assert(res).Equal(testObject.amount)
						})
						g.It("get id", func() {
							res := testCreditModel.GetId()
							g.Assert(res).Equal(testObject.id)
						})
						g.It("get duration", func() {
							res := testCreditModel.GetDuration()
							g.Assert(res).Equal(testObject.duration)
						})
						g.It("get annuity payment", func() {
							res := testCreditModel.getAnnuityPayment()
							g.Assert(res).Equal(testObject.annuityPayment)
						})
						g.It("get rounding", func() {
							res := testCreditModel.GetRounding()
							g.Assert(res).Equal(testObject.rounding)
						})
						g.It("get regular payment", func() {
							res := testCreditModel.getRegularPayment(testObject.annuityPayment)
							g.Assert(res).Equal(testObject.regularPayment)
						})
					})
					g.Describe("#Credit  model setters", func() {
						g.It("Credit model set id", func() {
							testCreditModel.SetId(i*2 + 1)
							g.Assert(testCreditModel.GetId()).Equal(i*2 + 1)
						})

					})
				})
			}
		})
		g.It("Negative tests", func() {
			for i := 0; i < numOfTestsNeg; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						testObject = negativeTestData[i]
					})
					g.BeforeEach(func() {
						testCreditModel = &Credit{testObject.id, testObject.person,
							testObject.amount, testObject.agreementAt,
							testObject.currency, testObject.duration,
							testObject.percent, testObject.rounding,
							0, testObject.regularPayment,
							[]valueObjects.PaymentInterface(nil)}
					})
					g.Describe("Credit model tests for errors", func() {
						g.It("create model", func() {
							res, _ := GenCredit(testObject.person, testObject.amount, testObject.agreementAt,
								testObject.currency, testObject.duration, testObject.percent)
							g.Assert(res).Equal(nil)
						})
						g.It("get rounding", func() {
							annuityPayment := testCreditModel.getAnnuityPayment()
							g.Assert(annuityPayment).Equal(testObject.annuityPayment)
							res, _ := testCreditModel.getRounding(annuityPayment)
							g.Assert(res).Equal(int32(-1))
						})

					})
				})
			}
		})
	})
}
