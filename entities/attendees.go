package entities

import "gorm.io/gorm"

type Attendees struct {
	gorm.Model
	EventId           uint `gorm:"not null" json:"event_id" form:"event_id"`
	UserId            uint `gorm:"not null" json:"user_id" form:"user_id"`
	TotalParticipants uint `gorm:"not null" json:"total_participants" form:"total_participants"`
	User              User `gorm:"foreignKey:UserId;references:ID"`
}
