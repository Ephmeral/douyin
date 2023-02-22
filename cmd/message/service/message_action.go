package service

import (
	"context"
	"errors"
	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/kitex_gen/message"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/Ephmeral/douyin/pkg/jwt"
)

type MessageActionService struct {
	ctx context.Context
}

// NewMessageActionService new MessageActionService
func NewMessageActionService(ctx context.Context) *MessageActionService {
	return &MessageActionService{ctx: ctx}
}

// MessageAction implement send action
// 如果actionType等于1，表示当前用户发送消息给目标用户，
// 新建一条消息记录
func (s *MessageActionService) MessageAction(req *message.MessageActionRequest) error {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, err := Jwt.CheckToken(req.Token)
	if err != nil {
		return err
	}

	users, err := db.QueryUserByIds(s.ctx, []int64{req.ToUserId})
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return errors.New("toUserId not exist")
	}

	if req.ActionType == constants.SendMessage {
		err := db.AddMessage(s.ctx, currentId, req.ToUserId, req.Content)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("ActionType Err")
}
