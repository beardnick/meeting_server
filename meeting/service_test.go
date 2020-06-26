package meeting

import (
	"meeting/data"
	"meeting/model"
	"meeting/testutil"
	"testing"

	. "gopkg.in/check.v1"
)

func init() {
	data.Mysql()
	data.Redis()
}

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestCreateEnd(c *C) {
	m, err := service.create(testutil.MockMeetings(1)[0])
	c.Assert(err, IsNil)
	c.Assert(m, NotNil)

	meet, err := service.search(m)
	c.Assert(err, IsNil)
	c.Assert(meet.Id, NotNil)

	err = service.end(meet.Id)
	c.Assert(err, IsNil)

	meet, err = service.search(meet.Id)
	c.Assert(err, NotNil)
}

func (s *MySuite) TestJoinLeave(c *C) {
	m, err := service.create(testutil.MockMeetings(1)[0])
	c.Assert(err, IsNil)
	defer service.end(m)

	users := testutil.MockUsers(10)
	err = service.joinUsers(users, m)
	c.Assert(err, IsNil)

	result, err := service.meetingUsers(m)
	c.Assert(err, IsNil)
	c.Assert(result, NotNil)
	c.Assert(result, HasLen, len(users))
	c.Assert(contains(users, result), Equals, true)
	c.Assert(contains(result, users), Equals, true)
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
