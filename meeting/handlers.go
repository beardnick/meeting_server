package meeting

import (
	"log"
	"meeting/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	service = MeetingService{}
)

func hello(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		"hello world",
	)
}

func createMeeting(c *gin.Context) {
	id, err := service.create()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithData(gin.H{"meeting": id}, c)
}

func stopMeeting(c *gin.Context) {
}

func joinMeeting(c *gin.Context) {
	meeting := c.Query("meeting")
	log.Println("meeting:", meeting)
	c.JSON(
		http.StatusOK,
		gin.H{
			"meeting": meeting,
		},
	)
}
