package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dal *Dal = new(Dal)

type Dal struct {
	DB *gorm.DB
}

type User struct {
	// gorm.Model
	Name string
	Age  uint
}

func (dal *Dal) InitDB() error {
	if dal.DB != nil {
		return nil
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	dal.DB = db
	return err
}

func (dal *Dal) CreateUser(u User) error {
	err := dal.DB.Create(&u).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func (dal *Dal) GetUser(name string) (*User, error) {
	u := &User{}
	err := dal.DB.First(u, "name = ?", name).Error
	return u, err
}

func (dal *Dal) DeleteUser(name string) error {
	err := dal.DB.Where("name = ?", name).Delete(&User{}).Error
	return err
}
