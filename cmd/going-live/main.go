package main

import (
	"fmt"
	"net/http"
)

var (
	port = "3001"
)

func main() {
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, MainRoutes())
}
