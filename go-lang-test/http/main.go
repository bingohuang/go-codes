package main

import (
	"net/http"
)

func main() {

	server := &http.Server{}
	server.ListenAndServe()
	server.Close()
	//server.Shutdown()
}
