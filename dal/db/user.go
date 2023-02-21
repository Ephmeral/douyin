package db

import (
	"github.com/Ephmeral/douyin/pkg/constants"
	"gorm.io/gorm"
)

// UserRaw User Gorm Data structures
type UserRaw struct {
	gorm.Model
	Name            string `gorm:"column:name;index:idx_username,unique;type:varchar(32);not null"`
	Password        string `gorm:"column:password;type:varchar(32);not null"`
	Avatar          string `gorm:"column:avatar;type:varchar(128);not null"`
	BackgroundImage string `gorm:"column:backGround;type:varchar(128);not null"`
	Signature       string `gorm:"column:signature;type:varchar(128);not null"`
}

func (UserRaw) TableName() string {
	return constants.UserTableName
}
