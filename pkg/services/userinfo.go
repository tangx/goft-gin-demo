package services

import (
	"strconv"

	"github.com/tangx/goft-gin-demo/pkg/daos"
	"github.com/tangx/goft-gin-demo/pkg/models"
)

type UserInfo struct {
	Dao *daos.UserInfo `inject:"-"`
}

func NewUserInfo() *UserInfo {
	return &UserInfo{}
}

func (u *UserInfo) GetUser(param string) *models.UserInfo {
	user := &models.User{}
	uid, err := strconv.Atoi(param)
	if err != nil {
		name := param
		user = u.Dao.FindUserByName(name)
	} else {
		user = u.Dao.FindUserById(uid)
	}

	if user == nil {
		return nil
	}

	return &models.UserInfo{
		Name:      user.Name,
		Cellphone: user.Cellphone,
	}
}
