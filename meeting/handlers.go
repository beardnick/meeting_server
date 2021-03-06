package meeting

import (
	"fmt"
	"meeting/model"
	"meeting/request"
	"meeting/response"
	"meeting/utils"
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
	meeting := model.MeetingModel{}
	err := c.ShouldBindJSON(&meeting)
	if err != nil {
		response.FailWithMessage(utils.ParamError(err).Error(), c)
		return
	}
	id, err := service.create(meeting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(gin.H{"meeting": id}, c)
}

func joinMeeting(c *gin.Context) {
	req := request.JoinMeetingReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(utils.ParamError(err).Error(), c)
		return
	}
	model, err := service.search(req.Meeting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if &model == nil {
		response.FailWithMessage(fmt.Sprintf("no meeting %s", req.Meeting), c)
		return
	}
	err = service.join(req.UserModel, req.Meeting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func leaveMeeting(c *gin.Context) {
	req := request.LeaveMeetingReq{}
	c.ShouldBindJSON(&req)
	err := service.leave(req.Meeting, req.User)
	if err != nil {
		response.FailWithMessage(utils.ParamError(err).Error(), c)
		return
	}
	response.Ok(c)
}

func closeMeeting(c *gin.Context) {
	meeting := c.Query("meeting")
	err := service.end(meeting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func meetingUsers(c *gin.Context) {
	meeting := c.Query("meeting")
	users, err := service.meetingUsers(meeting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(gin.H{"users": users}, c)
}
