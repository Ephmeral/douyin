package db

import (
	"context"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/Ephmeral/douyin/pkg/oss"
	"github.com/cloudwego/kitex/pkg/klog"
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

// QueryUserByIds 根据用户id获取用户信息
func QueryUserByIds(ctx context.Context, userIds []int64) ([]*UserRaw, error) {
	var users []*UserRaw
	err := DB.WithContext(ctx).Where("id in (?)", userIds).Find(&users).Error
	if err != nil {
		klog.Error("query user by ids fail " + err.Error())
		return nil, err
	}
	return users, nil
}

func QueryUserByName(ctx context.Context, name string) (*UserRaw, error) {
	var user UserRaw
	err := DB.WithContext(ctx).Where("name = ?", name).First(&user).Error
	if err != nil {
		klog.Error("query user by name fail " + err.Error())
		return nil, err
	}
	return &user, nil
}

func CreateUserInfo(ctx context.Context, username string, password string) (int64, error) {
	user := &UserRaw{
		Name:            username,
		Password:        password,
		Avatar:          oss.GetAvatar(),
		BackgroundImage: "",
		Signature:       "",
	}
	err := DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		klog.Error("Create user info fail " + err.Error())
		return 0, err
	}
	return int64(user.ID), nil
}
