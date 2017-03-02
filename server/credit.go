package server

import (
	"database/sql"
	"github.com/VinkDong/asset-alarm/log"
	"github.com/bitly/go-simplejson"
	"strconv"
	"fmt"
)

type Credit struct {
	Name           string
	Icon           string
	Credit         float64
	Debit          float64
	Balance        float64
	Account_date   int8
	Repayment_date int8
	Id             int64
}

type Alarm struct {
	Db      *sql.DB
	Credits []Credit
	DbPath  string
}

func (c *Credit) Save() error {
	var stmtSql string
	if c.Id == 0 {
		stmtSql = "insert into credit(name,icon,credit,debit,balance,account_date,repayment_date) values(?,?,?,?,?,?,?)"
	} else {
		stmtSql = "update credit set name = ?, icon =? ,credit =?,debit =?,balance =?,account_date =?,repayment_date =? where id = " + strconv.FormatInt(c.Id,4)
	}
	tx, stmt, err := prepareStmt(stmtSql)
	if err != nil {
		log.Error(err)
		return err
	}
	defer stmt.Close()
	r, err := stmt.Exec(c.Name, c.Icon, c.Credit, c.Debit, c.Balance, c.Account_date, c.Repayment_date)
	if err != nil {
		tx.Commit()
		return err
	}
	id, err := r.LastInsertId()
	if err != nil {
		tx.Commit()
		return err
	}
	tx.Commit()
	c.Id = id
	return nil
}

func (c *Credit) Browse(id int64) {
	stmtSql := `select * from credit where id = ?`
	r, err := Context.Db.Query(stmtSql, id)
	if err != nil {
		log.Errorf("Can't get credit id %d", id)
		return
	}
	if !r.Next(){
		return
	}
	err = c.ConvertFormRow(r)
	r.Close()
	if err != nil {
		log.Errorf("browse credit %d fail", id)
	}
}

func (c *Credit) ConvertFormRow(rows *sql.Rows) error {
	var err error
	if err = rows.Scan(&c.Id, &c.Name, &c.Icon, &c.Credit, &c.Debit, &c.Balance, &c.Account_date, &c.Repayment_date); err != nil {
		log.Error("convert rows to credit object error")
	}
	return err
}

func (c *Credit) ConvertFromJson(js *simplejson.Json) {
	c.Name = js.Get("name").MustString()
	c.Icon = js.Get("icon").MustString()
	c.Credit = js.Get("credit").MustFloat64()
	c.Debit = js.Get("debit").MustFloat64()
	c.Balance = js.Get("balance").MustFloat64()
	c.Account_date = int8(js.Get("account_date").MustInt())
	c.Repayment_date = int8(js.Get("repayment_date").MustInt())
	c.Id = int64(js.Get("id").MustInt())
}

func (c *Credit) ToJsonString() string {
	jsonStr := fmt.Sprintf(`{
	"name":"%s",
	"icon":"%s",
	"credit":%f,
	"debit":%f,
	"balance":%f,
	"account_date":%d,
	"repayment_date":%d,
	"id":%d
}`, c.Name, c.Icon, c.Credit, c.Debit, c.Balance, c.Account_date, c.Repayment_date, c.Id)
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