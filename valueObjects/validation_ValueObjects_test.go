package valueObjects

import (
	"github.com/franela/goblin"
	"strconv"
	"testing"
)

func TestValidationValueObject(t *testing.T) {
	g := goblin.Goblin(t)
	type TestData struct {
		messages []MessageInterface
	}
	var testValidation *Validation
	var testData []TestData
	var testDataObject []MessageInterface
	var numberOfTests int
	g.Describe("Message value object tests", func() {
		g.Before(func() {
			testData = []TestData{
				{[]MessageInterface{GenMessage("field1", "text1")}},
				{[]MessageInterface{GenMessage("field1", "text1"),
					GenMessage("field2", "text2")}},
				{[]MessageInterface{GenMessage("field1", "text1"),
					GenMessage("field2", "text2"),
					GenMessage("field3", "text3")}},
				{[]MessageInterface{GenMessage("field1", "text1"),
					GenMessage("field2", "text2"),
					GenMessage("field3", "text3"),
					GenMessage("field4", "text4")}},
				{[]MessageInterface{GenMessage("field1", "text1"),
					GenMessage("field2", "text2"),
					GenMessage("field3", "text3"),
					GenMessage("field4", "text4"),
					GenMessage("field5", "text5")}},
			}
			numberOfTests = len(testData)
		})
		g.It("all tests for validation", func() {
			for i := 0; i < numberOfTests; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						testDataObject = testData[i].messages
						testValidation = new(Validation)
					})
					g.It("add messages test", func() {
						for _, el := range testDataObject {
							testValidation.AddMessage(el)
						}
						g.Assert(testValidation.messages).Equal(testDataObject)
					})
					g.It("get messages", func() {
						g.Assert(testValidation.GetMessages()).Equal(testDataObject)
					})
				})
			}
		})
	})
}
