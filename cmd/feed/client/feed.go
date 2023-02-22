package client

import (
	"context"
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

	var favoriteMap map[int64]*db.FavoriteRaw
	var relationMap map[int64]*db.RelationRaw
	// 如果用户未登录
	if currentId == -1 {
		favoriteMap = nil
		relationMap = nil
	} else {
		var wg sync.WaitGroup
		wg.Add(2)
		var favoriteErr, relationErr error
		//获取点赞信息
		go func() {
			defer wg.Done()
			favoriteMap, err = db.QueryFavoriteByIds(s.ctx, currentId, videoIds)
			if err != nil {
				favoriteErr = err
				return
			}
		}()
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
		if favoriteErr != nil {
			return nil, 0, favoriteErr
		}
		if relationErr != nil {
			return nil, 0, relationErr
		}

	}
	// 打包视频信息
	videos, nextTime := pack.VideoInfo(currentId, videoData, userMap, favoriteMap, relationMap)
	return videos, nextTime, nil
}
