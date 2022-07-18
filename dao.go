package main

import (
	"context"

	"gorm.io/gorm"
)

func GetRecord(ctx context.Context, db *gorm.DB, in interface{}, options ...func(option *gorm.DB)) (err error) {
	for _, option := range options {
		option(db)
	}
	err = db.First(in).Error
	return
}
