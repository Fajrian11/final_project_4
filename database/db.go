package database

import (
	"final_project_4/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DBinit(dbHost, dbPort, dbUsername, dbPassword, dbName string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", dbUsername, dbPassword, dbHost, dbPort, dbName)

	fmt.Println(dsn)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf("ERROR: Failed to connect to Database -> %v\n", err)
	}
	db.AutoMigrate(models.User{}, models.Product{}, models.Category{}, models.TransactionHistory{})

	// database seeding
	var users = []models.User{}
	db.Where("role = ?", "admin").Find(&users)
	// fmt.Println(len(users))
	if len(users) == 0 {
		db.Create(&models.User{
			Full_name: "admin",
			Email:     "admin@gmail.com",
			Password:  "admin123",
			Role:      "admin",
		})
	}

	return db
}
