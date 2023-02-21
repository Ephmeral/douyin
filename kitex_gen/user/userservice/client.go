// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "github.com/ephmeral/douyin/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CheckUser(ctx context.Context, Req *user.CheckUserRequest, callOptions ...callopt.Option) (r *user.CheckUserResponse, err error)
	RegisterUser(ctx context.Context, Req *user.RegisterUserRequest, callOptions ...callopt.Option) (r *user.RegisterUserResponse, err error)
	UserInfo(ctx context.Context, Req *user.UserInfoRequest, callOptions ...callopt.Option) (r *user.UserInfoResponse, err error)
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
	return &kUserServiceClient{
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

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) CheckUser(ctx context.Context, Req *user.CheckUserRequest, callOptions ...callopt.Option) (r *user.CheckUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckUser(ctx, Req)
}

func (p *kUserServiceClient) RegisterUser(ctx context.Context, Req *user.RegisterUserRequest, callOptions ...callopt.Option) (r *user.RegisterUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RegisterUser(ctx, Req)
}

func (p *kUserServiceClient) UserInfo(ctx context.Context, Req *user.UserInfoRequest, callOptions ...callopt.Option) (r *user.UserInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserInfo(ctx, Req)
}
