package cache

import (
	"context"
	"fmt"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func init() {
	rdb = redis.NewClient(
		&redis.Options{
			Addr:     constants.RedisAddress,
			Password: constants.RedisPassword,
			DB:       constants.RedisDB,
		})

}

var (
	proxyIndexOperation ProxyIndexMap
)

type ProxyIndexMap struct {
}

func NewProxyIndexMap() *ProxyIndexMap {
	return &proxyIndexOperation
}

// 更新点赞状态
func (i *ProxyIndexMap) UpdateFavorState(userId int64, videoId int64, state bool) error {
	tx := rdb.TxPipeline()
	keyUser := fmt.Sprintf("%d_%s", userId, "favoriteVideo")
	keyVideo := fmt.Sprintf("%d_%s", videoId, "isFavoritedBy")
	if state {
		tx.SAdd(ctx, keyUser, videoId)
		tx.SAdd(ctx, keyVideo, userId)
	} else {
		tx.SRem(ctx, keyUser, videoId)
		tx.SRem(ctx, keyVideo, userId)
	}
	_, err := tx.Exec(ctx)
	return err
}

// 得到点赞状态
func (i *ProxyIndexMap) GetFavorState(userId int64, videoId int64) bool {
	key := fmt.Sprintf("%d_%s", userId, "favoriteVideo")
	state := rdb.SIsMember(ctx, key, videoId)
	return state.Val()
}

// 获取用户的点赞总数量
func (i *ProxyIndexMap) GetFavorCount(userId int64) (int64, error) {
	key := fmt.Sprintf("%d_%s", userId, "favoriteVideo")
	cnt, err := rdb.SCard(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

// 获取一个用户所有点过赞的视频id,以切片的形式返回
func (i *ProxyIndexMap) GetFavorVideoIds(userId int64) ([]int64, error) {
	key := fmt.Sprintf("%d_%s", userId, "favoriteVideo")
	videoIdsStr, err := rdb.SMembers(ctx, key).Result()
	if err != nil {
		return nil, nil
	}
	videoIds := make([]int64, len(videoIdsStr))
	for i, str := range videoIdsStr {
		videoIds[i], err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return videoIds, err
}

// 获取一个用户所有点过赞的视频id，以哈希表的形式返回
func (i *ProxyIndexMap) GetFavorVideoIdsBySet(userId int64) (map[int64]struct{}, error) {
	list, err := i.GetFavorVideoIds(userId)
	if err != nil {
		return nil, err
	}
	videoIdsSet := make(map[int64]struct{})
	for _, videoId := range list {
		videoIdsSet[videoId] = struct{}{}
	}
	return videoIdsSet, nil
}

// 获取一个视频被点赞的总数量
func (i *ProxyIndexMap) GetVideoIsFavoritedCount(videoId int64) (int64, error) {
	key := fmt.Sprintf("%d_%s", videoId, "favoriteVideo")
	cnt, err := rdb.SCard(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
