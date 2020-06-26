package meeting

import (
	"encoding/json"
	"math/rand"
	"meeting/global"
	"meeting/model"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
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

func (m *MeetingStore) join(user model.UserModel, id string) (err error) {
	data, err := json.Marshal(user)
	_, err = global.REDIS.Do("HSET", global.Meeting(id), user.Id, data)
	return
}

func (m *MeetingStore) leave(mid string, uid string) (err error) {
	_, err = global.REDIS.Do("HDEL", global.Meeting(mid), uid)
	return
}

func (m *MeetingStore) end(id string) (err error) {
	err = global.DB.Delete(&model.MeetingModel{Id: id}).Error
	if err != nil {
		return
	}
	_, err = global.REDIS.Do("DEL", global.Meeting(id))
	return
}

func (m *MeetingStore) search(id string) (meeting model.MeetingModel, err error) {
	db := global.DB.Model(&model.MeetingModel{})
	err = db.Where("id = ?", id).First(&meeting).Error
	return
}

func (m *MeetingStore) meetingUsers(id string) (users []model.UserModel, err error) {
	users = []model.UserModel{}
	strMap, err := redis.StringMap(global.REDIS.Do("HGETALL", global.Meeting(id)))
	if err != nil {
		return
	}
	user := model.UserModel{}
	for _, v := range strMap {
		err = json.Unmarshal([]byte(v), &user)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return
}
