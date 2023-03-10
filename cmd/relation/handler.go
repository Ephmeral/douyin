package main

import (
	"context"
	"github.com/Ephmeral/douyin/cmd/relation/service"
	"github.com/Ephmeral/douyin/dal/pack"
	relation "github.com/Ephmeral/douyin/kitex_gen/relation"
	"github.com/Ephmeral/douyin/pkg/errno"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	resp = new(relation.RelationActionResponse)

	if len(req.Token) == 0 || req.ToUserId == 0 || req.ActionType == 0 {
		resp.BaseResp = pack.BuildRelationBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildRelationBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildRelationBaseResp(errno.Success)
	return resp, nil
}

// FollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowList(ctx context.Context, req *relation.FollowListRequest) (resp *relation.FollowListResponse, err error) {
	resp = new(relation.FollowListResponse)

	if req.UserId == 0 {
		resp.BaseResp = pack.BuildRelationBaseResp(errno.ParamErr)
		return resp, nil
	}

	userList, err := service.NewFollowListService(ctx).FollowList(req)
	if err != nil {
		resp.BaseResp = pack.BuildRelationBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildRelationBaseResp(errno.Success)
	resp.UserList = userList
	return resp, nil
}

// FollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowerList(ctx context.Context, req *relation.FollowerListRequest) (resp *relation.FollowerListResponse, err error) {
	resp = new(relation.FollowerListResponse)

	if req.UserId == 0 {
		resp.BaseResp = pack.BuildRelationBaseResp(errno.ParamErr)
		return resp, nil
	}

	userList, err := service.NewFollowerListService(ctx).FollowerList(req)
	if err != nil {
		resp.BaseResp = pack.BuildRelationBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildRelationBaseResp(errno.Success)
	resp.UserList = userList
	return resp, nil
}

// FriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FriendList(ctx context.Context, req *relation.FriendListRequest) (resp *relation.FriendListResponse, err error) {
	resp = new(relation.FriendListResponse)

	if req.UserId == 0 {
		resp.BaseResp = pack.BuildRelationBaseResp(errno.ParamErr)
		return resp, nil
	}

	userList, err := service.NewFriendListService(ctx).FriendList(req)
	if err != nil {
		resp.BaseResp = pack.BuildRelationBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildRelationBaseResp(errno.Success)
	resp.UserList = userList
	return resp, nil
}
