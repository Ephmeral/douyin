package pack

import (
	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/kitex_gen/comment"
	"github.com/Ephmeral/douyin/pkg/constants"
)

func CommentInfo(commentRaw *db.CommentRaw, user *db.UserRaw) *comment.Comment {
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
			TotalFavorited:  0,
			WorkCount:       db.QueryVideoCountByUserId(int64(user.ID)),
			FavoriteCount:   0,
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

		var isFollow bool = false

		if currentId != -1 {
			_, ok := relationMap[commentRaw.UserId]
			if ok {
				isFollow = true
			}
		}

		commentList = append(commentList, &comment.Comment{
			Id: int64(commentRaw.ID),
			User: &comment.User{
				Id:            int64(commentUser.ID),
				Name:          commentUser.Name,
				FollowCount:   0,
				FollowerCount: 0,
				IsFollow:      isFollow,
			},
			Content:    commentRaw.Content,
			CreateDate: commentRaw.UpdatedAt.Format(constants.TimeFormat),
		})
	}
	return commentList
}
