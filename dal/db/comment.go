package db

import (
	"context"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
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

// CreateComment 通过一条评论创建一条评论记录并增加视频评论数
func CreateComment(ctx context.Context, comment *CommentRaw) error {
	err := DB.WithContext(ctx).Table(constants.CommentTableName).Create(&comment).Error
	if err != nil {
		klog.Error("create comment fail " + err.Error())
		return err
	}
	return nil
}

// DeleteComment 通过评论id号删除一条评论并减少视频评论数，返回该评论
func DeleteComment(ctx context.Context, commentId int64) (*CommentRaw, error) {
	var commentRaw *CommentRaw
	err := DB.WithContext(ctx).Where("id = ?", commentId).First(&commentRaw).Error
	if err == gorm.ErrRecordNotFound {
		klog.Errorf("not find comment %v, %v", commentRaw, err.Error())
		return nil, err
	}
	if err != nil {
		klog.Errorf("find comment %v fail, %v", commentRaw, err.Error())
		return nil, err
	}
	err = DB.WithContext(ctx).Where("id = ?", commentId).Delete(&CommentRaw{}).Error
	if err != nil {
		klog.Error("delete comment fail " + err.Error())
		return nil, err
	}

	return commentRaw, nil
}

// QueryCommentByCommentIds 通过评论id查询一组评论信息
func QueryCommentByCommentIds(ctx context.Context, commentIds []int64) ([]*CommentRaw, error) {
	var comments []*CommentRaw
	err := DB.WithContext(ctx).Table("comment").Where("id in (?)", commentIds).Find(&comments).Error
	if err != nil {
		klog.Error("query comment by comment id fail " + err.Error())
		return nil, err
	}
	return comments, nil
}

// QueryCommentByVideoId 通过视频id号倒序返回一组评论信息
func QueryCommentByVideoId(ctx context.Context, videoId int64) ([]*CommentRaw, error) {
	var comments []*CommentRaw
	err := DB.WithContext(ctx).Table("comment").Order("updated_at desc").Where("video_id = ?", videoId).Find(&comments).Error
	if err != nil {
		klog.Error("query comment by video id fail " + err.Error())
		return nil, err
	}
	return comments, nil
}
