package server

import (
	"database/sql"
	"../log"
	"fmt"
	"github.com/bitly/go-simplejson"
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
	Db      *sql.DB
	Credits []Credit
	DbPath  string
}

func prepareStmt(stmtSql string) (*sql.Tx, *sql.Stmt, error) {
	tx, err := Context.Db.Begin()
	if err != nil {
		log.Error(err)
	}
	stmt, err := tx.Prepare(stmtSql)
	return tx, stmt, err
}

func (c *Credit) Save() {
	var stmtSql string
	if c.Id == 0 {
		stmtSql = "insert into credit(name,icon,amount,debit,balance,account_date,repayment_date) values(?,?,?,?,?,?,?)"
	} else {
		stmtSql = "update credit set name = ?, icon =? ,amount =?,debit =?,balance =?,account_date =?,repayment_date =? where id =" +
			string(c.Id)
	}
	tx, stmt, err := prepareStmt(stmtSql)
	if err != nil {
		log.Error(err)
	}
	defer stmt.Close()
	r, err := stmt.Exec(c.Name, c.Icon, c.Amount, c.Debit, c.Balance, c.Account_date, c.Repayment_date)
	id, err := r.LastInsertId()
	tx.Commit()
	c.Id = int8(id)
}

func (c *Credit) ConventFormRow(rows *sql.Rows) error {
	var err error
	if err = rows.Scan(&c.Id, &c.Name, &c.Icon, &c.Amount, &c.Debit, &c.Balance, &c.Account_date, &c.Repayment_date); err != nil {
		log.Error("convert rows to credit object error")
	}
	return err
}

func (c *Credit) ToJsonString() string {
	jsonStr := fmt.Sprintf(`{
	"name":"%s",
	"icon":"%s",
	"amount":%f,
	"debit":%f,
	"balance":%f,
	"account_date":%d,
	"repayment_date":%d,
	"id":%d
}`, c.Name, c.Icon, c.Amount, c.Debit, c.Balance, c.Account_date, c.Repayment_date, c.Id)
	return jsonStr
}

func (c *Credit) ToJson() *simplejson.Json {
	jsData := c.ToJsonString()
	js, err := simplejson.NewJson([]byte(jsData))
	if err != nil {
		log.Error("parser credit credit to json error")
		return nil
	}
	return js
}