package service

import (
	"assignment_bd/global"
	"assignment_bd/model"
	"assignment_bd/utils"
	"time"
)

var timeMap = make(map[string]uint64)

// GetMessageList 通过自己的用户 ownID 和 别人的 otherID 来 获取聊天记录
func GetMessageList(ownID, otherID uint64) ([]model.Message, error) {
	var err error
	var messages []model.Message
	unionID := utils.GetUnionID(ownID, otherID)

	// 获取所有未读的消息
	err = global.DB.
		Where("from_user_id = ? AND to_user_id = ? AND create_time > ?", ownID, otherID, timeMap[unionID]).
		Or("from_user_id = ? AND to_user_id = ? AND create_time > ?", otherID, ownID, timeMap[unionID]).
		Order("create_time ASC").Find(&messages).Error

	//// 将所有未读的消息置为已读
	//for _, message := range messages {
	//	err = global.DB.Model(&message).Update("read_status", consts.READ).Error
	//}

	if err != nil {
		return nil, err
	}

	// 更新游标的时间戳
	timeMap[unionID] = uint64(time.Now().UnixNano())

	return messages, err
}

func SendMessage(fromUserID, toUserID uint64, content string) error {
	var err error
	unionID := utils.GetUnionID(fromUserID, toUserID)

	message := model.Message{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Content:    content,
		CreateTime: uint64(time.Now().UnixNano()),
	}

	err = global.DB.Create(&message).Error

	// 更新游标的时间戳
	timeMap[unionID] = uint64(time.Now().UnixNano())

	return err
}
