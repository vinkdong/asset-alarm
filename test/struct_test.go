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