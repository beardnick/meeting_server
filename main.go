package main

import (
	"meeting/data"
	"meeting/global"
	"meeting/meeting"
	"meeting/middleware"
	"meeting/video"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Cors())
	video.RegisterRouters(&router.RouterGroup)
	meeting.RegisterRouters(&router.RouterGroup)
	data.Mysql()
	data.Redis()
	defer global.DB.Close()
	defer global.REDIS.Close()
	router.Run(":9020")
}
