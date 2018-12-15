package valueObjects

import (
	"github.com/franela/goblin"
	"strconv"
	"testing"
)

func TestMessageValueObject(t *testing.T) {
	g := goblin.Goblin(t)
	type TestData struct {
		field string
		text  string
	}
	var testMessage *Message
	var field, text string
	var testData []TestData
	var testDataObject TestData
	var numberOfTests int
	g.Describe("Message value object tests", func() {
		g.Before(func() {
			testData = []TestData{
				{"field", "text"},
				{"FIELD", "TEXT"},
				{"f i e l d", "t e x t"},
				{"one field", "one text"},
				{"f", "t"},
			}
			numberOfTests = len(testData)
		})
		g.It("all tests for message", func() {
			for i := 0; i < numberOfTests; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						testDataObject = testData[i]
						field = testDataObject.field
						text = testDataObject.text
						testMessage = new(Message)
					})
					g.It("GenMessage test", func() {
						testMessage.text = text
						testMessage.field = field
						g.Assert(GenMessage(field, text)).Equal(testMessage)
					})
					g.It("Message getters", func() {
						g.Assert(testMessage.GetText()).Equal(text)
						g.Assert(testMessage.GetField()).Equal(field)
					})
				})
			}
		})
	})
}
