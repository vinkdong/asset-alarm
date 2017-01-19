package main

import (
	"./server"
	"flag"
	"./dbmanager"
	"./log"
	"database/sql"
)

var (
	dbPath = flag.String("db", "", "database path")
)

func main() {
	InitArgs()
	server.Context.DbPath = *dbPath
	log.Info("starting server...")
	InitDb()
	log.Infof("connected database: %s", *dbPath)
	server.Init()
	server.Start()
}

func InitArgs()  {
	flag.Parse()
}

func InitDb() {
	db := &sql.DB{}
	dbmanager.Init(db, server.Context.DbPath)
	if !dbmanager.Exists(db, "credit") {
		dbmanager.InitTables(db)
	}
	if !dbmanager.Exists(db, "record") {
		dbmanager.InitRecordTable(db)
	}
	server.Context.Db = db
}
