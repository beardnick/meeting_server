package meeting

import "github.com/gin-gonic/gin"

func RegisterRouters(group *gin.RouterGroup) {
	api := group.Group("/meeting")
	//api.POST("")
	api.POST("", hello)
	api.POST("create", createMeeting)
	api.GET("join", joinMeeting)

}
