package meeting

import "meeting/model"

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

func (m *MeetingService) leave(id string) (err error) {
	return store.leave(id)
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
