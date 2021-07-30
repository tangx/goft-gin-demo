package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"github.com/tangx/goft-gin-demo/pkg/services"
)

type UserInfo struct {
	Service *services.UserInfo `inject:"-"`
}

func (uf *UserInfo) Name() string {
	return "UserInfo"
}

func (uf *UserInfo) Build(goft *goft.Goft) {
	goft.Handle("GET", "/userinfo/:param", uf.handlerGetUserByID)
}

func (uf *UserInfo) handlerGetUserByID(c *gin.Context) goft.Json {
	param := c.Param("param")

	user := uf.Service.GetUser(param)
	if user != nil {
		return user
	}

	return gin.H{"error": "404, user not found"}

}
