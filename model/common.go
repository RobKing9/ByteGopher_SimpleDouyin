package model


// 状态码，0-成功，其他值-失败
const (
	SCodeSuccess	= 0	// 状态返回成功
	SCodeFalse		= 1	// 状态返回失败

	FansList		= "fans"		// 粉丝列表类型
	FollowList		= "follow"		// 关注列表类型

)

type Response struct {
	StatusCode	int32	`json:"status_code"`	// 状态码，SCodeSuccess-成功，SCodeFalse-失败
	StatusMsg	string	`json:"status_msg"`		// 返回状态描述
}
