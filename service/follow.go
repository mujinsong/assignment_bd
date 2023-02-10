package service

import (
	"assignment_bd/consts"
	"assignment_bd/global"
	"assignment_bd/model"
)

// Follow 关注用户
// userID 是被关注者id，followerID是关注者的ID（也就是当前登录用户的ID）
func Follow(userID, followerID uint64) error {
	var err error
	var follow model.Follow

	// 看看之前有没有这条记录
	errFind := global.DB.
		Where("user_id = ?", userID).
		Where("follower_id = ?", followerID).First(&follow).Error

	// 生成需要的列
	follow = model.Follow{
		UserID:     uint(userID),
		FollowerID: uint(followerID),
		ActionType: consts.FOLLOW,
	}

	if errFind != nil { // 如果之前没有这条记录，就插入
		err = global.DB.Create(&follow).Error
	} else { // 如果有，就更新这条记录（以前可能是未关注的状态）
		err = global.DB.Model(&model.Follow{}).
			Where("user_id", userID).
			Where("follower_id", followerID).
			Update("action_type", consts.FOLLOW).Error
	}

	// 如果没有出错，就更新user信息
	// user1 是被关注的用户，user2是当前登录用户（也就是user1的粉丝）
	var user1, user2 model.User

	// user1 的粉丝数 += 1
	global.DB.First(&user1, userID)
	global.DB.Model(&user1).Update("follower_count", user1.FollowerCount+1)

	// user2 的关注数 += 1
	global.DB.First(&user2, followerID)
	global.DB.Model(&user2).Update("follow_count", user2.FollowCount+1)

	return err
}

// UnFollow 和 UnFollow 基本一样， 取消关注用户
func UnFollow(userID, followerID uint64) error {
	var err error
	var follow model.Follow

	// 看看之前有没有这条记录
	errFind := global.DB.
		Where("user_id = ?", userID).
		Where("follower_id = ?", followerID).First(&follow).Error

	// 生成需要的列
	follow = model.Follow{
		UserID:     uint(userID),
		FollowerID: uint(followerID),
		ActionType: consts.UNFOLLOW,
	}

	if errFind != nil { // 如果之前没有这条记录，就插入
		err = global.DB.Create(&follow).Error
	} else { // 如果有，就更新这条记录（以前可能是未关注的状态）
		err = global.DB.Model(&model.Follow{}).
			Where("user_id", userID).
			Where("follower_id", followerID).
			Update("action_type", consts.UNFOLLOW).Error
	}

	// 如果没有出错，就更新user信息
	// user1 是被关注的用户，user2是当前登录用户（也就是user1的粉丝）
	var user1, user2 model.User

	// user1 的粉丝数 -= 1
	global.DB.First(&user1, userID)
	global.DB.Model(&user1).Update("follower_count", user1.FollowerCount-1)

	// user2 的关注数 -= 1
	global.DB.First(&user2, followerID)
	global.DB.Model(&user2).Update("follow_count", user2.FollowCount-1)

	return err
}

// GetFollowList 用来获取用户关注列表（参数userID是当前登录用户的ID）
func GetFollowList(userID uint64) ([]model.UserInfo, error) {
	var err error
	var users []model.UserInfo
	var follows []model.Follow

	// 先从 follows 表中查找出当前用户的所有关注的人
	err = global.DB.Where("follower_id = ?", userID).Where("action_type = ?", 1).Find(&follows).Error
	users = make([]model.UserInfo, len(follows))

	for i, follow := range follows {
		var user model.User
		err = global.DB.Where("id = ?", follow.UserID).Find(&user).Error

		// todo 查询关注数（待写）

		userInfo := model.UserInfo{
			ID:   user.ID,
			Name: user.Username,
			//FollowCount:   233,
			//FollowerCount: 233,
			IsFollow: true,
		}

		users[i] = userInfo
	}

	return users, err
}

