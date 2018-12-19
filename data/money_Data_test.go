package data

import (
	"github.com/franela/goblin"
	"strconv"
	"testing"
)

func TestMoneyData(t *testing.T) {
	g := goblin.Goblin(t)
	type TestData struct {
		money      string
		moneyInt   int64
		moneyFloat float64
	}
	var testMoney Money
	var testDataObject TestData
	var positiveTestData, negativeTestData []TestData
	var numberOfTestsPos, numberOfTestsNeg int
	g.Describe("Money data tests", func() {
		g.Before(func() {
			positiveTestData = []TestData{
				{"0", 0, 0.0},
				{"100", 100, 100.0},
				{"10000", 10000, 10000.0},
				{"1000000", 1000000, 1000000.0},
				{"100000000", 100000000, 100000000.0},
			}
			negativeTestData = []TestData{
				{"10000000000000000000000000000000", -1, -1.0},
			}
			numberOfTestsNeg = len(negativeTestData)
			numberOfTestsPos = len(positiveTestData)
		})
		g.It("Positive tests for money", func() {
			for i := 0; i < numberOfTestsPos; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						testMoney = Money(positiveTestData[i].moneyInt)
						testDataObject = positiveTestData[i]
					})
					g.It("str2mon test", func() {
						g.Assert(Str2Mon(testDataObject.money)).Equal(testMoney)
					})
					g.It("mon2str test", func() {
						g.Assert(testMoney.Mon2Str()).Equal(testDataObject.money)
					})
					g.It("mon2int test", func() {
						g.Assert(testMoney.Mon2Int64()).Equal(testDataObject.moneyInt)
					})
					g.It("mon2float test", func() {
						g.Assert(testMoney.Mon2Float64()).Equal(testDataObject.moneyFloat)
					})
				})
			}
		})
		g.It("Negative tests for money", func() {
			for i := 0; i < numberOfTestsNeg; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						testMoney = Money(negativeTestData[i].moneyInt)
						testDataObject = negativeTestData[i]
					})
					g.It("str2mon test", func() {
						g.Assert(Str2Mon(testDataObject.money)).Equal(testMoney)
					})
				})
			}
		})
	})
}
