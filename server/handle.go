package server

import (
	"net/http"
)

func HomePageHandler(resp http.ResponseWriter, req *http.Request)  {
	http.ServeFile(resp,req,"./static/index.html")
}
