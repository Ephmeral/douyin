package db

import (
	"context"
	"github.com/Ephmeral/douyin/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// RelationRaw Relation Gorm Data Structures
type RelationRaw struct {
	gorm.Model
	UserId   int64 `gorm:"column:user_id;not null;index:idx_userid"`
	ToUserId int64 `gorm:"column:to_user_id;not null;index:idx_touserid"`
}

func (RelationRaw) TableName() string {
	return constants.ReltaionTableName
}

// QueryRelationByIds 根据当前用户id和目标用户id获取关注信息
func QueryRelationByIds(ctx context.Context, currentId int64, toUserIds []int64) (map[int64]*RelationRaw, error) {
	if currentId == -1 {
		return nil, nil
	}
	var relations = make([]*RelationRaw, 0)
	err := DB.WithContext(ctx).Where("user_id = ? AND to_user_id IN ?", currentId, toUserIds).Find(&relations).Error
	if err != nil {
		klog.Error("query relation by ids " + err.Error())
		return nil, err
	}
	relationMap := make(map[int64]*RelationRaw)
	for _, relation := range relations {
		relationMap[relation.ToUserId] = relation
	}
	return relationMap, nil
}

// AddRelation 添加一条关注记录，表示currentId关注了toUserId
func AddRelation(ctx context.Context, currentId int64, toUserId int64) error {
	relationRaw := &RelationRaw{
		UserId:   currentId,
		ToUserId: toUserId,
	}
	err := DB.WithContext(ctx).Table("relation").Create(&relationRaw).Error
	if err != nil {
		klog.Error("create relation record fail " + err.Error())
		return err
	}
	return nil
}

// DeleteRelation 删除由currentId到toUserId的关注记录
func DeleteRelation(ctx context.Context, currentId int64, toUserId int64) error {
	var relationRaw *RelationRaw
	err := DB.WithContext(ctx).Table("relation").Where("user_id = ? AND to_user_id = ?", currentId, toUserId).Delete(&relationRaw).Error
	if err != nil {
		klog.Error("delete relation record fail " + err.Error())
		return err
	}
	return nil
}

// QueryFollowById 通过用户id，查询该用户关注的用户，返回两者之间的关注记录
func QueryFollowById(ctx context.Context, userId int64) ([]*RelationRaw, error) {
	var relations []*RelationRaw
	err := DB.WithContext(ctx).Table("relation").Where("user_id = ?", userId).Find(&relations).Error
	if err != nil {
		klog.Error("query follow by id fail " + err.Error())
		return nil, err
	}
	return relations, nil
}

// QueryFollowerById 通过用户id，查询该用户的粉丝，返回两者之间的关注记录
func QueryFollowerById(ctx context.Context, userId int64) ([]*RelationRaw, error) {
	var relations []*RelationRaw
	err := DB.WithContext(ctx).Table("relation").Where("to_user_id = ?", userId).Find(&relations).Error
	if err != nil {
		klog.Error("query follower by id fail " + err.Error())
		return nil, err
	}
	return relations, nil
}

// QueryFriendById 通过用户id，查询与之互相关注的用户Id
func QueryFriendById(ctx context.Context, currentId int64, toUserId int64) ([]int64, error) {
	var usersId []int64
	err := DB.WithContext(ctx).Raw("SELECT DISTINCT r1.to_user_id FROM relation r1 JOIN relation r2 ON r1.to_user_id = r2.user_id WHERE r1.user_id = ? AND r2.to_user_id = ?", currentId, toUserId).Find(&usersId).Error
	if err != nil {
		klog.Error("query friend by id fail " + err.Error())
		return nil, err
	}
	return usersId, nil
}

// QueryFollowCount 根据用户id，查询该用户关注其他用户的总数
func QueryFollowCount(userId int64) int64 {
	var followCount int64
	err := DB.Table(constants.ReltaionTableName).Where("user_id = ?", userId).Count(&followCount).Error
	if err != nil {
		return 0
	}
	return followCount
}

// QueryFollowerCount 根据用户id，查询该用户粉丝的总数
func QueryFollowerCount(userId int64) int64 {
	var followerCount int64
	err := DB.Table(constants.ReltaionTableName).Where("to_user_id = ?", userId).Count(&followerCount).Error
	if err != nil {
		return 0
	}
	return followerCount
}
