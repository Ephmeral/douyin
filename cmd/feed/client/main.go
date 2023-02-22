package client

import (
	"context"
	"fmt"
	"github.com/Ephmeral/douyin/kitex_gen/feed"
	"github.com/Ephmeral/douyin/kitex_gen/feed/feedservice"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/Ephmeral/douyin/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

func main() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := feedservice.NewClient(
		constants.FeedServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	response, err := c.Feed(context.Background(), &feed.FeedRequest{
		LatestTime: 100,
		Token:      "xcs",
	})
	if err != nil {
		fmt.Printf("message action error:%v", err)
		panic(err)
	}
	fmt.Println(response)
}
