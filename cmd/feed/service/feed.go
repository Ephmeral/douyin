package service

import (
	"context"
	"github.com/Ephmeral/douyin/dal/cache"
	"sync"

	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/dal/pack"
	"github.com/Ephmeral/douyin/kitex_gen/feed"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/Ephmeral/douyin/pkg/jwt"
)

type FeedService struct {
	ctx context.Context
}

// NewFeedService 创建新的FeedService
func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx: ctx}
}

// Feed 用于获取多个视频信息列表
func (s *FeedService) Feed(req *feed.FeedRequest) ([]*feed.Video, int64, error) {
	// 初始化 JWT 对象
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))

	// 检查用户登录状态
	currentId, _ := Jwt.CheckToken(req.Token)

	// 查询视频信息
	videoData, err := db.QueryVideoByLatestTime(s.ctx, req.LatestTime)
	if err != nil {
		return nil, 0, err
	}

	// 查询视频 ID 和用户 ID
	videoIds := make([]int64, 0)
	userIds := make([]int64, 0)
	for _, video := range videoData {
		videoIds = append(videoIds, int64(video.ID))
		userIds = append(userIds, video.UserId)
	}

	// 查询用户信息
	users, err := db.QueryUserByIds(s.ctx, userIds)
	if err != nil {
		return nil, 0, err
	}
	userMap := make(map[int64]*db.UserRaw)
	for _, user := range users {
		userMap[int64(user.ID)] = user
	}

	var relationMap map[int64]*db.RelationRaw
	// 如果用户未登录
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
			return nil, 0, relationErr
		}

	}
	// 从redis中查找当前用户所有点赞的视频集合
	var videoIdsSet map[int64]struct{}
	if currentId == -1 {
		//当前用户未登录
		videoIdsSet = make(map[int64]struct{}, 0)
	} else {
		videoIdsSet, err = cache.NewProxyIndexMap().GetFavorVideoIdsBySet(currentId)
		if err != nil {
			return nil, 0, err
		}
	}
	videos, nextTime := pack.VideoInfo(currentId, videoData, videoIdsSet, userMap, relationMap)
	return videos, nextTime, nil
}
