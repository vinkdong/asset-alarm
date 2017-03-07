package testcase

import "testing"
import (
	"github.com/VinkDong/asset-alarm/server"
	"github.com/VinkDong/asset-alarm/dbmanager"
	"github.com/bitly/go-simplejson"
)

func TestParseRowsToCreditList(t *testing.T) {
	TestExits(t)
	server.Context.Db = sou
	a := server.Credit{Name: "招商银行"}
	a.Save()
	b := server.Credit{Name: "Vink Bank"}
	b.Save()
	r, err := dbmanager.PatchData(sou, "credit")
	if err != nil{
		t.Error("patch database error")
		return
	}
	var c_list= &[]server.Credit{}
	server.ParseRowsToCreditList(r, c_list)
	c := * c_list
	if c[1].Name != "Vink Bank" {
		t.Errorf("test TestParseRowsToCreditList expect get Vink Bank but get %s", c[1].Name)
	}
}

func TestParserCreditsToJson(t *testing.T) {
	a := server.Credit{Name: "a bank", Credit: 1000.0}
	b := server.Credit{Name: "b bank", Debit: 2}
	l := make([]server.Credit, 0)
	l = append(l, a)
	l = append(l, b)
	js := server.ParserCreditsToJson(&l)
	credits := js.Get("credits")
	aAmount := credits.GetIndex(0).Get("credit").MustFloat64()
	if aAmount != 1000.0 {
		t.Errorf("test get credit:0's credit not correct get %f", aAmount)
	}
}

func TestCommonToJsonStr(t *testing.T)  {
	b := server.Bill{CreditId:1}
	jsonStr := server.CommonToJsonStr(&b)
    js,_ := simplejson.NewJson([]byte(jsonStr))
	if js.Get("credit_id").MustInt64()  != 1 {
		t.Errorf("expect credit_id of js object is 1 but got %d",js.Get("credit_id").MustInt64())
	}
}