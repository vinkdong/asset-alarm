package testcase

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/VinkDong/asset-alarm/server"
	"github.com/bitly/go-simplejson"
	"os"
	"strings"
	"github.com/VinkDong/asset-alarm/dbmanager"
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

func TestHandlerItemAdd(t *testing.T) {
	os.Remove("./t.db")
	TestExits(t)
	server.Context.Db = sou
	A := server.Credit{Name: "TH bank", Credit: 100000}
	A.Save()

	data := strings.NewReader(`
	{
	"version":"v0.1",
	"credit" : {
		"name":"Vink Bank",
		"icon":"../icon/vink.logo",
		"credit":10.000000,
		"debit":50.000000,
		"balance":10.000000,
		"account_date":8,
		"repayment_date":0
	}
}
	`)

	req, err := http.NewRequest("GET", "/api/item/add", data)
	if err != nil {
		t.Error("http request init fail")
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.HandLerAddItem)

	handler.ServeHTTP(rr, req)
	js, err := simplejson.NewFromReader(rr.Body)
	if err != nil{
		t.Error("response data is not json")
	}
	success := js.Get("success").MustBool()
	expect := true
	if success != expect {
		t.Error("add item should be true but go false")
	}

	r, err := dbmanager.PatchData(sou, "credit")
	if err != nil{
		t.Error("patch database error")
		return
	}
	defer r.Close()

	c := &server.Credit{}
	r.Next()
	for r.Next() {
		if err := c.ConvertFormRow(r); err != nil {
			t.Error("test patch data Scan error\n")
			t.Error(err)
		}
		if c.Name != "Vink Bank" {
			t.Errorf("test patch data content expect 'Vink Bank' but get %s", c.Name)
		}
	}
}

func TestHandlerRecordAdd(t *testing.T) {
	os.Remove("./t.db")
	TestExits(t)
	server.Context.Db = sou
	A := server.Record{CreditId: 1, Credit: 100000}
	A.Save()

	data := strings.NewReader(`
{
	"version":"v0.1",
	"record" : {
		"cid":1,
		"type":"out",
		"credit":10.000000,
		"debit":50.000000,
		"amount":10.000000,
		"time":"2017-01-21 20:08:09"
	}
}
	`)

	req, err := http.NewRequest("GET", "/api/record/add", data)
	if err != nil {
		t.Error("http request init fail")
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.HandLerAddRecord)

	handler.ServeHTTP(rr, req)
	js, err := simplejson.NewFromReader(rr.Body)
	if err != nil{
		t.Error("response data is not json")
	}
	success := js.Get("success").MustBool()
	expect := true
	if success != expect {
		t.Error("add item should be true but go false")
	}

	r, err := dbmanager.PatchData(sou, "record")
	if err != nil{
		t.Error("patch database error")
		return
	}
	defer r.Close()

	c := &server.Record{}
	r.Next()
	for r.Next() {
		if err := c.ConvertFormRow(r); err != nil {
			t.Error("test patch data Scan error\n")
			t.Error(err)
		}
		if c.Debit != 50.000 {
			t.Errorf("test patch data content expect 'Vink Bank' but get %s", c.Debit)
		}
	}
}

func TestCheckStaticResources(t *testing.T) {
	path := "a.js"
	expect := server.CheckStaticResources(path,"js","png")
	if expect != true {
		t.Error(".js should be see as static resources")
	}

	path = "b.png"
	expect = server.CheckStaticResources(path,"js","png")
	if expect != true {
		t.Error(".png should be see as static resources")
	}

	path = "c.do"
	expect = server.CheckStaticResources(path,"js","png")
	if expect == true {
		t.Error(".do shouldn't be see as static resources")
	}

}