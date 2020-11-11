package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	//路由
	router := gin.Default()
	router.GET("/", index)

	// /v1/login 就会匹配到这个组
	v1 := router.Group("/v1")
	if v1 != nil {
		v1.GET("/login", index)
		v1.POST("/submit", index)
		v1.POST("/read", index)
	}
	return router
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "hello gin",
	})
}
