package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用于获取和操作配置

// get config
// @Summary get config
// @Description 获取系统的部分配置
// @Tags 配置
// @Produce  json
// @Success 200
// @Param token query string true "token"
// @Param role query string false "不传入的话，默认值为fuckup"
// @Success 200 {object} User
// @Router /config [get]
func Config(ctx *gin.Context) {
	configToken := ctx.Query("token")
	role := ctx.DefaultQuery("role", "fuckup")
	ret := User{
		ID:      1000,
		Name:    configToken,
		Friends: nil,
		Role:    role,
	}
	ctx.JSON(http.StatusOK, ret)
}

// User 配置字段说明
type User struct {
	ID      int64  `json:"id"`      // 键值
	Name    string `json:"name"`    // configToken
	Friends []User `json:"friends"` // 有多少个与
	Role    string `json:"role"`    // 规则
}
