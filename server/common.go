package server

import (
	"database/sql"
	"github.com/bitly/go-simplejson"
)

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