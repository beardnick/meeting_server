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
