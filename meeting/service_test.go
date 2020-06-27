package meeting

import (
	"meeting/data"
	"meeting/model"
	"meeting/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	data.Mysql()
	data.Redis()
}

func TestCreateEnd(t *testing.T) {
	m, err := service.create(testutil.MockMeetings(1)[0])
	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(m)

	meet, err := service.search(m)
	assert.Nil(err)
	assert.NotNil(meet.Id)

	err = service.end(meet.Id)
	assert.Nil(err)

	meet, err = service.search(meet.Id)
	assert.NotNil(err)
}

func TestJoinLeave(t *testing.T) {
	assert := assert.New(t)
	m, err := service.create(testutil.MockMeetings(1)[0])
	assert.Nil(err)
	defer service.end(m)

	users := testutil.MockUsers(10)
	err = service.joinUsers(users, m)
	assert.Nil(err)

	result, err := service.meetingUsers(m)
	assert.Nil(err)
	assert.NotNil(result)
	assert.Len(result, len(users))
	// users == result
	assert.True(contains(users, result))
	assert.True(contains(result, users))
}

func contains(a []model.UserModel, b []model.UserModel) bool {
	for _, i := range b {
		contain := false
		for _, j := range a {
			if i == j {
				contain = true
				break
			}
		}
		if !contain {
			return false
		}
	}
	return true
}
