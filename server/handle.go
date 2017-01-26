package server

import (
	"net/http"
	"strings"
)

func HomePageHandler(resp http.ResponseWriter, req *http.Request)  {
	requestUrl := req.URL.Path
	if strings.HasSuffix(requestUrl, ".js") {
		http.ServeFile(resp,req,"."+requestUrl)
	} else {
		http.ServeFile(resp,req,"./static/index.html")
	}
}
