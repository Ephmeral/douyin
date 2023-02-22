package service

import (
	"context"
	"errors"

	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/kitex_gen/relation"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/Ephmeral/douyin/pkg/jwt"
)

type RelationActionService struct {
	ctx context.Context
}

// NewRelationActionService new RelationActionService
func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

// RelationAction 实现关注和取消关注操作
// 如果actionType等于1，表示当前用户关注其他用户，新建一条关注记录
// 如果actionType等于2，表示当前用户取消关注其他用户，删除该关注记录
func (s *RelationActionService) RelationAction(req *relation.RelationActionRequest) error {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, err := Jwt.CheckToken(req.Token)
	if err != nil {
		return err
	}

	// 检查toUserId是否存在
	users, err := db.QueryUserByIds(s.ctx, []int64{req.ToUserId})
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return errors.New("toUserId not exist")
	}

	// 不能自己关注自己
	if currentId == req.ToUserId {
		return errors.New("can not follow yourself")
	}

	// 当前用户关注其他用户
	if req.ActionType == constants.Follow {
		err := db.AddRelation(s.ctx, currentId, req.ToUserId)
		if err != nil {
			return err
		}
		return nil
	}
	// 当前用户取消关注其他用户
	if req.ActionType == constants.UnFollow {
		err := db.DeleteRelation(s.ctx, currentId, req.ToUserId)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("ActionType Err")
}
