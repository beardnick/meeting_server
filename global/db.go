package global

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

var (
	DB    *gorm.DB
	REDIS redis.Conn
)
