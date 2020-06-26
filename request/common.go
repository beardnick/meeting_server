package request

import "meeting/model"

type JoinMeetingReq struct {
	Meeting string `json:"meeting,omitempty"`
	model.UserModel
}

type LeaveMeetingReq struct {
	Meeting string `json:"meeting,omitempty"`
	User    string `json:"username,omitempty`
}
