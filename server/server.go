package server

import (
	"net/http"
)

func Init() {
	http.HandleFunc("/api/", apiHandler)
}

func Start() {
	http.ListenAndServe(":8001", nil)
}

func AddHandler(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, handlerFunc)
}