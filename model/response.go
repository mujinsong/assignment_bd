package model

// 取代了 backend 包，将封装的所有的数据结构都放在这里
// 并指明了各个响应对应的接口

// Response 所有响应的公共信息，包括状态码和状态信息
/*
	/douyin/publish/action/
	/douyin/favorite/action/
	/douyin/relation/action/
	/douyin/message/action/
*/
type Response struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// FeedResponse Feed流需要的响应
/*
	/douyin/feed/
*/
type FeedResponse struct {
	Response
	NextTime  uint64      `json:"next_time,omitempty"`
	VideoList []VideoInfo `json:"video_list,omitempty"`
}

// UserLoginResponse 用户 注册 或者 登录 响应的结构体
/*
	/douyin/user/register/
	/douyin/user/login/
*/
type UserLoginResponse struct {
	Response
	UserID uint64 `json:"user_id"`
	Token  string `json:"token"`
}

// UserInfoResponse 是响应的用户信息，是 UserInfo 的详细信息（关注数等等）
/*
	/douyin/user/
*/
type UserInfoResponse struct {
	Response
	UserInfo *UserInfo `json:"user"` // 用户信息
}

// VideoListResponse (用户发布过的视频的列表 和 用户点赞过的视频的列表) 的响应
/*
	/douyin/publish/list/
	/douyin/favorite/list/
*/
type VideoListResponse struct {
	Response
	VideoList []VideoInfo `json:"video_list"` // 用户发布的视频列表
}

// CommentResponse 是用户对视频评论操作的响应
/*
	/douyin/comment/action/
*/
type CommentResponse struct {
	Response
	Comment Comment `json:"comment"` // 评论成功返回评论内容，不需要重新拉取整个列表
}

// CommentListResponse 是视频的所有评论的响应（按发布时间倒序）
/*
	/douyin/comment/list/
*/
type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list"` // 评论列表
}

// UserListResponse 是返回用户列表的响应，可用来获取关注列表，粉丝列表，好友列表等等
/*
	/douyin/relation/follow/list/
	/douyin/relation/follower/list/
*/
type UserListResponse struct {
	Response
	UserList []UserInfo `json:"user_list"` // 用户信息列表
}

// FriendListResponse 是朋友列表响应的接口，FriendUser 是 UserInfo 的继承，除了基本内容外，还多了聊天记录
/*
	/douyin/relation/friend/list/
*/
type FriendListResponse struct {
	Response
	FriendList []FriendUser `json:"user_list"`
}

// MessageResponse 用户之间聊天信息的响应
/*
	/douyin/message/chat/
*/
type MessageResponse struct {
	Response
	MessageList []Message `json:"message_list"` // 用户列表
}
