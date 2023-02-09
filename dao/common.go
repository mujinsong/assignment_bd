package dao

type Message struct {
	ID         uint64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     uint64  `json:"user_id,omitempty"`
	ToUserId   uint64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId uint64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
