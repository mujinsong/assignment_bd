package service

import (
	"assignment_bd/consts"
	"assignment_bd/global"
	"assignment_bd/model"
	"time"
)

// GetMessageList 通过自己的用户 ownID 和 别人的 otherID 来 获取聊天记录
func GetMessageList(ownID, otherID uint64) ([]model.Message, error) {
	var err error
	var messages []model.Message

	// 获取所有未读的消息
	err = global.DB.
		Where("from_user_id = ? AND to_user_id = ? AND read_status = ?", ownID, otherID, consts.UNREAD).
		Or("from_user_id = ? AND to_user_id = ? AND read_status = ?", otherID, ownID, consts.UNREAD).
		Order("create_time ASC").Find(&messages).Error

	// 将所有未读的消息置为已读
	for _, message := range messages {
		err = global.DB.Model(&message).Update("read_status", consts.READ).Error
	}

	if err != nil {
		return nil, err
	}

	return messages, err
}

func SendMessage(fromUserID, toUserID uint64, content string) error {
	var err error
	message := model.Message{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		ReadStatus: consts.UNREAD,
		Content:    content,
		CreateTime: uint64(time.Now().Unix()),
	}

	err = global.DB.Create(&message).Error

	return err
}
