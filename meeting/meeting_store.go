package meeting

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

type MeetingStore struct {
}

func (m *MeetingStore) create() (id string, err error) {
	number := strconv.Itoa(rand.Int())
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	id = number + timestamp
	client, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Do("LPUSH", "meeting", id)
	if err != nil {
		log.Fatal(err)
	}
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
