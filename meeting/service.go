package meeting

var (
	store Store = &MeetingStore{}
)

type MeetingService struct {
}

func (m *MeetingService) create() (id string, err error) {
	return store.create()
}
