package server

import (
	"database/sql"
	"github.com/VinkDong/asset-alarm/log"
	"github.com/bitly/go-simplejson"
	"strconv"
)

type Record struct {
	Id       int64
	CreditId int64
	Type     string
	Amount   float64
	Credit   float64
	Debit    float64
	Time     string
}

func (c *Record) ConvertFormRow(rows *sql.Rows) error {
	var err error
	if err = rows.Scan(&c.Id, &c.CreditId, &c.Type, &c.Amount, &c.Credit, &c.Debit, &c.Time); err != nil {
		log.Error("convert rows to credit object error")
	}
	return err
}

func (r *Record) ConvertFromJson(js *simplejson.Json) {
	r.Id = js.Get("id").MustInt64()
	r.CreditId = js.Get("cid").MustInt64()
	r.Type = js.Get("type").MustString()
	r.Amount = js.Get("amount").MustFloat64()
	r.Credit = js.Get("credit").MustFloat64()
	r.Debit = js.Get("debit").MustFloat64()
	r.Time = js.Get("time").MustString()
}

func (r *Record) Save() error{
	var stmtSql string
	if r.Id == 0 {
		stmtSql = "INSERT INTO record(credit_id,type,amount,credit,debit,time) VALUES (?,?,?,?,?,?);"
	} else {
		stmtSql = "UPDATE record SET credit_id = ?,type = ? ,amount = ? ,credit = ? ,debit = ?, time = ? where id = " +
			strconv.FormatInt(r.Id, 8)
	}
	c := &Credit{}
	c.Browse(r.CreditId)
	c.Debit += r.Amount
	c.Balance += r.Amount
	err := c.Save()
	if err != nil {
		return err
	}
	tx, stmt, err := prepareStmt(stmtSql)
	if err != nil {
		log.Error(err)
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(r.CreditId, r.Type, r.Amount, r.Credit, r.Debit, r.Time)
	if err != nil{
		return err
	}
	id, err := result.LastInsertId()
	if err != nil{
		return err
	}
	tx.Commit()
	r.Id = id
	return nil
}