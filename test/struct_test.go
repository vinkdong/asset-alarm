package testcase

import (
	"testing"
	"../server"
	"../dbmanager"
	"database/sql"
	"os"
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
