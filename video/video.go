package video

import "github.com/gin-gonic/gin"

func RegisterRouters(group *gin.RouterGroup) {
	api := group.Group("/video")
	//api.POST("")
	api.POST("", saveVideo)

}
