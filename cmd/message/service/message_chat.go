package service

import (
	"context"
	"errors"
	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/kitex_gen/message"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/Ephmeral/douyin/pkg/jwt"
)

type MessageChatService struct {
	ctx context.Context
}

// NewMessageChatService new MessageChatService
func NewMessageChatService(ctx context.Context) *MessageChatService {
	return &MessageChatService{ctx: ctx}
}

// MessageChat implement message chat
// 查询用户id和目标用户之间的消息记录
func (s *MessageChatService) MessageChat(req *message.MessageChatRequest) ([]*message.Message, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, err := Jwt.CheckToken(req.Token)
	if err != nil {
		return nil, err
	}

	users, err := db.QueryUserByIds(s.ctx, []int64{req.ToUserId})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("toUserId not exist")
	}

	var result []*message.Message
	//获取目标用户关注的用户id号
	messages, err := db.QueryMessageById(s.ctx, currentId, req.ToUserId)
	if err != nil {
		return nil, err
	}
	for _, msg := range messages {
		result = append(result, &message.Message{Id: int64(msg.ID),
			ToUserId:   msg.ToUserId,
			FromUserId: msg.UserId,
			Content:    msg.Content,
			CreateTime: msg.CreateTime})
	}

	messages, err = db.QueryMessageById(s.ctx, req.ToUserId, currentId)
	if err != nil {
		return nil, err
	}
	for _, msg := range messages {
		result = append(result, &message.Message{Id: int64(msg.ID),
			ToUserId:   msg.ToUserId,
			FromUserId: msg.UserId,
			Content:    msg.Content,
			CreateTime: msg.CreateTime})
	}

	return result, nil
}
