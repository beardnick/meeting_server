package data

import (
	"log"
	"meeting/global"
	"meeting/model"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Mysql() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/meeting?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("init mysql root:123456@(127.0.0.1:3306)/meeting?charset=utf8&parseTime=True&loc=Local")
	global.DB = db
	// 自动创建表
	db.AutoMigrate(&model.MeetingModel{})
}

func Redis() {
	db, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("init redis 127.0.0.1:6379")
	global.REDIS = db
}
