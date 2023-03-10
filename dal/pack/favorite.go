package pack

import (
	"context"
	"github.com/Ephmeral/douyin/dal/cache"
	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/kitex_gen/favorite"
)

// VideoList pack video list info
func VideoList(currentId int64, videoData []*db.VideoRaw, userMap map[int64]*db.UserRaw, relationMap map[int64]*db.RelationRaw) []*favorite.Video {
	videoList := make([]*favorite.Video, 0)
	for _, video := range videoData {
		videoUser, ok := userMap[video.UserId]
		if !ok {
			videoUser = &db.UserRaw{
				Name: "未知用户",
			}
			videoUser.ID = 0
		}

		var isFollow = false

		if currentId != -1 {
			_, ok = relationMap[video.UserId]
			if ok {
				isFollow = true
			}
		}
		videoFavorCount, err := cache.NewProxyIndexMap().GetVideoIsFavoritedCount(int64(video.ID))
		if err != nil {
			videoFavorCount = 0
		}
		userFavorCount, err := cache.NewProxyIndexMap().GetFavorCount(video.UserId)
		if err != nil {
			userFavorCount = 0
		}
		videoList = append(videoList, &favorite.Video{
			Id: int64(video.ID),
			Author: &favorite.User{
				Id:             int64(videoUser.ID),
				Name:           videoUser.Name,
				FollowCount:    db.QueryFollowCount(int64(videoUser.ID)),
				FollowerCount:  db.QueryFollowerCount(int64(videoUser.ID)),
				IsFollow:       isFollow,
				TotalFavorited: db.QueryUserTotalFavorited(context.Background(), int64(videoUser.ID)),
				WorkCount:      db.QueryVideoCountByUserId(int64(videoUser.ID)),
				FavoriteCount:  userFavorCount,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: videoFavorCount,
			IsFavorite:    true,
			Title:         video.Title,
		})
	}

	return videoList
}
