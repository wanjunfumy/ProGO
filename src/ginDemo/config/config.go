package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	ctx.JSON(http.StatusOK, ret(configToken, role))
}

// get config1
// @Summary get config1
// @Description 获取系统的部分配置1
// @Tags 配置
// @Produce  json
// @Success 200
// @Param token query string true "token"
// @Param role query string false "不传入的话，默认值为fuckup"
// @Success 200 {object} User
// @Router /config/:id [get]
func Config1(ctx *gin.Context) {
	configToken := ctx.Query("token")
	id := ctx.Param("id")
	role := ctx.DefaultQuery("role", "fuckup")
	ret := ret(configToken, role)
	int64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		int64 = 0
	}
	ret.ID = int64
	ctx.JSON(http.StatusOK, ret)
}

func ret(configToken string, role string) (user User) {
	user = User{
		ID:      1000,
		Name:    configToken,
		Friends: nil,
		Role:    role,
	}
	return
}

type User struct {
	ID      int64  `json:"id"`      // 键值
	Name    string `json:"name"`    // configToken
	Friends []User `json:"friends"` // 有多少个与
	Role    string `json:"role"`    // 规则
}
