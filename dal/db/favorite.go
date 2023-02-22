package db

import (
	"context"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// FavoriteRaw Gorm Data Structures
type FavoriteRaw struct {
	gorm.Model
	UserId  int64 `gorm:"column:user_id;not null;index:idx_userid"`
	VideoId int64 `gorm:"column:video_id;not null;index:idx_videoid"`
}

func (FavoriteRaw) TableName() string {
	return constants.FavoriteTableName
}

func QueryUserFavoritedById(ctx context.Context, userId int64) (int64, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&FavoriteRaw{}).Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		klog.Error("query user favorited by id " + err.Error())
		return 0, err
	}
	return count, nil
}
