package config

import (
	"github.com/tangx/goft-gin-demo/pkg/daos"
	"github.com/tangx/goft-gin-demo/pkg/services"
)

type ServicesConfig struct {
}

func NewServicesConfig() *ServicesConfig {
	return &ServicesConfig{}
}

func (sc *ServicesConfig) UserInfoDao() *daos.UserInfo {
	return daos.NewUserInfo()
}

func (sc *ServicesConfig) UserInfoService() *services.UserInfo {
	return services.NewUserInfo()
}
