package service

import (
	"context"
	"errors"
	"github.com/Ephmeral/douyin/dal/cache"
	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/dal/pack"
	"github.com/Ephmeral/douyin/kitex_gen/favorite"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/Ephmeral/douyin/pkg/jwt"
	"sync"
)

type FavoriteListService struct {
	ctx context.Context
}

// NewFavoriteListService new FavoriteListService
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

// FavoriteList get video information that users like
func (s *FavoriteListService) FavoriteList(req *favorite.FavoriteListRequest) ([]*favorite.Video, error) {
	//获取用户id
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, _ := Jwt.CheckToken(req.Token)

	//检查用户是否存在
	user, err := db.QueryUserByIds(s.ctx, []int64{req.UserId})
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, errors.New("user not exist")
	}

	//获取目标用户的点赞视频id号
	videoIds, err := cache.NewProxyIndexMap().GetFavorVideoIds(req.UserId)
	if err != nil {
		return nil, err
	}

	//获取点赞视频的信息
	videoData, err := db.QueryVideoByVideoIds(s.ctx, videoIds)
	if err != nil {
		return nil, err
	}

	//获取点赞视频的发布用户id号
	userIds := make([]int64, 0)
	for _, video := range videoData {
		userIds = append(userIds, video.UserId)
	}

	//获取点赞视频的发布用户信息
	users, err := db.QueryUserByIds(s.ctx, userIds)
	if err != nil {
		return nil, err
	}
	userMap := make(map[int64]*db.UserRaw)
	for _, user := range users {
		userMap[int64(user.ID)] = user
	}

	var relationMap map[int64]*db.RelationRaw
	//if user not logged in
	if currentId == -1 {
		//favoriteMap = nil
		relationMap = nil
	} else {
		var wg sync.WaitGroup
		wg.Add(1)
		//var favoriteErr, relationErr error
		var relationErr error
		//获取关注信息
		go func() {
			defer wg.Done()
			relationMap, err = db.QueryRelationByIds(s.ctx, currentId, userIds)
			if err != nil {
				relationErr = err
				return
			}
		}()
		wg.Wait()
		if relationErr != nil {
			return nil, relationErr
		}

	}

	videoList := pack.VideoList(currentId, videoData, userMap, relationMap)
	return videoList, nil

}
