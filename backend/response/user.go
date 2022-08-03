package response

import "strconv"

// new 用户简易信息
type UserInfo struct {
	Id       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

// jwt接口方法
func (user UserInfo) GetUid() string {
	return strconv.Itoa(int(user.Id))
}
