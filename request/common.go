package request

import "meeting/model"

type JoinMeetingReq struct {
	Meeting string `json:"meeting,omitempty"`
	model.UserModel
}
