package valueObjects

import (
	"github.com/franela/goblin"
	"strconv"
	"testing"
)

func TestPersonValueObject(t *testing.T) {
	g := goblin.Goblin(t)
	type TestData struct {
		firstName string
		lastName  string
	}
	var testPerson *Person
	var lastName, firstName string
	var testData []TestData
	var testDataObject TestData
	var numberOfTests int
	g.Describe("Person value object tests", func() {
		g.Before(func() {
			testData = []TestData{
				{"firstName", "lastName"},
				{"firstName", "lastName"},
				{"f i r s t n a m e", "l  a s t n a m e"},
				{"one first name", "one last name"},
				{"f", "l"},
			}
			numberOfTests = len(testData)
		})
		g.It("all tests for person", func() {
			for i := 0; i < numberOfTests; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						testDataObject = testData[i]
						lastName = testDataObject.lastName
						firstName = testDataObject.firstName
						testPerson = new(Person)
					})
					g.It("GenPerson test", func() {
						testPerson.FirstName = firstName
						testPerson.LastName = lastName
						g.Assert(GenPerson(firstName, lastName)).Equal(testPerson)
					})
					g.It("Person getters", func() {
						g.Assert(testPerson.GetFirstName()).Equal(firstName)
						g.Assert(testPerson.GetLastName()).Equal(lastName)
					})
				})
			}
		})
	})
}
