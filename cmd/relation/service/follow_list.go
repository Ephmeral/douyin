package service

import (
	"context"
	"errors"
	"github.com/Ephmeral/douyin/dal/pack"

	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/kitex_gen/relation"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/Ephmeral/douyin/pkg/jwt"
)

type FollowListService struct {
	ctx context.Context
}

// NewFollowListService new FollowListService
func NewFollowListService(ctx context.Context) *FollowListService {
	return &FollowListService{ctx: ctx}
}

// FollowList get user follow list info
func (s *FollowListService) FollowList(req *relation.FollowListRequest) ([]*relation.User, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, _ := Jwt.CheckToken(req.Token)

	user, err := db.QueryUserByIds(s.ctx, []int64{req.UserId})
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, errors.New("userId not exist")
	}

	// 获取目标用户关注的用户id号
	relations, err := db.QueryFollowById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	userIds := make([]int64, 0)
	for _, relation := range relations {
		userIds = append(userIds, relation.ToUserId)
	}

	// 获取用户id号对应的用户信息
	users, err := db.QueryUserByIds(s.ctx, userIds)
	if err != nil {
		return nil, err
	}

	relationMap, err := db.QueryRelationByIds(s.ctx, currentId, userIds)
	if err != nil {
		return nil, err
	}
	userList := pack.UserList(currentId, users, relationMap)
	return userList, nil
}
