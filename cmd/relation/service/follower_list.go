package service

import (
	"context"
	"errors"

	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/dal/pack"
	"github.com/Ephmeral/douyin/kitex_gen/relation"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/Ephmeral/douyin/pkg/jwt"
)

type FollowerListService struct {
	ctx context.Context
}

// NewFollowerListService new FollowerListService
func NewFollowerListService(ctx context.Context) *FollowerListService {
	return &FollowerListService{ctx: ctx}
}

// FollowerList 查询用户粉丝列表
func (s *FollowerListService) FollowerList(req *relation.FollowerListRequest) ([]*relation.User, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, err := Jwt.CheckToken(req.Token)
	if err != nil {
		return nil, err
	}
	// 检查请求的用户是否存在
	user, err := db.QueryUserByIds(s.ctx, []int64{req.UserId})
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, errors.New("userId not exist")
	}

	//查询目标用户的被关注记录
	relations, err := db.QueryFollowerById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	//获取这些记录的关注方id
	userIds := make([]int64, 0)
	for _, relation := range relations {
		userIds = append(userIds, relation.UserId)
	}

	//获取关注方的信息
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
