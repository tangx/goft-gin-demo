package config

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConfig struct {
}

func NewMysqlConfig() *MysqlConfig {
	return &MysqlConfig{}
}

// DB_demo 方法可以任意取名字。
// 	1. 必须是公开方法， 即大写开头
// 	2. 如果存在多个方法， 在初始化时，每个方法都要执行。
// 	2.1. 但在连接使用时， 每次都选择了错误的？ 和顺序无关
func (my *MysqlConfig) DB_demo() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:Mysql12345@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqldb, _ := db.DB()
	sqldb.SetMaxIdleConns(5)
	sqldb.SetMaxOpenConns(10)
	sqldb.SetConnMaxLifetime(time.Second * 30)

	return db
}
