package dbmanager

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
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
amount float,
debit float,
balance float,
account_date int,
repayment_date int
)`
	_, err := db.Exec(sqlStmt)
	return err
}

func Exists(db *sql.DB, table string) bool {
	query := fmt.Sprintf(SQL_EXIT, table)
	r, err := db.Query(query)
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
	return db.Query(query)
}