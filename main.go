package main

import (
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
	router.Run(":9020")
}
