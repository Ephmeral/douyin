package main

import (
	"context"
	message "github.com/Ephmeral/douyin/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {
	// TODO: Your code here...

	return &message.MessageActionResponse{
		BaseResp: &message.BaseResp{
			StatusCode:    100,
			StatusMessage: req.Content,
			ServiceTime:   100,
		},
	}, nil
}

// MessageList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageList(ctx context.Context, req *message.MessageListRequest) (resp *message.MessageListResponse, err error) {
	// TODO: Your code here...
	return
}
