package server

import (
	"net/http"
)

func apiHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte(`hello world`))
}
