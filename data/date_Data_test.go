package data

import (
	"github.com/franela/goblin"
	"strconv"
	"testing"
	"time"
)

func TestDateData(t *testing.T) {
	g := goblin.Goblin(t)
	type TestData struct {
		date string
	}
	var testDate Date
	var testDataObject string
	var positiveTestData, negativeTestData []TestData
	var numberOfTestsPos, numberOfTestsNeg int
	g.Describe("Currency data tests", func() {
		g.Before(func() {
			positiveTestData = []TestData{
				{"2018-01-01"},
				{"2006-01-02"},
			}
			negativeTestData = []TestData{
				{"2018/01/02"},
				{"2018.01.02"},
				{"01.02.2018"},
				{"01/02.2018"},
				{"01-02-2018"},
				{"2018 01 02"},
				{"01 02 2018"},
				{"2019-13-20"},
				{"2018-01-35"},
				{"2019-3-20"},
				{"2018-01-5"},
			}
			numberOfTestsNeg = len(negativeTestData)
			numberOfTestsPos = len(positiveTestData)
		})
		g.It("Positive tests for currency", func() {
			for i := 0; i < numberOfTestsPos; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						date, _ := time.Parse("2006-01-02", positiveTestData[i].date)
						testDate = Date(date)
						testDataObject = positiveTestData[i].date
					})
					g.It("str2cur test", func() {
						g.Assert(Str2Date(testDataObject)).Equal(testDate)
					})
					g.It("cur2str test", func() {
						g.Assert(testDate.Date2Str()).Equal(testDataObject)

					})
				})
			}
		})
		g.It("Negative tests for currency", func() {
			for i := 0; i < numberOfTestsNeg; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						date, err := time.Parse("2006-01-02", negativeTestData[i].date)
						if err != nil {
							date = time.Now()
						}
						testDate = Date(date)
						testDataObject = negativeTestData[i].date
					})

					g.It("str2date test", func() {
						g.Assert(Str2Date(testDataObject).Date2Str()).Equal(testDate.Date2Str())
					})
					g.It("date2str test", func() {
						g.Assert(testDate.Date2Str()).Equal(Date(time.Now()).Date2Str())
					})
				})
			}
		})
	})
}
