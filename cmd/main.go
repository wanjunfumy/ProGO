package main

import (
	. "fmt"
	"github.com/Unknwon/goconfig"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type String struct {
	string
	value string
}

func (s String) len() int {
	return len(s.string)
}

func route(w http.ResponseWriter, r *http.Request) {
	Fprint(w, "let us GO!")
}

/**
 */
func index(w http.ResponseWriter, r *http.Request) {
	Fprint(w, "我是跟目录，我返回你一个String")
}

func newAccount(w http.ResponseWriter, r *http.Request) {
	Fprint(w, "我需要很多参数的哦")
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

/**
 */
func Chann(ch chan int, stopCh chan bool) {
	var i int
	i = 10
	for j := 0; j < 10; j++ {
		ch <- i
		time.Sleep(time.Second)
	}
	stopCh <- true
}

func mapp() {
	var a map[string]string = map[string]string{
		"1": "2",
	}
	a["2"] = "2"
	a["2"] = "3"
	a["3"] = "5"
	Println(a)
}

func main() {
	var a = [...]int{1, 5, 3, 2, 6}
	sort.Ints(a[:])
	Println(a)
	mapp()
	Fprintf()
	/*ch := make(chan int)
	c := 0
	stopCh := make(chan bool)

	go Chann(ch, stopCh)

	for {
		select {
		case c = <-ch:
			Println("Receive", c)
			Println("channel")
		case s := <-ch:
			Println("Receive", s)
		case _ = <-stopCh:
			goto end
		}
	}
	end:

	loop()*/

	address, port := getAddress()
	// 我需要一个路由，来管理我所有的请求。
	http.HandleFunc("/id", route)
	http.HandleFunc("/newAccount", newAccount)
	http.HandleFunc("/", index)
	// 数据量可能有点大，加载Redis？
	// 加载搜索引擎？GO，你有么？
	// 我需要结合爬虫，去抓一些最新动态、新闻等。
	err := http.ListenAndServe(address+":"+port, nil)
	if nil != err {
		log.Fatal("ListenAndServe", err)
	}

}

func loop() {
	//for i := 0; i < 49; i++ {
	//	Println("我打印：", i)
	//	if i % 3 == 0 {
	//		println("我叫万新蕊，嘿嘿！")
	//	}
	//	if i % 7 == 0 {
	//		println()
	//		println("是不是很好玩？")
	//		println()
	//	}
	//	time.Sleep(time.Second)
	//}

	pipe := make(chan int, 2)
	pipe <- 5
	pipe <- 4
	a := len(pipe)
	for i := 0; i < a; i++ {
		s := <-pipe
		println(s)
	}

	f := makeSuffix(".bmp")
	println(f("test"))
	println(f("panic"))
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			suffix = ".png" // 如果是传参进来的，这里是无法改变初始化的值
			return name
		}
		return name + suffix
	}
}
