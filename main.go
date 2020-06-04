package main

import (
	"meeting/middleware"
	"meeting/video"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Cors())
	video.RegisterRouters(&router.RouterGroup)
	router.Run(":9020")
}
