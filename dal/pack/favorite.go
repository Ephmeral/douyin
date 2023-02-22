package pack

import (
	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/kitex_gen/favorite"
)

// VideoList pack video list info
<<<<<<< HEAD
func VideoList(currentId int64, videoData []*db.VideoRaw, userMap map[int64]*db.UserRaw, favoriteMap map[int64]*db.FavoriteRaw, relationMap map[int64]*db.RelationRaw) []*favorite.Video {
=======
func VideoList(currentId int64, videoData []*db.VideoRaw, userMap map[int64]*db.UserRaw, relationMap map[int64]*db.RelationRaw) []*favorite.Video {
>>>>>>> 17f73c6b966bdb94c10b445e5c0232fb8972bb5d
	videoList := make([]*favorite.Video, 0)
	for _, video := range videoData {
		videoUser, ok := userMap[video.UserId]
		if !ok {
			videoUser = &db.UserRaw{
				Name: "未知用户",
			}
			videoUser.ID = 0
		}

<<<<<<< HEAD
		var isFavorite = false
		var isFollow = false

		if currentId != -1 {
			_, ok := favoriteMap[int64(video.ID)]
			if ok {
				isFavorite = true
			}
=======
		var isFavorite bool = true
		var isFollow bool = false

		if currentId != -1 {
>>>>>>> 17f73c6b966bdb94c10b445e5c0232fb8972bb5d
			_, ok = relationMap[video.UserId]
			if ok {
				isFollow = true
			}
		}
		videoList = append(videoList, &favorite.Video{
			Id: int64(video.ID),
			Author: &favorite.User{
<<<<<<< HEAD
				Id:            int64(videoUser.ID),
				Name:          videoUser.Name,
				FollowCount:   db.QueryFollowCount(int64(videoUser.ID)),
				FollowerCount: db.QueryFollowerCount(int64(videoUser.ID)),
				IsFollow:      isFollow,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: 0, // TODO
			CommentCount:  0, // TODO
			IsFavorite:    isFavorite,
			Title:         video.Title,
=======
				Id:   int64(videoUser.ID),
				Name: videoUser.Name,
				//FollowCount:   videoUser.FollowCount,
				//FollowerCount: videoUser.FollowerCount,
				IsFollow: isFollow,
			},
			PlayUrl:  video.PlayUrl,
			CoverUrl: video.CoverUrl,
			//FavoriteCount: video.FavoriteCount,
			//CommentCount:  video.CommentCount,
			IsFavorite: isFavorite,
			Title:      video.Title,
>>>>>>> 17f73c6b966bdb94c10b445e5c0232fb8972bb5d
		})
	}

	return videoList
}
