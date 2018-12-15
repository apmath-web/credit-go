package data

import (
	"github.com/franela/goblin"
	"strconv"
	"testing"
)

func TestStateData(t *testing.T) {
	g := goblin.Goblin(t)
	type TestData struct {
		state string
	}
	var testType State
	var testDataObject string
	var positiveTestData, negativeTestData []TestData
	var numberOfTestsPos, numberOfTestsNeg int
	g.Describe("State data tests", func() {
		g.Before(func() {
			positiveTestData = []TestData{
				{"upcoming"},
				{"paid"},
			}
			negativeTestData = []TestData{
				{""},
				{"Upcoming"},
				{"Paid"},
				{"UPCOMING"},
				{"PAID"},
				{"u p c o m i n g"},
				{"REGULAR"},
				{"p a i d"},
				{"\n"},
				{" "},
			}
			numberOfTestsNeg = len(negativeTestData)
			numberOfTestsPos = len(positiveTestData)
		})
		g.It("Positive tests for state", func() {
			for i := 0; i < numberOfTestsPos; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						testType = State(positiveTestData[i].state)
						testDataObject = positiveTestData[i].state
					})
					g.It("state2str state", func() {
						g.Assert(testType.State2Str()).Equal(testDataObject)

					})
				})
			}
		})
		g.It("Negative tests for state", func() {
			for i := 0; i < numberOfTestsNeg; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						testType = State(negativeTestData[i].state)
						testDataObject = negativeTestData[i].state
					})
					g.It("state2str test", func() {
						g.Assert(testType.State2Str()).Equal("")

					})
				})
			}
		})
	})
}
