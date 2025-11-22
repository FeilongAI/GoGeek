package dao

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicate = errors.New("user duplicated")
	ErrUserNotFound  = gorm.ErrRecordNotFound
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
	err := dao.db.Create(&u).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			//邮箱冲突
			return ErrUserDuplicate
		}
	}
	return err
}
func (dao *UserDao) Update(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Ctime = now
	return dao.db.Model(&User{}).Where("id = ?", u.Id).UpdateColumns(&u).Error
}

func (dao *UserDao) FindByEmail(c *gin.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(c).Where("email = ?", email).First(&u).Error
	return u, err
}

type User struct {
	Id          int64  `gorm:"primary_key,autoIncrement"`
	Email       string `gorm:"unique"`
	Nickname    string
	Birthday    time.Time `gorm:"DEFAULT:null"`
	Description string
	Password    string
	Ctime       int64
	Utime       int64
}
