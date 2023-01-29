package service

import (
	"assignment_bd/dao"
	"assignment_bd/global"
)

// GetFollowStatusForUpdate 获取关注状态，此处是针对 AddFollow 和 CancelFollow
func GetFollowStatusForUpdate(followerID, celebrityID uint64) (bool, error) {
	var followList []dao.Follow
	if result := global.DB.Select("celebrity_id", "is_follow").Model(&dao.Follow{}).
		Where("follower_id = ?", followerID).Find(&followList); result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// GetFollowStatus 获取关注状态，此处是针对非更新操作
func GetFollowStatus(followerID, celebrityID uint64) (bool, error) {
	followStatus, err := GetFollowStatusForUpdate(followerID, celebrityID)
	if err == nil || err.Error() == "no tracking information" {
		return followStatus, nil
	}
	return false, err
}
