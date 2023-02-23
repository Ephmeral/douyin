package pack

import (
	"github.com/Ephmeral/douyin/dal/cache"
	"time"

	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/kitex_gen/feed"
)

// VideoInfo pack video list info
func VideoInfo(currentId int64, videoData []*db.VideoRaw, videoIdsSet map[int64]struct{}, userMap map[int64]*db.UserRaw, relationMap map[int64]*db.RelationRaw) ([]*feed.Video, int64) {
	videoList := make([]*feed.Video, 0)
	var nextTime int64
	for _, video := range videoData {
		videoUser, ok := userMap[video.UserId]
		if !ok {
			videoUser = &db.UserRaw{
				Name: "未知用户",
				//FollowCount:   0,
				//FollowerCount: 0,
			}
			videoUser.ID = 0
		}

		var isFavorite bool = false
		var isFollow bool = false

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
		userFavorCount, err := cache.NewProxyIndexMap().GetFavorCount(video.UserId)
		if err != nil {
			userFavorCount = 0
		}
		videoFavorCount, err := cache.NewProxyIndexMap().GetVideoIsFavoritedCount(int64(video.ID))
		if err != nil {
			videoFavorCount = 0
		}
		videoList = append(videoList, &feed.Video{
			Id: int64(video.ID),
			Author: &feed.User{
				Id:            int64(videoUser.ID),
				Name:          videoUser.Name,
				FollowCount:   db.QueryFollowCount(int64(videoUser.ID)),
				FollowerCount: db.QueryFollowerCount(int64(videoUser.ID)),
				IsFollow:      isFollow,
				FavoriteCount: userFavorCount,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: videoFavorCount,
			//CommentCount:  video.CommentCount,
			IsFavorite: isFavorite,
			Title:      video.Title,
		})
	}

	if len(videoData) == 0 {
		nextTime = time.Now().UnixMilli()
	} else {
		nextTime = videoData[len(videoData)-1].UpdatedAt.UnixMilli()
	}

	return videoList, nextTime
}
