package models

import "gorm.io/gorm"

const TableNameUser = "user"

type User struct {
	gorm.Model
	Id       int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name     string `gorm:"column:name;size:256" json:"name"`
	Email    string `gorm:"column:email;uniqueIndex;size:256" json:"email"`
	Password string `gorm:"column:password;size:256" json:"password"`
}

func (*User) TableName() string {
	return TableNameUser
}
