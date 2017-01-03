package testcase

import (
	"testing"
	"database/sql"
	"../dbmanager"
	"os"
)

func TestDbInit(t *testing.T) {
	os.Remove("./t.db")
	sou := &sql.DB{}
	err := dbmanager.Init(sou, "./t.db")
	if err != nil {
		t.Error("dbmanager init test fail")
	}

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
