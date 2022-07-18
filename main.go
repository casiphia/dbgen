package main

import (
	"context"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(
		new(User),
		new(UserAddr),
	)
	if err != nil {
		log.Panic(err)
	}
	db = db.Debug().Where("1 = 1")
	ctx := context.Background()
	var user User
	if err := GetRecord(ctx,
		db,
		&user,
		ID(1),
	); err != nil {
		log.Fatal(err)
	}
	log.Println(user)
}
