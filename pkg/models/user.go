package models

import (
	"log"

	"github.com/rakhazufar/go-jwt/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model        
	NamaLengkap string `gorm:"varchar(300)" json:"nama_lengkap"`
	Username string `gorm:"varchar(300)" json:"username"`
	Email string `gorm:"varchar(300)" json:"email"`
	Password string `gorm:"varchar(300)" json:"password"`
}

func init() {
	config.ConnectDatabase()
	db = config.GetDB()
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("error in miggration: %v", err)
	}
}


func CreateUser (user *User) error {
	if err:= db.Create(&user).Error; err != nil {
		return err
	} else {
		return nil
	}
}


func GetUserByUsername (username string) (*User, error) {
	var user User

	result := db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetUserByEmail (email string) (*User, error) {
	var user User
	result := db.Where("email = ?", email).First(&user)
	
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}