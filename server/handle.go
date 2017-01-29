package server

import (
	"net/http"
	"strings"
)

func CheckStaticResources(path string, suffix ...string) bool {
	for i := 0; i < len(suffix); i++ {
		if strings.HasSuffix(path, suffix[i]) {
			return true
		}
	}
	return false
}

func HomePageHandler(resp http.ResponseWriter, req *http.Request) {
	requestUrl := req.URL.Path
	if CheckStaticResources(requestUrl, "js", "png", "jpg", "css") {
		http.ServeFile(resp, req, "."+requestUrl)
	} else {
		http.ServeFile(resp,req,"./static/index.html")
	}
}
