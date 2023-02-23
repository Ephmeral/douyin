package rpc

import (
	"context"
	"github.com/Ephmeral/douyin/kitex_gen/message"
	"github.com/Ephmeral/douyin/kitex_gen/message/messageservice"
	"time"

	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/Ephmeral/douyin/pkg/errno"
	"github.com/Ephmeral/douyin/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var messageClient messageservice.Client

func initMessageRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := messageservice.NewClient(
		constants.MessageServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(500*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	messageClient = c
}

// MessageAction implement send message actions
func MessageAction(ctx context.Context, req *message.MessageActionRequest) error {
	resp, err := messageClient.MessageAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// MessageChat implement send and receive message
func MessageChat(ctx context.Context, req *message.MessageChatRequest) ([]*message.Message, error) {
	resp, err := messageClient.MessageChat(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.MessageList, nil
}
