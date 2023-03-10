package rpc

import (
	"context"
	"time"

	"github.com/Ephmeral/douyin/kitex_gen/favorite"
	"github.com/Ephmeral/douyin/kitex_gen/favorite/favoriteservice"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/Ephmeral/douyin/pkg/errno"
	"github.com/Ephmeral/douyin/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var favoriteClient favoriteservice.Client

func initFavoriteRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := favoriteservice.NewClient(
		constants.FavoriteServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

// FavoriteAction implement like and unlike operations
func FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) error {
	resp, err := favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// FavoriteList get favorite list info
func FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) ([]*favorite.Video, error) {
	resp, err := favoriteClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.VideoList, nil
}
