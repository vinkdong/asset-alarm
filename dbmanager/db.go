package dbmanager

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/VinkDong/asset-alarm/log"
	"fmt"
)

const (
	SQL_EXIT = "SELECT name FROM sqlite_master WHERE type='table' AND name='%s'"
)

func Init(db *sql.DB, db_source string) error {
	dbInit, err := sql.Open("sqlite3", db_source)
	*db = *dbInit
	return err
}

func InitTables(db *sql.DB) error {
	sqlStmt := `create table credit(
id integer not null primary key,
name text,
icon text,
credit float,
debit float,
balance float,
account_date int,
repayment_date int
)`
	_, err := db.Exec(sqlStmt)
	return err
}

func InitRecordTable(db *sql.DB) error {
	sqlStmt := `
CREATE TABLE record (
  id        INTEGER NOT NULL PRIMARY KEY,
  credit_id INT,
  type      TEXT,
  amount    FLOAT,
  credit    FLOAT,
  debit     FLOAT,
  time      DATETIME
)
`
	_, err := db.Exec(sqlStmt)
	return err
}

func InitBillTable(db *sql.DB) error {
	sqlStmt := `
CREATE TABLE bill (
  id        INTEGER NOT NULL PRIMARY KEY,
  credit_id INT,
  year      int,
  month     int,
  credit    FLOAT,
  amount    FLOAT,
  balance   FLOAT
)
`
	_, err := db.Exec(sqlStmt)
	return err
}

func Exists(db *sql.DB, table string) bool {
	query := fmt.Sprintf(SQL_EXIT, table)
	r, err := db.Query(query)
	defer r.Close()
	if err != nil {
		return false
	}
	if r.Next(){
		return true
	}
	return false
}

func PatchData(db *sql.DB, table string) (*sql.Rows, error) {
	query := fmt.Sprintf("select * from %s", table)
	r, err := db.Query(query)
	if err != nil {
		log.Errorf("patch data form table %s error", table)
	}
	return r, err
}