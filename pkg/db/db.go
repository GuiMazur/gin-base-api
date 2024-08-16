package db

import (
	"fmt"
	"gin-base-api/pkg/config"
	"gin-base-api/pkg/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func New() *gorm.DB {
	if dbInstance == nil {
		db, err := setup()
		if err != nil {
			panic(err)
		}
		dbInstance = db
	}

	return dbInstance
}

func setup() (*gorm.DB, error) {
	const retryDelay = 3

	config := config.New()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name)

	var (
		db  *gorm.DB
		err error
	)

	for i := 0; i < 10; i++ {
		fmt.Println("Connecting to database...")
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
		if err == nil {
			break
		} else {
			fmt.Printf("Failed to connect to database, retrying in %v seconds...", retryDelay)
			fmt.Println()
			time.Sleep(retryDelay * time.Second)
		}
	}

	if err != nil {
		return db, err
	}

	db.AutoMigrate(models.ToMigrate()...)
	return db, nil
}
