package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/Ephmeral/douyin/dal/db"
	User "github.com/Ephmeral/douyin/kitex_gen/user"
	"github.com/Ephmeral/douyin/pkg/errno"
)

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

// CheckUser check user info
func (s *CheckUserService) CheckUser(request *User.CheckUserRequest) (int64, error) {
	hash := md5.New()
	if _, err := hash.Write([]byte(request.Password)); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", hash.Sum(nil))

	userName := request.Username
	user, err := db.QueryUserByName(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if user.Password != passWord {
		return 0, errno.LoginErr
	}
	return int64(user.ID), nil
}
