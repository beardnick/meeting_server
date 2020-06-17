package meeting

import "meeting/model"

type Store interface {
	create(meeting model.MeetingModel) (id string, err error)
	join(id string) (err error)
	leave(id string) (err error)
	end(id string) (err error)
	//search(id string) (,err error)
}
