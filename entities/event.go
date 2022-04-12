package entities

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	UserId          uint        `gorm:"not null" json:"user_id" form:"user_id"`
	CatagoryId      uint        `gorm:"not null" json:"catagory_id" form:"catagory_id"`
	NameEvent       string      `gorm:"not null" json:"name_event" form:"name_event"`
	HostedBy        string      `gorm:"not null" json:"hosted_by" form:"hosted_by"`
	MaxParticipants uint        `gorm:"not null" json:"max_participants" form:"max_participants"`
	Date            time.Time   `gorm:"not null" json:"date" form:"date"`
	Location        string      `gorm:"not null" json:"location" form:"location"`
	DetailEvent     string      `gorm:"not null" json:"detail_event" form:"detail_event"`
	UrlImage        string      `gorm:"not null" json:"url_image" form:"url_image"`
	Attendees       []Attendees `gorm:"foreignKey:EventId;references:ID"`
	Comment         []Comment   `gorm:"foreignKey:EventId;references:ID"`
}
