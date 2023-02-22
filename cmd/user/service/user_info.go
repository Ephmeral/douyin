package service

import (
	"context"
	"errors"
	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/dal/pack"
	User "github.com/Ephmeral/douyin/kitex_gen/user"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/Ephmeral/douyin/pkg/jwt"
)

type UserInfoService struct {
	ctx context.Context
}

// NewUserInfoService new UserInfoService
func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{
		ctx: ctx,
	}
}

func (s *UserInfoService) UserInfo(request *User.UserInfoRequest) (*User.User, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, err := Jwt.CheckToken(request.Token)
	if err != nil {
		return nil, err
	}

	userIds := []int64{request.UserId}
	users, err := db.QueryUserByIds(s.ctx, userIds)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("user not exist")
	}
	user := users[0]

	relationMap, err := db.QueryRelationByIds(s.ctx, currentId, userIds)
	if err != nil {
		return nil, err
	}

	isFollow := false
	_, ok := relationMap[request.UserId]
	if ok {
		isFollow = true
	}

	follow, err := db.QueryFollowById(s.ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	follower, err := db.QueryFollowerById(s.ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	totalFavorited, err := db.QueryUserFavoritedById(s.ctx, request.UserId)
	if err != nil {
		return nil, err
	}
	video, err := db.QueryVideoByUserId(s.ctx, request.UserId)
	var favoriteCount int64
	for _, v := range video {
		count, err := db.QueryUserFavoritedById(s.ctx, int64(v.ID))
		if err != nil {
			return nil, err
		}
		favoriteCount += count
	}

	userInfo := pack.UserInfo(user, isFollow, int64(len(follow)), int64(len(follower)), totalFavorited, int64(len(video)), favoriteCount)
	return userInfo, nil
}
