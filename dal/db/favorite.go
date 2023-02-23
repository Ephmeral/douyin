package db

import (
	"context"

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

// 根据当前用户id和视频id获取点赞信息
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
	cache.NewProxyIndexMap().UpdateFavorState(favorite.UserId, favorite.VideoId, true)
	return nil
}

// DeleteFavorite Delete a record in the favorite table and reduce the number of video likes
func DeleteFavorite(ctx context.Context, currentId int64, videoId int64) error {
	cache.NewProxyIndexMap().UpdateFavorState(currentId, videoId, false)
	return nil
}

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

func QueryUserFavoritedById(ctx context.Context, userId int64) (int64, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&FavoriteRaw{}).Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		klog.Error("query user favorite by id " + err.Error())
		return 0, err
	}
	return count, nil
}
