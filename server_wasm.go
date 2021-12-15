package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const dir = "web"

var port = 8000

var fs http.Handler

func serve(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Cache-Control", "no-cache")
	if strings.HasSuffix(req.URL.Path, ".wasm") {
		resp.Header().Set("content-type", "application/wasm")
	}
	fs.ServeHTTP(resp, req)
}

func main() {
	if len(os.Args) > 1 {
		if p, err := strconv.Atoi(os.Args[1]); err == nil {
			port = p
		}
	}

	fs = http.FileServer(http.Dir(dir))
	fmt.Printf("Serving Moving Square on localhost: %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), http.HandlerFunc(serve))
}
