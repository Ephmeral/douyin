package main

import (
	"context"
	"github.com/Ephmeral/douyin/cmd/feed/service"
	"github.com/Ephmeral/douyin/dal/pack"
	feed "github.com/Ephmeral/douyin/kitex_gen/feed"
	"github.com/Ephmeral/douyin/pkg/errno"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// 创建一个新的响应结构体
	resp = new(feed.FeedResponse)

	if req.LatestTime <= 0 { // 如果LatestTime小于等于0，返回参数错误
		// 用参数错误的信息构建一个BaseResp结构体并存储到响应结构体中
		resp.BaseResp = pack.BuildFeedBaseResp(errno.ParamErr)
		return resp, nil
	}

	// 如果LatestTime大于0，则调用Feed服务中的服务实现处理Feed请求
	videos, nextTime, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		resp.BaseResp = pack.BuildFeedBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildFeedBaseResp(errno.Success)
	resp.VideoList = videos  // 将视频列表设置到响应结构体中
	resp.NextTime = nextTime // 将视频列表设置到响应结构体中
	return resp, nil
}
