package meeting

import "meeting/model"

type Store interface {
	create(meeting model.MeetingModel) (id string, err error)
	join(user model.UserModel, id string) (err error)
	leave(mid string, uid string) (err error)
	end(id string) (err error)
	search(id string) (meeting model.MeetingModel, err error)
	meetingUsers(id string) (users []model.UserModel, err error)
}
