package repostitory

import (
	"context"
	"github.com/FeilongAI/GoGeek/baisic-go/webhook/internal/domain"
	"github.com/FeilongAI/GoGeek/baisic-go/webhook/internal/repostitory/dao"
)

type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{dao: dao}
}
func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}
func (r *UserRepository) UpdateUserInfo(ctx context.Context, u domain.EditUserInfo) error {
	return r.dao.Update(ctx, dao.User{
		Nickname:    u.Nickname,
		Birthday:    u.Birthday,
		Description: u.Description,
	})

}
