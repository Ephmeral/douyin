package oss

import (
	"github.com/Ephmeral/douyin/pkg/constants"
	"os"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var Bucket *oss.Bucket
var Path string

func Init() {
	// 算出绝对路径，防止service层测试时路径错误
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	Path = strings.Split(dir, "/cmd")[0]
	// 打开oss的Bucket
	endpoint := constants.OssEndPoint
	accessKeyId := constants.OssAccessKeyId
	accessKeySecret := constants.OssAccessKeySecret
	bucket := constants.OssBucket
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	Bucket, err = client.Bucket(bucket)
	if err != nil {
		panic(err)
	}

	// 上传默认头像
	err = PublishAvatarInit()
	if err != nil {
		panic(err)
	}
	// 上传默认背景
	err = PublishBackgroundInit()
	if err != nil {
		panic(err)
	}
}
