package service

import (
	"assignment_bd/global"
	"assignment_bd/model"
)

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
func FollowAndFollowedCount(userID int64) (followCount, followedCount int64, err error) {
	global.DB.Select("COUNT(*)").Where("user_id = ? AND action_type = 1", userID).Model(model.Follow{}).Take(&followedCount)
	global.DB.Select("COUNT(*)").Where("follower_id = ? AND action_type = 1", userID).Model(model.Follow{}).Take(&followCount)
	return
}

// IsFollow 判断两者followerID是否关注masterID
func IsFollow(masterID, followerID int64) bool {
	var Type int8
	global.DB.Model(model.Follow{}).Select("action_type").Where("user_id = ? AND follower_id = ?", masterID, followerID).Take(&Type)
	if Type == 1 {
		return true
	}
	return false
}
