package model

type Message struct {
	ID         uint64 `json:"id,omitempty"`
	FromUserID uint64 `json:"from_user_id"`
	ToUserID   uint64 `json:"to_user_id"`
	ReadStatus uint64 `json:"read_status"`
	Content    string `json:"content,omitempty"`
	CreateTime uint64 `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserID     uint64 `json:"user_id,omitempty"`
	ToUserId   uint64 `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId uint64 `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
