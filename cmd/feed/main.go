package main

import (
	"github.com/Ephmeral/douyin/dal"
	feed "github.com/Ephmeral/douyin/kitex_gen/feed/feedservice"
	"github.com/Ephmeral/douyin/pkg/bound"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/Ephmeral/douyin/pkg/middleware"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

func Init() {
	dal.Init()
}
func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress}) //r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.FeedAddress)
	if err != nil {
		panic(err)
	}
	Init()
	svr := feed.NewServer(new(FeedServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.FeedServiceName}), //server name
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       //address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), //limit
		server.WithMuxTransport(),                                          //Multiplex
		server.WithBoundHandler(bound.NewCpuLimitHandler()),                //BoundHandler
		server.WithRegistry(r),                                             //registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
