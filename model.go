package main

import "time"

type Base struct {
	IsDeleted int8      `gorm:"size:4;default:0;comment:删除状态:0-正常 1-删除"`
	Status    int8      `gorm:"size:4;default:1;comment:状态,0-禁用 1-正常"`
	CreatedAt time.Time `gorm:"comment:创建时间"`
	UpdatedAt time.Time `gorm:"comment:更新时间"`
}

type User struct {
	Base
	Id       int64  `gorm:"primaryKey;size:64;autoIncrement;not null;comment:Id"`
	UserName string `gorm:"size:255;comment:用户名"`
	Email    string `gorm:"size:255;comment:邮箱"`
	Address  string `gorm:"size:255;default:null;comment:当前登陆钱包地址"`
}

type UserAddr struct {
	Base
	Id      int64  `gorm:"primaryKey;size:64;autoIncrement;not null;comment:Id"`
	UserId  int64  `gorm:"size:64;comment:用户ID"`
	Address string `gorm:"size:255;comment:钱包地址"`
}
