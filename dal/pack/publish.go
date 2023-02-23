package pack

import (
	"github.com/Ephmeral/douyin/dal/cache"
	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/kitex_gen/publish"
)

// PublishInfo pack video list info
func PublishInfo(currentId int64, videoData []*db.VideoRaw, userMap map[int64]*db.UserRaw, videoIdsSet map[int64]struct{}, relationMap map[int64]*db.RelationRaw) []*publish.Video {
	videoList := make([]*publish.Video, 0)
	for _, video := range videoData {
		videoUser, ok := userMap[video.UserId]
		if !ok {
			videoUser = &db.UserRaw{
				Name: "未知用户",
			}
			videoUser.ID = 0
		}

		var isFavorite = false
		var isFollow = false

		if currentId != -1 {
			_, ok := videoIdsSet[int64(video.ID)]
			if ok {
				isFavorite = true
			}
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
		videoList = append(videoList, &publish.Video{
			Id: int64(video.ID),
			Author: &publish.User{
				Id:            int64(videoUser.ID),
				Name:          videoUser.Name,
				FollowCount:   db.QueryFollowCount(int64(videoUser.ID)),
				FollowerCount: db.QueryFollowerCount(int64(videoUser.ID)),
				FavoriteCount: userFavorCount,
				IsFollow:      isFollow,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: videoFavorCount,
			CommentCount: db.QueryCommentCountByVideoId(int64(video.ID)),
			IsFavorite:   isFavorite,
			Title:        video.Title,
		})
	}

	return videoList
}
