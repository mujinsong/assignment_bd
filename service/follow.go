package service

import (
	"assignment_bd/dao"
	"assignment_bd/global"
)

// GetFollowStatusForUpdate 获取关注状态，此处是针对 AddFollow 和 CancelFollow
//func GetFollowStatusForUpdate(followerID, userID uint64) (bool, error) {
//	var followList []dao.Follow
//	if result := global.DB.Select("user_id", "action_type").Model(&dao.Follow{}).
//		Where("follower_id = ?", followerID).Find(&followList); result.Error != nil {
//		return false, result.Error
//	}
//	return true, nil
//}

// GetFollowStatus 获取关注状态，此处是针对非更新操作
func GetFollowStatus(followerID, userID uint64) (bool, error) {
	//followStatus, err := GetFollowStatusForUpdate(followerID, userID)
	//if err == nil || err.Error() == "no tracking information" {
	//	return followStatus, nil
	//}
	//return false, err
	followst := &dao.Follow{}
	if result := global.DB.Select("user_id", "action_type").Model(&dao.Follow{}).
		Where("follower_id = ? AND user_id = ?", followerID, userID).Take(followst); result.Error != nil {
		return false, result.Error

	}
	if followst.ActionType == 0 {
		return false, nil
	}
	return true, nil
}

func GetFollowStatusList(followerID uint, userIDList []uint, isfollowerList []bool) error {
	var temp []dao.Follow
	if result := global.DB.Select("user_id", "action_type").Model(&dao.Follow{}).
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
