package server

import (
	"database/sql"
	"github.com/bitly/go-simplejson"
	"reflect"
	"bytes"
	"github.com/VinkDong/asset-alarm/log"
	"fmt"
)

func prepareStmt(stmtSql string) (*sql.Tx, *sql.Stmt, error) {
	tx, err := Context.Db.Begin()
	if err != nil {
		log.Error(err)
	}
	stmt, err := tx.Prepare(stmtSql)
	return tx, stmt, err
}

func Interface2map(r interface{})  map[string]string{
	v := reflect.ValueOf(r)
	indirect := reflect.Indirect(v)

	vls := make(map[string]string)
	for i := 0; i < v.Elem().NumField();i++ {
		file_name := indirect.Type().Field(i).Name
		vls[PackToCol(file_name)] = ConvertString(indirect.Field(i))
	}
	return vls
}

func FormatString(buf bytes.Buffer) string {
	b := buf.Bytes()
	return string(b[:len(b)-1])
}

func CommonSave(r interface{}, name string) error {
	val := Interface2map(r)
	stmtSql := GenerateSql(val, name)
	tx, stmt, err := prepareStmt(stmtSql)
	if err != nil {
		log.Error(err)
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec()
	if err != nil{
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	tx.Commit()
	SaveId(r, id)
	return nil
}

func SaveId(x interface{}, id int64) {
	v := reflect.ValueOf(x)
	indirect_type := reflect.Indirect(v).Type()
	el := v.Elem()
	for i := 0; i < el.NumField(); i++ {
		if indirect_type.Field(i).Name == "Id" {
			el.Field(i).SetInt(id)
		}
	}
}

func ParseRowsToCreditList(row *sql.Rows, c_list *[]Credit) {
	defer row.Close()
	for row.Next() {
		var c = &Credit{}
		c.ConvertFormRow(row)
		*c_list = append(*c_list, *c)
	}
}

func ParserCreditsToJson(cl *[]Credit) *simplejson.Json {
	creditList := *cl
	js := simplejson.New()
	js.Set("version", VERSION)
	creditJsonList := make([]interface{}, 0)
	for i := 0; i < len(creditList); i++ {
		jsonPod := creditList[i].ToJson()
		ma := jsonPod.MustMap()
		creditJsonList = append(creditJsonList, ma)
	}
	js.Set("credits", creditJsonList)
	return js
}

func GenerateSql(val map[string]string, name string) string {
	var stmtSql string
	if val["id"] == "0" {
		var header bytes.Buffer
		var values bytes.Buffer
		for k, v := range val {
			if k == "id" {
				continue
			}
			header.WriteString(k + ",")
			values.WriteString(v + ",")
		}
		stmtSql = fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", name, FormatString(header), FormatString(values))
	} else {
		var values bytes.Buffer
		values.WriteString("UPDATE record SET ")
		for k, v := range val {
			if k == "id" {
				continue
			}
			values.WriteString(fmt.Sprintf("%s=%s,", k, v))
		}
		stmtSql = FormatString(values) + " where id = " + val["id"]
	}
	return stmtSql
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