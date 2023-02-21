package db

import (
	"github.com/Ephmeral/douyin/pkg/constants"
	"gorm.io/gorm"
)

// CommentRaw Comment Gorm Data Structures
type CommentRaw struct {
	gorm.Model
	UserId  int64  `gorm:"column:user_id;not null;index:idx_userid"`
	VideoId int64  `gorm:"column:video_id;not null;index:idx_videoid"`
	Content string `gorm:"column:content;type:varchar(255);not null"`
}

func (CommentRaw) TableName() string {
	return constants.CommentTableName
}
