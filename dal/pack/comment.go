package pack

import (
	"context"
	"github.com/Ephmeral/douyin/dal/cache"
	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/kitex_gen/comment"
	"github.com/Ephmeral/douyin/pkg/constants"
)

func CommentInfo(commentRaw *db.CommentRaw, user *db.UserRaw) *comment.Comment {
	userFavorCount, err := cache.NewProxyIndexMap().GetFavorCount(int64(user.ID))
	if err != nil {
		userFavorCount = 0
	}
	return &comment.Comment{
		Id: int64(commentRaw.ID),
		User: &comment.User{
			Id:              int64(user.ID),
			Name:            user.Name,
			FollowCount:     db.QueryFollowCount(int64(user.ID)),
			FollowerCount:   db.QueryFollowerCount(int64(user.ID)),
			IsFollow:        false,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  db.QueryUserTotalFavorited(context.Background(), int64(user.ID)),
			WorkCount:       db.QueryVideoCountByUserId(int64(user.ID)),
			FavoriteCount:   userFavorCount,
		},
		Content:    commentRaw.Content,
		CreateDate: commentRaw.CreatedAt.Format(constants.TimeFormat),
		LikeCount:  0,
		TeaseCount: 0,
	}
}

func CommentList(currentId int64, comments []*db.CommentRaw, userMap map[int64]*db.UserRaw, relationMap map[int64]*db.RelationRaw) []*comment.Comment {
	commentList := make([]*comment.Comment, 0)
	for _, commentRaw := range comments {
		commentUser, ok := userMap[commentRaw.UserId]
		if !ok {
			commentUser = &db.UserRaw{
				Name: "未知用户",
			}
			commentUser.ID = 0
		}

		var isFollow = false

		if currentId != -1 {
			_, ok := relationMap[commentRaw.UserId]
			if ok {
				isFollow = true
			}
		}
		userFavorCount, err := cache.NewProxyIndexMap().GetFavorCount(int64(commentUser.ID))
		if err != nil {
			userFavorCount = 0
		}
		commentList = append(commentList, &comment.Comment{
			Id: int64(commentRaw.ID),
			User: &comment.User{
				Id:            int64(commentUser.ID),
				Name:          commentUser.Name,
				FollowCount:   db.QueryFollowCount(int64(commentUser.ID)),
				FollowerCount: db.QueryFollowCount(int64(commentUser.ID)),
				FavoriteCount: userFavorCount,
				IsFollow:      isFollow,
			},

			Content:    commentRaw.Content,
			CreateDate: commentRaw.UpdatedAt.Format(constants.TimeFormat),
		})
	}
	return commentList
}
