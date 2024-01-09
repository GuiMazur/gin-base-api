package db

import (
	"fmt"
	"gin-base-api/config"
	"gin-base-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
  
  func SetupDB(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	  return db, err
	}
	
	db.AutoMigrate(&models.User{})

	return db, nil
  }