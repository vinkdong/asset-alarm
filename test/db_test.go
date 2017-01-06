package testcase

import (
	"testing"
	"database/sql"
	"../dbmanager"
	"../server"
	"os"
)

var sou = &sql.DB{}

func TestDbInit(t *testing.T) {
	os.Remove("./t.db")
	sou = &sql.DB{}
	err := dbmanager.Init(sou, "./t.db")
	if err != nil {
		t.Error("dbmanager init test fail")
	}
}

func TestExits(t *testing.T) {
	TestDbInit(t)
	exits := dbmanager.Exists(sou, "credit")
	if exits == true {
		t.Error("dbmanager init test fail exit table asset should be false")
	}

	dbmanager.InitTables(sou)
	exits = dbmanager.Exists(sou, "credit")
	if exits == false {
		t.Error("dbmanager init test fail exit table asset should be true")
	}
}

func TestPatchData(t *testing.T) {
	TestExits(t)
	server.Context.Db = sou
	a := server.Credit{Name: "招商银行"}
	a.Save()
	r, err := dbmanager.PatchData(sou, "credit")
	if err != nil{
		t.Error("patch database error")
		return
	}
	_, err = r.Columns()
	if err != nil {
		t.Error("convert rows to string[] error")
	}
}