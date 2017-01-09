package testcase

import "testing"
import (
	"../server"
	"../dbmanager"
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