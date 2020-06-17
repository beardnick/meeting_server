package model

type MeetingModel struct {
	Id   string `json:"id,omitempty" gorm:"not null;unique;primary_key;" gorm:"comment:'会议ID'"`
	Name string `json:"name,omitempty" gorm:"default:''" gorm:"comment:'会议名'"`
}
