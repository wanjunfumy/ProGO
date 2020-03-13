package main

import (
	. "fmt"
	"github.com/Unknwon/goconfig"
	"log"
	"net/http"
)

func route(w http.ResponseWriter, r *http.Request) {
	Fprint(w, "let us GO!")
}

/**
 */
func index(w http.ResponseWriter, r *http.Request) {
	Fprint(w, "我是跟目录，我返回你一个String")
}

/*
获取地址配置
*/
func getAddress() (string, string) {
	cof, err := goconfig.LoadConfigFile("configs/config.ini")
	if nil != err {
		log.Print("no config.ini found")
	}
	var address = ""
	var port = ""
	if nil == cof {
		log.Print("config.ini, error!")
		port = ":8080"
	} else {
		address, _ = cof.GetValue("address", "address")
		port, _ = cof.GetValue("address", "port")
	}
	return address, port
}

func main() {
	address, port := getAddress()
	// 我需要一个路由，来管理我所有的请求。
	http.HandleFunc("/id", route)
	http.HandleFunc("/", index)
	// 数据量可能有点大，加载Redis？
	// 加载搜索引擎？GO，你有么？
	// 我需要结合爬虫，去抓一些最新动态、新闻等。
	err := http.ListenAndServe(address+":"+port, nil)
	if nil != err {
		log.Fatal("ListenAndServe", err)
	}
}
