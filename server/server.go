package server

import (
	"net/http"
	"database/sql"
	"../dbmanager"
	"../log"
)

var Context = Alarm{}

func Init() {
	http.HandleFunc("/api/", apiHandler)
}

func Start() {
	http.ListenAndServe(":8001", nil)
}

func AddHandler(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, handlerFunc)
}

func ReInitDb() *sql.DB {
	var db = &sql.DB{}
	if Context.Db == nil {
		err := dbmanager.Init(db, Context.DbPath)
		if err != nil {
			log.Error(err)
		}
	} else {
		db = Context.Db
	}
	return db
}