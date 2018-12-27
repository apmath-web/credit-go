package data

import (
	"github.com/franela/goblin"
	"strconv"
	"testing"
)

func TestCurrencyData(t *testing.T) {
	g := goblin.Goblin(t)
	type TestData struct {
		currency string
	}
	var testCurrency Currency
	var testDataObject string
	var positiveTestData, negativeTestData []TestData
	var numberOfTestsPos, numberOfTestsNeg int
	g.Describe("Currency data tests", func() {
		g.Before(func() {
			positiveTestData = []TestData{
				{"USD"},
				{"EUR"},
				{"RUR"},
			}
			negativeTestData = []TestData{
				{""},
				{"RUB"},
				{"rur"},
				{"usd"},
				{"eur"},
				{" "},
				{"something"},
				{"E U R"},
				{"TeSt"},
			}
			numberOfTestsNeg = len(negativeTestData)
			numberOfTestsPos = len(positiveTestData)
		})
		g.It("Positive tests for type_", func() {
			for i := 0; i < numberOfTestsPos; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						testCurrency = Currency(positiveTestData[i].currency)
						testDataObject = positiveTestData[i].currency
					})
					g.It("str2cur test", func() {
						g.Assert(Str2Cur(testDataObject)).Equal(testCurrency)
					})
					g.It("cur2str test", func() {
						g.Assert(testCurrency.Cur2Str()).Equal(testDataObject)

					})
				})
			}
		})
		g.It("Negative tests for type_", func() {
			for i := 0; i < numberOfTestsNeg; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						testCurrency = Currency(negativeTestData[i].currency)
						testDataObject = negativeTestData[i].currency
					})
					g.It("str2cur test", func() {
						g.Assert(Str2Cur(testDataObject)).Equal(Currency(""))
					})
					g.It("cur2str test", func() {
						g.Assert(testCurrency.Cur2Str()).Equal("")

					})
				})
			}
		})
	})
}
