package main

import (
	"fmt"
	"net/http"
)

var (
	port = "3000"
)

func main() {
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, MainRoutes())
}
