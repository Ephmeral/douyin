package db

import (
	"context"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
)

// PublishVideoData 向video表添加一条记录
func PublishVideoData(ctx context.Context, videoData *VideoRaw) error {
	if err := DB.WithContext(ctx).Create(&videoData).Error; err != nil {
		klog.Error("PublishVideoData error " + err.Error())
		return err
	}
	return nil
}

// QueryVideoByUserId 新增通过用户id获取视频数据的功能
func QueryVideoByUserId(ctx context.Context, userId int64) ([]*VideoRaw, error) {
	var videos []*VideoRaw
	err := DB.Table(constants.VideoTableName).WithContext(ctx).Order("update_time desc").Where("user_id = ?", userId).Find(&videos).Error
	if err != nil {
		klog.Error("QueryVideoByUserId find video error " + err.Error())
		return nil, err
	}
	return videos, nil
}
