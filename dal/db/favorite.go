package db

import (
	"github.com/Ephmeral/douyin/pkg/constants"
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
