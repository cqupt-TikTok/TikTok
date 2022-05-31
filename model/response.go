package model

// Response 基本响应
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// FeedResponse 视频流响应
type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"` // 视频列表
	NextTime  int64   `json:"next_time,omitempty"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

// UserResponse 用户注册，登录响应
type UserResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

// UserInfoResponse 用户信息响应
type UserInfoResponse struct {
	Response
	User `json:"user"` //用户信息
}

// PostListResponse 发布列表响应
type PostListResponse struct {
	Response
	VideoList []Video `json:"video_list"` // 用户发布的视频列表
}

// PointListResponse 点赞列表响应
type PointListResponse struct {
	Response
	VideoList []Video `json:"video_list"` // 用户发布的视频列表
}

// CommentActionResponse 评论操作响应
type CommentActionResponse struct {
	Response
	Comment `json:"comment"` // 评论成功返回评论内容，不需要重新拉取整个列表
}

// CommentListResponse 评论列表响应
type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list"` // 评论列表
}

// FollowListResponse 关注列表响应
type FollowListResponse struct {
	Response
	UserList []User `json:"user_list"` // 用户信息列表
}

// FollowerListResponse 粉丝列表响应
type FollowerListResponse struct {
	Response
	UserList []User `json:"user_list"` // 用户列表
}
