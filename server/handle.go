package server

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"../log"
	"../dbmanager"
)

func apiHandler(resp http.ResponseWriter, req *http.Request) {
	uri := req.URL.Path[5:]

	switch uri {
	case "list":
		if checkAccess(resp, req) {
			HandlerList(resp, req)
		}
		break
	default:
		HandlerApiHome(resp, req)
	}

	//urls := req.URL
	r, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("get request body error")
	}
	fmt.Println(string(r))
	resp.Write([]byte(`hello world`))
}

func checkAccess(resp http.ResponseWriter, req *http.Request) bool {
	return true
}

func HandlerApiHome(resp http.ResponseWriter, req *http.Request) {

}

func HandlerList(resp http.ResponseWriter, req *http.Request) {
	db := ReInitDb()
	r, err := dbmanager.PatchData(db, "credit")
	if err != nil {
		return
	}
	var cl = &[]Credit{}
	ParseRowsToCreditList(r, cl)
	js := ParserCreditsToJson(cl)
	respData, err := js.MarshalJSON()
	if err != nil {
		log.Error("convert json to bytes error")
		resp.Write([]byte(`500 SERVER INTERNAL ERROR`))
	}
	resp.Write(respData)
}
