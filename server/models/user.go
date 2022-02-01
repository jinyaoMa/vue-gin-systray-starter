package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account    string `gorm:"unique;size:32;<-:create"` // 3 <= len(account) <= 32
	Password   string `gorm:"size:40"`                  // sha1 hex string length = 40
	Name       string `gorm:"size:64"`                  // 3 <= len(name) <= 64
	Permission string ``                                //
	Reserved   bool   `gorm:"default:false;<-:create"`  // mark reserved user account
	Active     bool   `gorm:"default:true;<-:update"`   // mark active user account
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (u *User) Find(db *gorm.DB) (users []User, err error) {
	err = db.Find(&users).Error
	return
}
