package db

import (
	"context"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// MessageRaw Message Gorm Data Structures
type MessageRaw struct {
	gorm.Model
	UserId     int64  `gorm:"column:user_id;not null;index:idx_userid"`
	ToUserId   int64  `gorm:"column:to_user_id;not null;index:idx_touserid"`
	Content    string `gorm:"column:content;not null;index:idx_content"`
	CreateTime int64  `gorm:"column:create_time;not null;index:idx_createtime;autoCreateTime"` // 使用时间戳秒数填充创建时间
}

func (MessageRaw) TableName() string {
	return constants.MessageTableName
}

// AddMessage 添加一条消息记录
func AddMessage(ctx context.Context, currentId int64, toUserId int64, content string) error {
	messageRaw := &MessageRaw{
		UserId:   currentId,
		ToUserId: toUserId,
		Content:  content,
	}
	err := DB.WithContext(ctx).Table("message").Create(&messageRaw).Error
	if err != nil {
		klog.Error("add message record fail " + err.Error())
		return err
	}
	return nil
}

// QueryMessageById 根据用户id和目标用户id，查询对应的消息记录
func QueryMessageById(ctx context.Context, userId int64, toUserId int64) ([]*MessageRaw, error) {
	var messages []*MessageRaw
	err := DB.WithContext(ctx).Table(constants.MessageTableName).
		Where("user_id = ? and to_user_id = ?", userId, toUserId).Find(&messages).Error
	if err != nil {
		klog.Error("query message by id fail " + err.Error())
		return nil, err
	}
	return messages, nil
}

// QueryFriendLastMessage 根据用户id和目标用户id，查询最新的消息记录
func QueryFriendLastMessage(ctx context.Context, userId int64, toUserId int64) (*MessageRaw, error) {
	var msg *MessageRaw
	err := DB.WithContext(ctx).Table(constants.MessageTableName).
		Where("user_id = ? and to_user_id = ? ", userId, toUserId).
		Or("user_id = ? and to_user_id = ? ", toUserId, userId).
		Order("create_time DESC").Find(&msg).Error
	if err != nil {
		klog.Error("query friend last message by id fail " + err.Error())
		return nil, err
	}
	return msg, nil
}
