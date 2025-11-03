package service

import (
	"context"
	"github.com/FeilongAI/GoGeek/baisic-go/webhook/internal/domain"
	"github.com/FeilongAI/GoGeek/baisic-go/webhook/internal/repostitory"
)

type UserService struct {
	repo *repostitory.UserRepository
}

func NewUserService(repo *repostitory.UserRepository) *UserService {
	return &UserService{repo: repo}

}
func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	return svc.repo.Create(ctx, u)
}
func (svc *UserService) Edit(ctx context.Context, u domain.EditUserInfo) error {
	return svc.repo.UpdateUserInfo(ctx, u)
}
