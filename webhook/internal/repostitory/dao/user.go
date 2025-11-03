package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}
func (dao *UserDao) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	return dao.db.Create(&u).Error
}
func (dao *UserDao) Update(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Ctime = now
	return dao.db.Model(&User{}).Where("id = ?", u.Id).UpdateColumns(&u).Error
}

type User struct {
	Id          int64  `gorm:"primary_key,auto_increment"`
	Email       string `gorm:"unique"`
	Nickname    string
	Birthday    time.Time
	Description string
	Password    string
	Ctime       int64
	Utime       int64
}
