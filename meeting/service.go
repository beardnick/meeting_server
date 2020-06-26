package meeting

import (
	"meeting/model"
)

var (
	store Store = &MeetingStore{}
)

type MeetingService struct {
}

func (m *MeetingService) create(meeting model.MeetingModel) (id string, err error) {
	return store.create(meeting)
}

func (m *MeetingService) join(user model.UserModel, id string) (err error) {
	return store.join(user, id)
}

func (m *MeetingService) joinUsers(users []model.UserModel, id string) (err error) {
	// 如果没有全部插入成功就要回滚已经成功的
	rollBack := []string{}
	for _, v := range users {
		err = store.join(v, id)
		if err != nil {
			break
		}
		rollBack = append(rollBack, v.Id)
	}
	if err == nil {
		return
	}
	for _, v := range rollBack {
		err = store.leave(id, v)
		if err != nil {
			return
		}
	}
	return
}

func (m *MeetingService) leave(mid string, uid string) (err error) {
	return store.leave(mid, uid)
}

func (m *MeetingService) end(id string) (err error) {
	return store.end(id)
}

func (m *MeetingService) search(id string) (meeting model.MeetingModel, err error) {
	return store.search(id)
}

func (m *MeetingService) meetingUsers(id string) (users []model.UserModel, err error) {
	return store.meetingUsers(id)
}
