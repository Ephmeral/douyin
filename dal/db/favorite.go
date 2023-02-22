package db

import (
	"context"
	"errors"
	"github.com/Ephmeral/douyin/dal/cache"
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

// CreateFavorite add a record to the favorite table through a transaction, and add the number of video likes
func CreateFavorite(ctx context.Context, favorite *FavoriteRaw, videoId int64) error {
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 修改video表之前先查询点赞状态
		state := cache.NewProxyIndexMap().GetFavorState(favorite.UserId, videoId)
		if state {
			// 已经点过赞了，不能重复点赞
			return errors.New("you have already liked it")
		}
		err := tx.Table("video").Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
		if err != nil {
			klog.Error("AddFavoriteCount error " + err.Error())
			return err
		}

		cache.NewProxyIndexMap().UpdateFavorState(favorite.UserId, favorite.VideoId, true)
		return nil
	})
	return nil
}

// DeleteFavorite Delete a record in the favorite table and reduce the number of video likes
func DeleteFavorite(ctx context.Context, currentId int64, videoId int64) error {
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 修改video表之前先查询点赞状态
		state := cache.NewProxyIndexMap().GetFavorState(currentId, videoId)
		if !state {
			// 不能对没有点赞的视频取消点赞
			return errors.New("you can't unlike a video that doesn't have a like")
		}
		err := tx.Table("video").Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
		if err != nil {
			klog.Error("SubFavoriteCount error " + err.Error())
			return err
		}

		cache.NewProxyIndexMap().UpdateFavorState(currentId, videoId, false)
		return nil
	})
	return nil
}
