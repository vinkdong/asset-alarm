package server

import (
	"database/sql"
	"github.com/VinkDong/asset-alarm/log"
	"github.com/bitly/go-simplejson"
	"strconv"
	"fmt"
	"reflect"
	"bytes"
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

type Record struct {
	Id       int64
	CreditId int64
	Type     string
	Amount   float64
	Credit   float64
	Debit    float64
	Time     string
}

type Bill struct {
	Id       int64
	CreditId int64
	Year     int
	Month    int
	Amount   float64
	Balance  float64
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

func CommonSave(r interface{}) error {
	v := reflect.ValueOf(r)
	indirect := reflect.Indirect(v)
	//fmt.Println(indirectType.Field(1).Name,ind.Field(1).Int())

	vls := make(map[string]string)
	for i := 0; i < v.NumField(); i++ {
		file_name := indirect.Type().Field(i).Name
		vls[file_name] = ConvertString(indirect.Field(i))
	}

	return nil
}

func PackToCol(key string) string {
	var buf bytes.Buffer
	for index, v := range []byte(key) {
		if v >= 65 && v <= 90 {
			if index > 0 {
				buf.WriteByte('_')
			}
			buf.WriteByte(v + 32)
		} else {
			buf.WriteByte(v)
		}
	}
	return buf.String()
}

func ConvertString(i reflect.Value) string {
	switch i.Kind() {
	case reflect.String:
		return fmt.Sprintf("\"%s\"", i.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", i.Int())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%.2f", i.Float())
	case reflect.Bool:
		if i.Bool() {
			return "true"
		}
		return "false"
	}
	return ""
}