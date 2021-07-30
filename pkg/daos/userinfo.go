package daos

import (
	"github.com/tangx/goft-gin-demo/pkg/models"
	"gorm.io/gorm"
)

type UserInfo struct {
	Db *gorm.DB `inject:"-"`
}

func NewUserInfo() *UserInfo {
	return &UserInfo{}
}

func (u UserInfo) FindUserById(uid int) *models.User {

	user := &models.User{}
	tx := u.Db.Table("users").Where("id = ?", uid).First(user)
	if tx.Error != nil {
		return nil
	}

	return user
}

func (u UserInfo) FindUserByName(name string) *models.User {

	user := &models.User{}
	tx := u.Db.Table("users").Where("name = ?", name).First(user)
	if tx.Error != nil {
		return nil
	}

	return user
}
