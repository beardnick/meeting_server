package main

import (
	"meeting/data"
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
	router.Run(":9020")
}
