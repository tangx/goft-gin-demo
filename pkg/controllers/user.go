package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"github.com/tangx/goft-gin-demo/pkg/models"
	"gorm.io/gorm"
)

type User struct {
	// 用于 IoC 注入，
	// 1. 名字随意但必须是公开的。
	// 2. 字段类型必须与 goft.Config 中的配置字段类型一致
	Db *gorm.DB `inject:"-"`
}

func (u *User) Name() string {
	return "User"
}

func (u *User) Build(goft *goft.Goft) {
	goft.Handle("GET", "/user/:uid", u.handlerUser)
	goft.Handle("GET", "/user/:uid/simple", u.handlerUserSimpleQuery)
}

func (u *User) handlerUser(c *gin.Context) goft.Json {
	uid := c.Param("uid")

	user := models.User{}
	// u.GDB.Gdb.Table("users").Where("id = ?", uid).First(&user)
	u.Db.Table("users").Where("id = ?", uid).First(&user)
	return user
}

func (u *User) handlerUserSimpleQuery(c *gin.Context) goft.SimpleQuery {
	uid := c.Param("uid")

	sql := fmt.Sprintf("SELECT * FROM users WHERE id = %s", uid)

	return goft.SimpleQuery(sql)
}
