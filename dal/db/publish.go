package db

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
)

// QueryVideoByUserId 新增通过用户id获取视频数据的功能
func QueryVideoByUserId(ctx context.Context, userId int64) ([]*VideoRaw, error) {
	var videos []*VideoRaw
	err := DB.Table("video").WithContext(ctx).Order("update_time desc").Where("user_id = ?", userId).Find(&videos).Error
	if err != nil {
		klog.Error("QueryVideoByUserId find video error " + err.Error())
		return nil, err
	}
	return videos, nil
}
