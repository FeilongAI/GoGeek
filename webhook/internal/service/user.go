package service

import (
	"context"
	"errors"
	"github.com/FeilongAI/GoGeek/baisic-go/webhook/internal/domain"
	"github.com/FeilongAI/GoGeek/baisic-go/webhook/internal/repostitory"
	"github.com/FeilongAI/GoGeek/baisic-go/webhook/internal/repostitory/dao"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserDuplicate = dao.ErrUserDuplicate
var ErrInvalidUserOrPassword = errors.New("账号或邮箱错误")

type UserService struct {
	repo *repostitory.UserRepository
}

func NewUserService(repo *repostitory.UserRepository) *UserService {
	return &UserService{repo: repo}

}
func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(ctx, u)
}
func (svc *UserService) Edit(ctx context.Context, u domain.EditUserInfo) error {
	return svc.repo.UpdateUserInfo(ctx, u)
}

func (svc *UserService) Login(c *gin.Context, req domain.User) (domain.User, error) {
	user, err := svc.repo.FindByEmail(c, req.Email)
	if err == repostitory.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {

		return domain.User{}, ErrInvalidUserOrPassword
	}
	return user, nil
}
