package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/apmath-web/credit-go/valueObjects"
	"github.com/franela/goblin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func GenRequestFromUrlAndJson(url string, Json string) *http.Request {
	req, _ := http.NewRequest("GET", url, ioutil.NopCloser(bytes.NewBufferString(Json)))
	return req
}

func GenRequestFromUrl(url string) *http.Request {
	return GenRequestFromUrlAndJson(url, "")
}

func TestActionsUtils(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("get id test", func() {
		g.It("/credit/<id>", func() {
			for i := 0; i < 15; i++ {
				id, _ := getId("/credit/" + strconv.Itoa(i))
				g.Assert(id).Equal(i)
			}
		})
		g.It("/credit/<id>/", func() {
			for i := 0; i < 15; i++ {
				id, _ := getId("/credit/" + strconv.Itoa(i) + "/")
				g.Assert(id).Equal(i)
			}
		})
		g.It("/credit/<id>/payments", func() {
			for i := 0; i < 15; i++ {
				id, _ := getId("/credit/" + strconv.Itoa(i) + "/payments")
				g.Assert(id).Equal(i)
			}
		})
	})
	g.Describe("validate params test", func() {
		g.Describe("positive tests", func() {
			g.It("empty", func() {
				res := validateParam("", []string{"next"}, true)
				g.Assert(res).Equal(nil)
			})
			g.It("next in params", func() {
				res := validateParam("next", []string{"next", "early", "regular"}, true)
				g.Assert(res).Equal(nil)
			})
			g.It("early in params", func() {
				res := validateParam("early", []string{"next", "early", "regular"}, true)
				g.Assert(res).Equal(nil)
			})
			g.It("regular in params", func() {
				res := validateParam("regular", []string{"next", "early", "regular"}, true)
				g.Assert(res).Equal(nil)
			})
		})
		g.Describe("negative tests", func() {
			g.It("empty", func() {
				res := validateParam("", []string{"next"}, false)
				g.Assert(res != nil).IsTrue()
			})
			g.It("not in params", func() {
				res := validateParam("None", []string{"next", "early", "regular"}, true)
				g.Assert(res != nil).IsTrue()
			})
		})
	})
	g.Describe("error message test", func() {
		g.It("empty message", func() {
			var respTest, respExt httptest.ResponseRecorder
			respExt.WriteHeader(200)
			fmt.Fprintf(&respExt, "{\"message\":\"%s\"}", "")
			errorMessage("", 200, &respTest)
			g.Assert(respTest).Equal(respExt)
		})
		g.It("not empty message", func() {
			var respTest, respExt httptest.ResponseRecorder
			respExt.WriteHeader(200)
			fmt.Fprintf(&respExt, "{\"message\":\"%s\"}", "OK")
			errorMessage("OK", 200, &respTest)
			g.Assert(respTest).Equal(respExt)
		})
		g.It("http status 400", func() {
			var respTest, respExt httptest.ResponseRecorder
			respExt.WriteHeader(400)
			fmt.Fprintf(&respExt, "{\"message\":\"%s\"}", "OK")
			errorMessage("OK", 400, &respTest)
			g.Assert(respTest).Equal(respExt)
		})
		g.It("http status 404", func() {
			var respTest, respExt httptest.ResponseRecorder
			respExt.WriteHeader(404)
			fmt.Fprintf(&respExt, "{\"message\":\"%s\"}", "OK")
			errorMessage("OK", 404, &respTest)
			g.Assert(respTest).Equal(respExt)
		})
	})
	g.Describe("message to json test", func() {
		g.It("empty description", func() {
			res := ptrMessagesToJsonErrMessage("message", nil)
			g.Assert(res).Equal("{\"message\":\"message\"}")
		})
		g.It("non empty description", func() {
			validator := new(valueObjects.Validation)
			validator.AddMessage(valueObjects.GenMessage("type", "test"))
			resStr := ptrMessagesToJsonErrMessage("message", validator.GetMessages())
			exp := map[string]interface{}{"description": map[string]interface{}{"type": "test"}, "message": "message"}
			var res interface{}
			json.Unmarshal([]byte(resStr), &res)
			g.Assert(res).Equal(exp)
		})
	})
	g.Describe("get param test", func() {
		g.It("get exist param", func() {
			req := GenRequestFromUrl("http://aaaaa.aa/?type=next")
			res := getParam(req, "type")
			g.Assert(res).Equal("next")
		})
		g.It("get wrong param", func() {
			req := GenRequestFromUrl("http://aaaaa.aa/?type=next")
			res := getParam(req, "state")
			g.Assert(res).Equal("")
		})
	})
	g.Describe("to json test", func() {
		g.It("positive test", func() {
			resp := httptest.ResponseRecorder{}
			req := GenRequestFromUrlAndJson("", "{\"message\":\"message\"}")
			res := toJson(&resp, req)
			g.Assert(res).Equal(map[string]interface{}{"message": "message"})
		})
		g.It("negative test", func() {
			resp := httptest.ResponseRecorder{}
			req := GenRequestFromUrlAndJson("", "{\"message\":message\"}")
			res := toJson(&resp, req)
			g.Assert(res).Equal(map[string]interface{}(nil))
		})
	})
}