// GetFollowerList 用来获取用户粉丝列表（参数userID是当前登录用户的ID）
func GetFollowerList(userID uint64) ([]model.UserInfo, error) {
	var err error
	var users []model.UserInfo
	var follows []model.Follow

	// 先从 follows 表中查找出当前用户的所有粉丝
	err = global.DB.Where("user_id = ?", userID).Where("action_type = ?", 1).Find(&follows).Error
	users = make([]model.UserInfo, len(follows))

	for i, follow := range follows {
		var user model.User
		var checkFollow model.Follow
		err = global.DB.Where("id = ?", follow.FollowerID).Find(&user).Error

		// record not found 未找到记录说明未关注当前粉丝
		global.DB.
			Where("user_id", follow.FollowerID).
			Where("follower_id = ?", userID).
			Where("action_type = ?", 1).First(&checkFollow)

		// todo 查询粉丝数（待写）

		userInfo := model.UserInfo{
			ID:   user.ID,
			Name: user.Username,
			//FollowCount:   233,
			//FollowerCount: 233,
			IsFollow: checkFollow.ID != 0,
		}

		users[i] = userInfo
	}

	return users, err
}

/*

==========================================================================================

*/

// GetFollowStatus 获取关注状态，此处是针对非更新操作
//func GetFollowStatus(followerID, userID uint64) (bool, error) {
//	//followStatus, err := GetFollowStatusForUpdate(followerID, userID)
//	//if err == nil || err.Error() == "no tracking information" {
//	//	return followStatus, nil
//	//}
//	//return false, err
//	followst := &model.Follow{}
//	if result := global.DB.Select("user_id", "action_type").Model(&model.Follow{}).
//		Where("follower_id = ? AND user_id = ?", followerID, userID).Take(followst); result.Error != nil {
//		return false, result.Error
//
//	}
//	if followst.ActionType == 0 {
//		return false, nil
//	}
//	return true, nil
//}

func GetFollowStatusList(followerID uint, userIDList []uint, isfollowerList []bool) error {
	var temp []model.Follow
	if result := global.DB.Select("user_id", "action_type").Model(&model.Follow{}).
		Where("follower_id = ? AND user_id IN ?", followerID, userIDList).Find(&temp); result.Error != nil {
		return result.Error
	}
	isfollowerList = make([]bool, 0, len(userIDList))
	mp := make(map[uint]bool)
	for _, v := range temp {
		mp[v.UserID] = func() bool {
			if v.ActionType == 0 {
				return false
			} else {
				return true
			}
		}()
	}
	for _, v := range userIDList {
		isfollowerList = append(isfollowerList, mp[v])
	}
	return nil
}

// FollowAndFollowedCount 获取userID的所有关注和被关注人数
func FollowAndFollowedCount(userID uint64) (followCount, followedCount uint64, err error) {
	global.DB.Select("COUNT(*)").Where("user_id = ? AND action_type = 1", userID).Model(model.Follow{}).Take(&followedCount)
	global.DB.Select("COUNT(*)").Where("follower_id = ? AND action_type = 1", userID).Model(model.Follow{}).Take(&followCount)
	return
}

// IsFollow 判断两者followerID是否关注masterID
// func IsFollow(masterID, followerID uint64) bool {
// 	var Type int8
// 	global.DB.Model(model.Follow{}).Select("action_type").Where("user_id = ? AND follower_id = ?", masterID, followerID).Take(&Type)
// 	if Type == 1 {
// 		return true
// 	}
// 	return false
// }

/*
判断是否关注该用户
uid  当前登录的用户id
userID 表示被查询的用户id
*/
func IsFollow(userID, uid uint64) bool {
	// 自己肯定关注自己
	//if userID == uid {
	//	return true
	//}
	//var likeLog model.Follow
	//global.DB.Model(model.Follow{}).Where("user_id = ? AND follower_id = ?", userID, uid).Take(&likeLog)
	//if likeLog.ActionType == 1 {
	//	return true
	//} else {
	//	return false
	//}7
	return true
}
