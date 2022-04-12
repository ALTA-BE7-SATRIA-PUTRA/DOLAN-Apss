package utils

import (
	"fmt"
	"group-project/dolan-planner/configs"
	"group-project/dolan-planner/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	InitialMigration(db)
	return db
}

func InitialMigration(db *gorm.DB) {

	fmt.Println("ini jalan")
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Catagory{})
	db.AutoMigrate(&entities.Event{})
	db.AutoMigrate(&entities.Comment{})
	db.AutoMigrate(&entities.Attendees{})
}
