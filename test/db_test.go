package testcase

import (
	"testing"
	"database/sql"
	"../dbmanager"
	"../server"
	"os"
	"fmt"
)

var sou = &sql.DB{}

func TestDbInit(t *testing.T) {
	sou = &sql.DB{}
	err := dbmanager.Init(sou, "./t.db")
	if err != nil {
		t.Error("dbmanager init test fail")
	}
}

func TestExits(t *testing.T) {
	os.Remove("./t.db")
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
	defer r.Close()
	for r.Next() {
		var accout_data, repayment_date int8
		var id int
		var amount, debit, balance float64
		var name, icon string
		if err := r.Scan(&id, &name, &icon, &amount, &debit, &balance, &accout_data, &repayment_date); err != nil {
			t.Error("test patch data Scan error\n")
			t.Error(err)
		}
		if name != "招商银行" {
			t.Errorf("test patch data content expect '招商银行' but get %s", name)
		}
		fmt.Println(name)
	}
}