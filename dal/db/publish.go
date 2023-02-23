package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

func PublishVideoData(ctx context.Context, videoData *VideoRaw) error {
	if err := DB.WithContext(ctx).Create(&videoData).Error; err != nil {
		klog.Error("PublishVideoData error " + err.Error())
		return err
	}
	return nil
}

func QueryVideoByUserId(ctx context.Context, userId int64) ([]*VideoRaw, error) {
	var videos []*VideoRaw
	err := DB.Table("video").WithContext(ctx).Order("update_time desc").Where("user_id = ?", userId).Find(&videos).Error
	if err != nil {
		klog.Error("QueryVideoByUserId find video error " + err.Error())
		return nil, err
	}
	return videos, nil
}
