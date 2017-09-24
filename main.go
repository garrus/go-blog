package main

import (
	"fmt"
	"net/http"
	"github.com/garrus/go-blog/core"
	"github.com/garrus/go-blog/config"
)

func main() {

	server := &core.HttpServer{}
	server.Router = core.Router{}
	config.SetRoutes(&server.Router)
	addr := "127.0.0.1:13232";

	fmt.Printf("Server starting ... listening to %s. \n", addr)

	err := http.ListenAndServe(addr, server)
	if err != nil {
		fmt.Errorf("Server stopped. %s\n", err.Error())
	}
}