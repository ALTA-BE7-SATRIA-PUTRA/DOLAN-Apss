package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string      `gorm:"not null" json:"name" form:"name"`
	City      string      `gorm:"not null" json:"city" form:"city"`
	UrlImage  string      `gorm:"not null;default:https://th.bing.com/th/id/R.99537a7e756793f5f03580876cee89a7?rik=Lh%2fc3JEffloN7g&riu=http%3a%2f%2fgetdrawings.com%2ffree-icon-bw%2fprofile-png-icon-17.png&ehk=m4AnM9fgkGBAl3fGtmV1jtB4Lz1mjEIB1arLLHcVHjc%3d&risl=&pid=ImgRaw&r=0" json:"url_image" form:"url_image"`
	Email     string      `gorm:"unique;not null" json:"email" form:"email"`
	Password  string      `gorm:"not null" json:"password" form:"password"`
	Event     []Event     `gorm:"foreignKey:UserId;references:ID"`
	Attendees []Attendees `gorm:"foreignKey:UserId;references:ID"`
	Comment   []Comment   `gorm:"foreignKey:UserId;references:ID"`
}
