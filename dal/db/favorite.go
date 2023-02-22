package db

import (
	"context"
<<<<<<< HEAD
=======
	"errors"
	"github.com/Ephmeral/douyin/dal/cache"
>>>>>>> 17f73c6b966bdb94c10b445e5c0232fb8972bb5d
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

<<<<<<< HEAD
//根据当前用户id和视频id获取点赞信息
func QueryFavoriteByIds(ctx context.Context, currentId int64, videoIds []int64) (map[int64]*FavoriteRaw, error) {
	var favorites []*FavoriteRaw
	err := DB.WithContext(ctx).Where("user_id = ? AND video_id IN ?", currentId, videoIds).Find(&favorites).Error
	if err != nil {
		klog.Error("quert favorite record fail " + err.Error())
		return nil, err
	}
	favoriteMap := make(map[int64]*FavoriteRaw)
	for _, favorite := range favorites {
		favoriteMap[favorite.VideoId] = favorite
	}
	return favoriteMap, nil
}

// CreateFavorite add a record to the favorite table through a transaction, and add the number of video likes
func CreateFavorite(ctx context.Context, favorite *FavoriteRaw, videoId int64) error {
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
=======
// CreateFavorite add a record to the favorite table through a transaction, and add the number of video likes
func CreateFavorite(ctx context.Context, favorite *FavoriteRaw, videoId int64) error {
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 修改video表之前先查询点赞状态
		state := cache.NewProxyIndexMap().GetFavorState(favorite.UserId, videoId)
		if state {
			// 已经点过赞了，不能重复点赞
			return errors.New("you have already liked it")
		}
>>>>>>> 17f73c6b966bdb94c10b445e5c0232fb8972bb5d
		err := tx.Table("video").Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
		if err != nil {
			klog.Error("AddFavoriteCount error " + err.Error())
			return err
		}

<<<<<<< HEAD
		err = tx.Table("favorite").Create(favorite).Error
		if err != nil {
			klog.Error("create favorite record fail " + err.Error())
			return err
		}

=======
		cache.NewProxyIndexMap().UpdateFavorState(favorite.UserId, favorite.VideoId, true)
>>>>>>> 17f73c6b966bdb94c10b445e5c0232fb8972bb5d
		return nil
	})
	return nil
}

<<<<<<< HEAD
//DeleteFavorite Delete a record in the favorite table and reduce the number of video likes
func DeleteFavorite(ctx context.Context, currentId int64, videoId int64) error {
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var favorite *FavoriteRaw
		err := tx.Table("favorite").Where("user_id = ? AND video_id = ?", currentId, videoId).Delete(&favorite).Error
		if err != nil {
			klog.Error("delete favorite record fail " + err.Error())
			return err
		}

		err = tx.Table("video").Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
=======
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
>>>>>>> 17f73c6b966bdb94c10b445e5c0232fb8972bb5d
		if err != nil {
			klog.Error("SubFavoriteCount error " + err.Error())
			return err
		}
<<<<<<< HEAD
=======

		cache.NewProxyIndexMap().UpdateFavorState(currentId, videoId, false)
>>>>>>> 17f73c6b966bdb94c10b445e5c0232fb8972bb5d
		return nil
	})
	return nil
}
<<<<<<< HEAD

// QueryFavoriteById 通过一个用户id查询出该用户点赞的所有视频id号
func QueryFavoriteById(ctx context.Context, userId int64) ([]int64, error) {
	var favorites []*FavoriteRaw
	err := DB.WithContext(ctx).Table("favorite").Where("user_id = ?", userId).Find(&favorites).Error
	if err != nil {
		klog.Error("query favorite record fail " + err.Error())
		return nil, err
	}
	videoIds := make([]int64, 0)
	for _, favorite := range favorites {
		videoIds = append(videoIds, favorite.VideoId)
	}
	return videoIds, nil
}
=======
>>>>>>> 17f73c6b966bdb94c10b445e5c0232fb8972bb5d
