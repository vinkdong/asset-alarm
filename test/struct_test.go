package testcase

import (
	"testing"
	"../server"
	"../dbmanager"
	"database/sql"
	"os"
	"github.com/bitly/go-simplejson"
)

func TestCreditSave(t *testing.T) {
	os.Remove("./t.db")
	a := server.Credit{Name: "招商银行"}
	db := &sql.DB{}
	dbmanager.Init(db, "./t.db")
	server.Context.Db = db
	dbmanager.InitTables(db)
	a.Save()
	if a.Id != 1{
		t.Error("insert one record id should be 1")
	}
}

func TestCreditObj2JsonString(t *testing.T)  {
	a := server.Credit{Name:"Vink Bank"}
	str := a.ToJsonString()
	js,err := simplejson.NewJson([]byte(str))
	if err != nil {
		t.Error("convert json error")
		return
	}
	name := js.Get("name").MustString()
	expect := "Vink Bank"
	if name != expect {
		t.Error("convert json value is not like expect")
	}
}

func TestCreditOjb2Json(t *testing.T){
	a := server.Credit{Name:"Vink Bank"}
	js := a.ToJson()
	expect := "Vink Bank"
	if js.Get("name").MustString() != expect {
		t.Error("get convert json value is not like expect")
	}
}

func TestConvertFromJson(t *testing.T) {
	js, err := simplejson.NewJson([]byte(`
{
	"name":"Vink Bank",
	"icon":"../icon/vink.logo",
	"credit":10.000000,
	"debit":50.000000,
	"balance":10.000000,
	"account_date":8,
	"repayment_date":0,
	"id":9
}
	`))
	if err != nil{
		t.Error("convert credit from json error")
	}
	expect := "Vink Bank"
	result := js.Get("name").MustString()
	if expect != result {
		t.Errorf("expect json->name is %s but got %s", expect, result)
	}
}

func TestRecordSave(t *testing.T) {
	os.Remove("./t.db")
	TestDbInit(t)
	TestExits(t)
	server.Context.Db = sou
	a := server.Record{CreditId: 1, Credit: 9, Amount: 8, Time: "2017-01-20 20:22:01"}
	a.Save()
	if a.Id != 1 {
		t.Error("insert one record id should be 1")
	}
}