package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/kitex_gen/user"
	"github.com/Ephmeral/douyin/pkg/errno"
)

type RegisterUserService struct {
	ctx context.Context
}

func NewRegisterUserService(ctx context.Context) *RegisterUserService {
	return &RegisterUserService{
		ctx: ctx,
	}
}

// RegisterUser register user info
func (s *RegisterUserService) RegisterUser(request *user.RegisterUserRequest) (int64, error) {
	_, err := db.QueryUserByName(s.ctx, request.Username)
	if err == nil {
		return 0, errno.UserAlreadyExistErr
	}

	hash := md5.New()
	if _, err := hash.Write([]byte(request.Password)); err != nil {
		return 0, err
	}
	password := fmt.Sprintf("%x", hash.Sum(nil))

	userId, err := db.CreateUserInfo(s.ctx, request.Username, password)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
