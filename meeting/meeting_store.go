package meeting

import (
	"math/rand"
	"meeting/global"
	"meeting/model"
	"strconv"
	"time"
)

type MeetingStore struct {
}

func (m *MeetingStore) create(meeting model.MeetingModel) (id string, err error) {
	number := strconv.Itoa(rand.Int())
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	id = number + timestamp
	meeting.Id = id
	err = global.DB.Create(&meeting).Error
	return
}

func (m *MeetingStore) join(id string) (err error) {
	return
}

func (m *MeetingStore) leave(id string) (err error) {
	return
}

func (m *MeetingStore) end(id string) (err error) {
	return
}
