package viewModelTests

import (
	"github.com/apmath-web/credit-go/viewModels"
	"github.com/franela/goblin"
	"testing"
)

func TestPersonViewModelPositive(t *testing.T) {
	g := goblin.Goblin(t)

	testData := map[string]interface{}{
		"firstName": "Fname",
		"lastName":  "Lname",
	}
	testPersonViewModel := new(viewModels.Person)

	g.Describe("#PersonViewModelPositive", func() {
		g.It("validate correct json package", func() {
			testPersonViewModel.Fill(testData)
			res := testPersonViewModel.Validate()
			g.Assert(res).IsTrue()
		})
		g.It("person view model save data", func() {
			savedData := testPersonViewModel.Fetch()
			g.Assert(savedData).Equal(savedData)
		})
	})
}
