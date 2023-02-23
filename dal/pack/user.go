package pack

import (
	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/kitex_gen/user"
)

func UserInfo(userRaw *db.UserRaw, isFollow bool, followCount int64, followerCount int64, totalFavortited int64, workCount int64, favoriteCount int64) *user.User {
	userInfo := &user.User{
		Id:              int64(userRaw.ID),
		Name:            userRaw.Name,
		FollowCount:     followCount,
		FollowerCount:   followerCount,
		IsFollow:        isFollow,
		Avatar:          userRaw.Avatar,
		BackgroundImage: userRaw.BackgroundImage,
		Signature:       userRaw.Signature,
		TotalFavorited:  totalFavortited,
		WorkCount:       workCount,
		FavoriteCount:   favoriteCount,
	}
	return userInfo
}
