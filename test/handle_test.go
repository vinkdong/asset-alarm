package testcase

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"../server"
	"github.com/bitly/go-simplejson"
)

func TestHandlerList(t *testing.T) {
	TestExits(t)
	server.Context.Db = sou
	A := server.Credit{Name: "TH bank", Credit: 100000}
	B := server.Credit{Name: "B bank", Credit: 200000}
	A.Save()
	B.Save()

	req, err := http.NewRequest("GET", "/api/list", nil)
	if err != nil {
		t.Error("http request init fail")
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.HandlerList)

	handler.ServeHTTP(rr, req)
	js, err := simplejson.NewFromReader(rr.Body)
	if err != nil{
		t.Error("response data is not json")
	}
	name0 := js.Get("credits").GetIndex(0).Get("name").MustString()
	expect := "TH bank"
	if name0 != expect {
		t.Errorf("expect bank0's name is %s but got %s", expect, name0)
	}
}
