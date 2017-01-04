package server

import (
	"database/sql"
	"log"
)

type Credit struct {
	Name           string
	Icon           string
	Amount         float64
	Debit          float64
	Balance        float64
	Account_date   int8
	Repayment_date int8
	Id             int8
}

type Alarm struct {
	Db *sql.DB
}

func (c *Credit) Save() {
	var stmtSql string
	if c.Id == 0 {
		stmtSql = "insert into credit(name,icon,amount,debit,balance,account_date,repayment_date) values(?,?,?,?,?,?,?)"
	} else {
		stmtSql = "update credit set name = ?, icon =? ,amount =?,debit =?,balance =?,account_date =?,repayment_date =? where id =" +
			string(c.Id)
	}
	tx, err := Context.Db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(stmtSql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	r, err := stmt.Exec(c.Name, c.Icon, c.Amount, c.Debit, c.Balance, c.Account_date, c.Repayment_date)
	id, err := r.LastInsertId()
	tx.Commit()
	c.Id = int8(id)
}