package pack

import (
	"github.com/Ephmeral/douyin/dal/db"
	"github.com/Ephmeral/douyin/kitex_gen/relation"
)

func UserList(currentId int64, users []*db.UserRaw, relationMap map[int64]*db.RelationRaw) []*relation.User {
	userList := make([]*relation.User, 0)
	for _, user := range users {
		var isFollow = false

		if currentId != -1 {
			_, ok := relationMap[int64(user.ID)]
			if ok {
				isFollow = true
			}
		}
		userList = append(userList, &relation.User{
			Id:            int64(user.ID),
			Name:          user.Name,
			FollowCount:   db.QueryFollowCount(int64(user.ID)),
			FollowerCount: db.QueryFollowerCount(int64(user.ID)),
			IsFollow:      isFollow,
		})
	}
	return userList
}

func FriendList(currentId int64, users []*db.UserRaw, messageMap map[int64]*db.MessageRaw) []*relation.FriendUser {
	userList := make([]*relation.FriendUser, 0)
	for _, user := range users {
		var msg string
		var msgType int64
		ret, ok := messageMap[int64(user.ID)]
		if ok {
			msg = ret.Content
			if ret.ToUserId == currentId {
				msgType = 0
			} else {
				msgType = 1
			}
		}
		userList = append(userList, &relation.FriendUser{
			Id:            int64(user.ID),
			Name:          user.Name,
			FollowCount:   db.QueryFollowCount(int64(user.ID)),
			FollowerCount: db.QueryFollowerCount(int64(user.ID)),
			IsFollow:      true,
			Message:       msg,
			MsgType:       msgType,
		})
	}
	return userList
}
