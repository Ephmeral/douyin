package service

import (
	"context"
	"errors"
	"github.com/Ephmeral/douyin/dal/cache"
	"sync"

	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/dal/pack"
	"github.com/Ephmeral/douyin/kitex_gen/publish"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/Ephmeral/douyin/pkg/jwt"
)

type PublishListService struct {
	ctx context.Context
}

// NewPublishListService new PublishService
func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

// PublishList get publish list by userid
func (s *PublishListService) PublishList(req *publish.PublishListRequest) ([]*publish.Video, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, _ := Jwt.CheckToken(req.Token)

	videoData, err := db.QueryVideoByUserId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	videoIds := make([]int64, 0)
	userIds := []int64{req.UserId}
	for _, video := range videoData {
		videoIds = append(videoIds, int64(video.ID))
	}

	users, err := db.QueryUserByIds(s.ctx, userIds)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("user not exist")
	}
	userMap := make(map[int64]*db.UserRaw)
	for _, user := range users {
		userMap[int64(user.ID)] = user
	}

	var relationMap map[int64]*db.RelationRaw
	if currentId == -1 {
		relationMap = nil
	} else {
		var wg sync.WaitGroup
		wg.Add(1)
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
	// 从redis中查找当前用户所有点赞的视频集合
	var videoIdsSet map[int64]struct{}
	if currentId == -1 {
		videoIdsSet = make(map[int64]struct{}, 0)
	} else {
		videoIdsSet, err = cache.NewProxyIndexMap().GetFavorVideoIdsBySet(currentId)
		if err != nil {
			return nil, err
		}
	}
	videoList := pack.PublishInfo(currentId, videoData, userMap, videoIdsSet, relationMap)
	return videoList, nil
}
