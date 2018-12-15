package data

import (
	"github.com/franela/goblin"
	"strconv"
	"testing"
)

func TestTypeData(t *testing.T) {
	g := goblin.Goblin(t)
	type TestData struct {
		type_ string
	}
	var testType Type
	var testDataObject string
	var positiveTestData, negativeTestData []TestData
	var numberOfTestsPos, numberOfTestsNeg int
	g.Describe("Type data tests", func() {
		g.Before(func() {
			positiveTestData = []TestData{
				{"regular"},
				{"next"},
				{"early"},
				{""},
			}
			negativeTestData = []TestData{
				{" "},
				{"Next"},
				{"Regular"},
				{"Early"},
				{"eur"},
				{"NEXT"},
				{"REGULAR"},
				{"EARLY"},
				{"\n"},
			}
			numberOfTestsNeg = len(negativeTestData)
			numberOfTestsPos = len(positiveTestData)
		})
		g.It("Positive tests for type", func() {
			for i := 0; i < numberOfTestsPos; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						testType = Type(positiveTestData[i].type_)
						testDataObject = positiveTestData[i].type_
					})
					g.It("str2type test", func() {
						g.Assert(Str2Type(testDataObject)).Equal(testType)
					})
					g.It("type2str test", func() {
						g.Assert(testType.Type2Str()).Equal(testDataObject)

					})
				})
			}
		})
		g.It("Negative tests for type", func() {
			for i := 0; i < numberOfTestsNeg; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						testType = Type(negativeTestData[i].type_)
						testDataObject = negativeTestData[i].type_
					})
					g.It("str2type test", func() {
						g.Assert(Str2Type(testDataObject)).Equal(None)
					})
					g.It("type2str test", func() {
						g.Assert(testType.Type2Str()).Equal("")

					})
				})
			}
		})
	})
}
