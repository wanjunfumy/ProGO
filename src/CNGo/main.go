package main

import (
	//handles "CNGo/routes"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	//"log"
	"net/http"
)

var router *httprouter.Router

func Cors() gin.HandlerFunc {
	return func(c gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
func main() {
	//router = httprouter.New()
	//for _, v := range handles.Cells {
	//	v.Init(router)
	//}
	//log.Fatal(http.ListenAndServe(":8080", router)) // 用了Fatal 就不需要处理err，Fatal中已经有 os.Exit(1)

	g := gin.Default()
	g.Use(Cors())
}
