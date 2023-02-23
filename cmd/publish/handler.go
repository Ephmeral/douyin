package main

import (
	"context"
	"github.com/Ephmeral/douyin/cmd/publish/service"
	"github.com/Ephmeral/douyin/dal/pack"
	publish "github.com/Ephmeral/douyin/kitex_gen/publish"
	"github.com/Ephmeral/douyin/pkg/errno"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	resp = new(publish.PublishActionResponse)

	if len(req.Token) == 0 || len(req.Title) == 0 || req.Data == nil {
		resp.BaseResp = pack.BuildPublishBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewPublishService(ctx).Publish(req)
	if err != nil {
		resp.BaseResp = pack.BuildPublishBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildPublishBaseResp(errno.Success)
	return resp, nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	resp = new(publish.PublishListResponse)

	if req.UserId <= 0 {
		resp.BaseResp = pack.BuildPublishBaseResp(errno.ParamErr)
		return resp, nil
	}

	videoList, err := service.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		resp.BaseResp = pack.BuildPublishBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildPublishBaseResp(errno.Success)
	resp.VideoList = videoList
	return resp, nil
}
