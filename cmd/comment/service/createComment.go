package service

import (
	"context"
	"errors"
	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/dal/pack"
	"github.com/Ephmeral/douyin/kitex_gen/comment"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/Ephmeral/douyin/pkg/jwt"
	"sync"
)

type CreateCommentService struct {
	ctx context.Context
}

// NewCreateCommentService new CreateCommentService
func NewCreateCommentService(ctx context.Context) *CreateCommentService {
	return &CreateCommentService{ctx: ctx}
}

// CreateComment add comment
func (s *CreateCommentService) CreateComment(req *comment.CreateCommentRequest) (*comment.Comment, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, err := Jwt.CheckToken(req.Token)
	if err != nil {
		return nil, err
	}

	videos, err := db.QueryVideoByVideoIds(s.ctx, []int64{req.VideoId})
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 {
		return nil, errors.New("video not exist")
	}

	commentRaw := &db.CommentRaw{
		UserId:  currentId,
		VideoId: req.VideoId,
		Content: req.CommentText,
	}

	var wg sync.WaitGroup
	wg.Add(2)
	var user *db.UserRaw
	var commentErr, userErr error
	//创建评论记录并增加视频评论数
	go func() {
		defer wg.Done()
		err := db.CreateComment(s.ctx, commentRaw)
		if err != nil {
			commentErr = err
			return
		}
	}()
	//获取当前用户信息
	go func() {
		defer wg.Done()
		users, err := db.QueryUserByIds(s.ctx, []int64{currentId})
		if err != nil {
			userErr = err
			return
		}
		user = users[0]
	}()
	wg.Wait()
	if commentErr != nil {
		return nil, commentErr
	}
	if userErr != nil {
		return nil, userErr
	}

	comment := pack.CommentInfo(commentRaw, user)
	return comment, nil

}
