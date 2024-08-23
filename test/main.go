package main

import (
	models "ldchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	// utils.SystemInit()
	MysqlTest()
}

func MysqlTest() {
	dsn := "root:root@tcp(127.0.0.1:3306)/ldchat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})

	row := []models.UserBasic{
		{
			Model:    gorm.Model{ID: 1},
			Name:     "test1",
			Password: "123456",
		},
		{
			Model:    gorm.Model{ID: 2},
			Name:     "test2",
			Password: "123456",
		},
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&row)
}
