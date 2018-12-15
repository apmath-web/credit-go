package repositories

import (
	"github.com/apmath-web/credit-go/data"
	"github.com/apmath-web/credit-go/models"
	"github.com/apmath-web/credit-go/valueObjects"
	"github.com/franela/goblin"
	"strconv"
	"testing"
	"time"
)

func TestCreditRepository(t *testing.T) {
	g := goblin.Goblin(t)
	type TestData struct {
		valueObject models.CreditInterface
	}
	var valueObjectTest models.CreditInterface
	var testRepository CreditsRepositoryInterface
	var testData []TestData
	var numOfTests int
	//var numberOfCredits int32
	g.Describe("CreditRepository tests", func() {
		g.Before(func() {

			testData = []TestData{
				{nil},
				{nil},
				{nil},
			}
			testData[0].valueObject, _ = models.GenCredit(valueObjects.GenPerson("FName", "dfs"),
				2000, data.Str2Date("2017-12-10"), "RUR", 12, 10)
			testData[1].valueObject, _ = models.GenCredit(valueObjects.GenPerson("FName", "dfs"),
				600, data.Str2Date(data.Date(time.Now()).Date2Str()), "EUR", 6, 1)
			testData[2].valueObject, _ = models.GenCredit(valueObjects.GenPerson("FName", "dfs"),
				3000000000000000, data.Str2Date(data.Date(time.Now()).Date2Str()), "USD", 1200, 300)
			numOfTests = len(testData)
		})
		g.It("all tests", func() {
			g.Describe("Create empty repository", func() {
				g.Before(func() {
					testRepository = &CreditRepository{make(map[int]models.CreditInterface), 0}
				})
				g.It("Gen repository", func() {
					res := GenRepository()
					g.Assert(res).Equal(testRepository)
				})
			})
			for i := 0; i < numOfTests; i++ {
				g.Describe("Test #"+strconv.Itoa(i+1), func() {
					g.Before(func() {
						valueObjectTest = testData[i].valueObject

					})
					g.It("Add credit", func() {
						testRepository.Store(valueObjectTest)
						g.Assert(valueObjectTest.GetId()).Equal(i + 1)
					})
					g.It("Get credit", func() {
						res := testRepository.Get(i + 1)
						g.Assert(res).Equal(valueObjectTest)
					})
					g.It("Del credit", func() {
						testRepository.Remove(valueObjectTest)
						res := testRepository.Get(i + 1)
						g.Assert(res).Equal(nil)
					})
					g.It("Get credit not found", func() {
						res := testRepository.Get(i * 100)
						g.Assert(res).Equal(nil)
					})
				})
			}
		})
	})
}
