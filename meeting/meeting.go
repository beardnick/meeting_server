package meeting

import "github.com/gin-gonic/gin"

func RegisterRouters(group *gin.RouterGroup) {
	api := group.Group("/meeting")
	api.POST("create", createMeeting)
	api.POST("join", joinMeeting)
	api.POST("leave", leaveMeeting)
	api.POST("close", closeMeeting)
	api.GET("users", meetingUsers)
}
