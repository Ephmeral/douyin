package db

import (
	"github.com/Ephmeral/douyin/pkg/constants"
	"gorm.io/gorm"
)

// MessageRaw Message Gorm Data Structures
type MessageRaw struct {
	gorm.Model
	Content    string `gorm:"column:contents;type:varchar(255);not null"`
	ToUserId   int64  `gorm:"column:from_user_id;not null;index:idx_touserid"`
	FromUserId int64  `gorm:"column:from_user_id;not null;index:idx_fromuserid"`
	CreateTime int64  `gorm:"column:create_time;not null;index:idx_create"`
}

func (MessageRaw) TableName() string {
	return constants.RelationServiceName
}
