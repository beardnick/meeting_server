package testutil

import (
	"math/rand"
	"meeting/model"
	"strconv"
)

func MockUsers(n int) []model.UserModel {
	users := []model.UserModel{}
	for i := 0; i < n; i++ {
		users = append(users, model.UserModel{
			Id:   strconv.Itoa(rand.Int()),
			Name: strconv.Itoa(rand.Int()),
		})
	}
	return users
}

func MockMeetings(n int) []model.MeetingModel {
	meetings := []model.MeetingModel{}
	for i := 0; i < n; i++ {
		meetings = append(meetings, model.MeetingModel{
			Id:   strconv.Itoa(rand.Int()),
			Name: strconv.Itoa(rand.Int()),
		})
	}
	return meetings
}
