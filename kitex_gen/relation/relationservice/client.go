// Code generated by Kitex v0.4.4. DO NOT EDIT.

package relationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	relation "github.com/ephmeral/douyin/kitex_gen/relation"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RelationAction(ctx context.Context, Req *relation.RelationActionRequest, callOptions ...callopt.Option) (r *relation.RelationActionResponse, err error)
	FollowList(ctx context.Context, Req *relation.FollowListRequest, callOptions ...callopt.Option) (r *relation.FollowListResponse, err error)
	FollowerList(ctx context.Context, Req *relation.FollowerListRequest, callOptions ...callopt.Option) (r *relation.FollowerListResponse, err error)
	FriendList(ctx context.Context, Req *relation.FriendListRequest, callOptions ...callopt.Option) (r *relation.FriendListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kRelationServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) RelationAction(ctx context.Context, Req *relation.RelationActionRequest, callOptions ...callopt.Option) (r *relation.RelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationAction(ctx, Req)
}

func (p *kRelationServiceClient) FollowList(ctx context.Context, Req *relation.FollowListRequest, callOptions ...callopt.Option) (r *relation.FollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowList(ctx, Req)
}

func (p *kRelationServiceClient) FollowerList(ctx context.Context, Req *relation.FollowerListRequest, callOptions ...callopt.Option) (r *relation.FollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowerList(ctx, Req)
}

func (p *kRelationServiceClient) FriendList(ctx context.Context, Req *relation.FriendListRequest, callOptions ...callopt.Option) (r *relation.FriendListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FriendList(ctx, Req)
}
