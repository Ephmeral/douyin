package handlers

import (
	"context"
	"github.com/Ephmeral/douyin/kitex_gen/message"
	"strconv"

	"github.com/Ephmeral/douyin/cmd/api/rpc"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/Ephmeral/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

// MessageAction implement message actions
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserIdStr := c.Query("to_user_id")
	actionTypeStr := c.Query("action_type")
	content := c.Query("content")

	if len(token) == 0 {
		SendResponse(c, errno.ParamErr)
		return
	}

	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamParseErr)
		return
	}

	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamParseErr)
		return
	}
	if actionType != constants.SendMessage {
		SendResponse(c, errno.ActionTypeErr)
		return
	}
	if len(content) == 0 {
		SendResponse(c, errno.ParamParseErr)
		return
	}

	req := &message.MessageActionRequest{Token: token, ToUserId: toUserId, ActionType: int32(actionType), Content: content}
	err = rpc.MessageAction(context.Background(), req)
	if err != nil {
		SendResponse(c, err)
		return
	}
	SendResponse(c, errno.Success)
}

// MessageChat get user message info
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserIdStr := c.Query("to_user_id")
	preMsgTime := c.Query("pre_msg_time")

	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamParseErr)
		return
	}
	msgTime, err := strconv.ParseInt(preMsgTime, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamParseErr)
		return
	}

	req := &message.MessageChatRequest{Token: token, ToUserId: toUserId, PreMsgTime: msgTime}
	msgList, err := rpc.MessageChat(context.Background(), req)
	if err != nil {
		SendResponse(c, err)
		return
	}
	SendMessageListResponse(c, errno.Success, msgList)
}
