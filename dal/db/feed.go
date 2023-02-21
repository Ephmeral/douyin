package db

import (
	"github.com/Ephmeral/douyin/pkg/constants"
	"gorm.io/gorm"
	"time"
)

// VideoRaw Video Gorm Data Structures
type VideoRaw struct {
	gorm.Model
	UserId    int64     `gorm:"column:user_id;not null;index:idx_userid"`
	Title     string    `gorm:"column:title;type:varchar(128);not null"`
	PlayUrl   string    `gorm:"column:play_url;varchar(128);not null"`
	CoverUrl  string    `gorm:"column:cover_url;varchar(128);not null"`
	UpdatedAt time.Time `gorm:"column:update_time;not null;index:idx_update"`
}

func (v *VideoRaw) TableName() string {
	return constants.VideoTableName
}
