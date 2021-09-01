package db

import (
	"github.com/yongjie0203/go-trade-order/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var Coon *gorm.DB

func InitDB() *gorm.DB {
	dsn := "root:root@(127.0.0.1:3310)/trade?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
	})

	if err != nil {
		log.Print(err.Error())
	}

	sqldb, e := db.DB()
	if e != nil {
		log.Print(e.Error())
	}

	e = sqldb.Ping()
	if e != nil {
		panic(e)
	}
	sqldb.SetMaxIdleConns(100)
	sqldb.SetMaxOpenConns(20)
	Coon = db
	return db
}

func DBUpdate() {
	Coon.AutoMigrate(&model.Order{})
	Coon.AutoMigrate(&model.TransactionPrice{})
	Coon.AutoMigrate(&model.Trade{})
	Coon.AutoMigrate(&model.Token{})
	Coon.AutoMigrate(&model.Pair{})
	Coon.AutoMigrate(&model.PairGroup{})
}

var id int64

// NextId 可更换为Redis实现
func NextId(tableName string) int64 {
	id = id + 1
	return id
}
