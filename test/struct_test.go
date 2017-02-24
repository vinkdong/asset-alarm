package testcase

import (
	"testing"
	"github.com/VinkDong/asset-alarm/server"
	"github.com/VinkDong/asset-alarm/dbmanager"
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

func TestModifyCreditWhenNewRecordAdded(t *testing.T)  {
	os.Remove("./t.db")
	TestDbInit(t)
	TestExits(t)
	server.Context.Db = sou
	c := server.Credit{Debit: 100.0,Credit:100000}
	c.Save()
	a := server.Record{CreditId: 1, Credit: 9, Amount: 8, Time: "2017-01-20 20:22:01"}
	a.Save()
	c.Browse(1)
	if c.Debit != 108.00 {
		t.Errorf("modify credit want get 108.00 but got value %f", c.Debit)
	}
}

func TestRecordFromJson(t *testing.T) {
	js, _ := simplejson.NewJson([]byte(`
{
		"cid":1,
		"type":"out",
		"credit":10.000000,
		"debit":50.000000,
		"amount":10.000000,
		"time":"2017-01-21 20:08:09"
}
	`))
	r := server.Record{}
	r.ConvertFromJson(js)
	expect := "out"
	if r.Type != expect {
		t.Errorf("expect record type is %s but got %s", expect, r.Type)
	}
}

func TestCreditBrowse(t *testing.T) {
	TestCreditSave(t)
	c := server.Credit{}
	c.Browse(1)
	expect := "招商银行"
	if c.Name != expect {
		t.Errorf("expect browse 1 of credit name is %s but got %s", expect, c.Name)
	}
}

func TestInterface2map(t *testing.T) {
	a := server.Bill{Id:5,Balance:10}
	a_v := server.Interface2map(a)

	expect_a_id := "5"
	if a_v["id"] != expect_a_id{
		t.Errorf("expect id of a is %s but got %s",expect_a_id,a_v["id"])
	}
	b := server.Credit{Name:"CMB Bank"}
	b_v := server.Interface2map(b)
	expect_b_name := "\"CMB Bank\""
	if b_v["name"] != expect_b_name{
		t.Errorf("expect name of b is %s but got %s",expect_b_name,b_v["name"])
	}
}

func TestKeyToColumn(t *testing.T) {
	a := "apple"
	b := "deskNote"
	if ea := server.PackToCol(a); ea != "apple" {
		t.Errorf("expect apple to apple but got %s", ea)
	}
	if eb := server.PackToCol(b); eb != "desk_note" {
		t.Errorf("expect deskNote to desk_note but got %s", eb)
	}
}
