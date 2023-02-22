package cache

import (
	"fmt"
	"testing"
)

func TestProxyIndexMap_UpdateFavorState(t *testing.T) {
	NewProxyIndexMap().UpdateFavorState(2, 2, true)
	NewProxyIndexMap().UpdateFavorState(2, 3, true)
	NewProxyIndexMap().UpdateFavorState(2, 4, true)
	fmt.Println("用户1对视频2的状态是：", NewProxyIndexMap().GetFavorState(1, 2))
	fmt.Println("用户1对视频3的状态是：", NewProxyIndexMap().GetFavorState(1, 3))
	list, err := NewProxyIndexMap().GetFavorVideoIds(2)
	if err != nil {
		panic(err)
	}
	fmt.Println("用户2喜欢的视频列表为：", list)
}
