package entities

type Catagory struct {
	ID           uint    `gorm:"primarykey"`
	CatagoryName string  `gorm:"not null" json:"catagory_name" form:"catagory_name"`
	Event        []Event `gorm:"foreignKey:CatagoryId;references:ID"`
}
