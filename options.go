package main

import (
	"gorm.io/gorm"
)

type Option func(*gorm.DB)

func ID(userId int64) Option {
	return func(db *gorm.DB) {
		db.Where("id = ?", userId)
	}
}

func UserName(userName string) Option {
	return func(db *gorm.DB) {
		db.Where("user_name = ?", userName)
	}
}

func Address(address string) Option {
	return func(db *gorm.DB) {
		db.Where("address = ?", address)
	}
}

func UserId(userId string) Option {
	return func(db *gorm.DB) {
		db.Where("user_id = ?", userId)
	}
}

func Email(email string) Option {
	return func(db *gorm.DB) {
		db.Where("email = ?", email)
	}
}
