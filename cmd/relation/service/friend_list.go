package service

import (
	"context"
	"errors"
	"github.com/chenmengangzhi29/douyin/dal/db"
	"github.com/chenmengangzhi29/douyin/dal/pack"
	"github.com/chenmengangzhi29/douyin/kitex_gen/relation"
	"github.com/chenmengangzhi29/douyin/pkg/constants"
	"github.com/chenmengangzhi29/douyin/pkg/jwt"
)

type FriendListService struct {
	ctx context.Context
}

func NewFriendListService(ctx context.Context) *FriendListService {
	return &FriendListService{ctx: ctx}
}

func (s *FriendListService) FriendList(req *relation.FriendListRequest) ([]*relation.FriendUser, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, _ := Jwt.CheckToken(req.Token)

	// 检查请求的用户是否存在
	user, err := db.QueryUserByIds(s.ctx, []int64{req.UserId})
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, errors.New("userId not exist")
	}

	// 查询目标用户的好友记录
	userIds, err := db.QueryFriendById(s.ctx, currentId, req.UserId)
	if err != nil {
		return nil, err
	}

	// 获取好友id

	// 获取好友的信息
	users, err := db.QueryUserByIds(s.ctx, userIds)
	if err != nil {
		return nil, err
	}

	// 查询当前用户与好友之间的最新消息记录
	var messageMap = make(map[int64]*db.MessageRaw)
	for _, user := range users {
		msg, err := db.QueryFriendLastMessage(s.ctx, currentId, int64(user.ID))
		if err != nil {
			continue
		}
		messageMap[int64(user.ID)] = msg
	}

	userList := pack.FriendList(currentId, users, messageMap)
	return userList, nil
}
