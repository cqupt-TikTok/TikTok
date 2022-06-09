package model

// BaseResponse 基本响应
type BaseResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// FeedResponse 视频流响应
type FeedResponse struct {
	BaseResponse
	VideoList []VideoResp `json:"video_list"` // 视频列表
	NextTime  int64       `json:"next_time"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

// UserResponse 用户注册，登录响应
type UserResponse struct {
	BaseResponse
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}

// UserInfoResponse 用户信息响应
type UserInfoResponse struct {
	BaseResponse
	UserResp `json:"user"` //用户信息
}

// PostListResponse 发布列表响应
type PostListResponse struct {
	BaseResponse
	VideoList []VideoResp `json:"video_list"` // 用户发布的视频列表
}

// FavoriteListResponse 点赞列表响应
type FavoriteListResponse struct {
	BaseResponse
	VideoList []VideoResp `json:"video_list"` // 用户发布的视频列表
}

// CommentActionResponse 评论操作响应
type CommentActionResponse struct {
	BaseResponse
	CommentResp `json:"comment"` // 评论成功返回评论内容，不需要重新拉取整个列表
}

// CommentListResponse 评论列表响应
type CommentListResponse struct {
	BaseResponse
	CommentList []CommentResp `json:"comment_list"` // 评论列表
}

// FollowListResponse 关注列表响应
type FollowListResponse struct {
	BaseResponse
	UserList []UserResp `json:"user_list"` // 用户信息列表
}

// FollowerListResponse 粉丝列表响应
type FollowerListResponse struct {
	BaseResponse
	UserList []UserResp `json:"user_list"` // 用户列表
}
