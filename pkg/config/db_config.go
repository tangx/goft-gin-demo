package config

import (
	"database/sql"
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
func (my *MysqlConfig) DB_demo() *MysqlConfigAdapter {
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
	return &MysqlConfigAdapter{
		db: db,
	}
}

// MysqlConfigAdapter 适配器模式
type MysqlConfigAdapter struct {
	db *gorm.DB
}

// DB 如果要支持 goft.SimpleQuery 返回， 适配器必须包含该方法。
// 	1. 方法签名必须是 `func (apt *Adapter) DB() *sql.DB`
func (mya *MysqlConfigAdapter) DB() *sql.DB {
	sqldb, _ := mya.db.DB()
	return sqldb
}

// Get 此方法可以随意命名， 也可以不存在。
// 	1. 随意命名是是为了在 controller 中能够获取 *gorm.DB， 并使用 *gorm.DB 语法进行操作。
// 		(ex. controller User 中， controllers.User.handlerUser 方法)
// 	2. 不存在
// 		2.1. Adapter 如果能正常匿名内嵌（ ex. Xorm ) 则可以直接调用
// 		2.2. Adapter 字段使用公开字段入 `GormDB *gorm.DB`, 这样就可以直接调用不用方法返回。
func (mya *MysqlConfigAdapter) Get() *gorm.DB {
	return mya.db
}
