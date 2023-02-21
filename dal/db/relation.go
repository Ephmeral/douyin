package db

import (
	"github.com/Ephmeral/douyin/pkg/constants"
	"gorm.io/gorm"
)

// RelationRaw Relation Gorm Data Structures
type RelationRaw struct {
	gorm.Model
	UserId   int64 `gorm:"column:user_id;not null;index:idx_userid"`
	ToUserId int64 `gorm:"column:to_user_id;not null;index:idx_touserid"`
}

func (RelationRaw) TableName() string {
	return constants.ReltaionTableName
}
