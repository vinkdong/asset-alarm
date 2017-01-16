package server

import (
	"net/http"
	"../log"
	"../dbmanager"
	"github.com/bitly/go-simplejson"
)

func apiHandler(resp http.ResponseWriter, req *http.Request) {
	uri := req.URL.Path[5:]

	switch uri {
	case "list":
		if checkAccess(resp, req) {
			HandlerList(resp, req)
		}
		break
	case "item/add":
		if checkAccess(resp, req) {
			HandLerAddItem(resp, req)
		}
	case "item/del":
		if checkAccess(resp, req) {
			HandLerDelItem(resp, req)
		}
	case "item/update":
		if checkAccess(resp, req) {
			HandLerUpdateItem(resp, req)
		}
	default:
		HandlerApiHome(resp, req)
	}
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
	resp.Header().Set("content-type","application/json")
	resp.Write(respData)
}

func HandLerAddItem(resp http.ResponseWriter, req *http.Request) {
	js, err := simplejson.NewFromReader(req.Body)
	if err != nil {
		log.Error("catch add item error cant convert to json object")
	}
	version := js.Get("version").MustString()
	if version != VERSION {
		resp.Write([]byte(`{"error":"api version is not support"}`))
		return
	}
	creditJson := js.Get("credit")
	var c = Credit{}
	c.ConvertFromJson(creditJson)
	c.Save()
	resp.Header().Set("content-type", "application/json")
	resp.Write([]byte(`{"success":true}`))
}

func HandLerDelItem(resp http.ResponseWriter, req *http.Request) {

}

func HandLerUpdateItem(resp http.ResponseWriter, req *http.Request) {

}

