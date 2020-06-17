package global

import "strings"

var (
	SPLITER = ":"
	MEETING = "meeting"
)

func Meeting(id string) string {
	return strings.Join([]string{MEETING, id}, SPLITER)
}
