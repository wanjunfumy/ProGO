package main

import (
	handles "CNGo/routes"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var router *httprouter.Router

func main() {
	router = httprouter.New()
	for _, v := range handles.Cells {
		v.Init(router)
	}
	log.Fatal(http.ListenAndServe(":8080", router)) // 用了Fatal 就不需要处理err，Fatal中已经有 os.Exit(1)
}
