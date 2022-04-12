package entities

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	EventId uint   `gorm:"not null" json:"event_id" form:"event_id"`
	UserId  uint   `gorm:"not null" json:"user_id" form:"user_id"`
	Comment string `gorm:"not null" json:"comment" form:"comment"`
	User    User   `gorm:"foreignKey:UserId;references:ID"`
}
