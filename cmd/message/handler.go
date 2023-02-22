package main

import (
	"context"
	"github.com/Ephmeral/douyin/cmd/message/service"
	"github.com/Ephmeral/douyin/dal/pack"
	message "github.com/Ephmeral/douyin/kitex_gen/message"
	"github.com/Ephmeral/douyin/pkg/errno"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {
	resp = new(message.MessageActionResponse)

	if len(req.Token) == 0 || req.ToUserId == 0 || req.ActionType == 0 {
		resp.BaseResp = pack.BuildMessageBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewMessageActionService(ctx).MessageAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildMessageBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildMessageBaseResp(errno.Success)
	return resp, nil
}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	resp = new(message.MessageChatResponse)

	if len(req.Token) == 0 || req.ToUserId == 0 {
		resp.BaseResp = pack.BuildMessageBaseResp(errno.ParamErr)
		return resp, nil
	}

	messages, err := service.NewMessageChatService(ctx).MessageChat(req)
	if err != nil {
		resp.BaseResp = pack.BuildMessageBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildMessageBaseResp(errno.Success)
	resp.MessageList = messages
	return resp, nil
}
