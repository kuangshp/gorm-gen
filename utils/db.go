package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm_gen/dao"
)

const dsn = "root:123456@(localhost:3306)/test001?charset=utf8mb4&parseTime=True&loc=Local"

var DB = dao.Query{}

func InitDB() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true, // 自动创建表的时候不创建外键
		NamingStrategy: schema.NamingStrategy{ // 自动创建表时候表名的配置
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	dao.SetDefault(db)
	q := dao.Q
	DB = *q
}

func init() {
	InitDB()
}
