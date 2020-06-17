package meeting

type Store interface {
	create() (id string, err error)
	join(id string) (err error)
	leave(id string) (err error)
	end(id string) (err error)
	//search(id string) (,err error)
}
