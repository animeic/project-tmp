package response

type CommentPost struct {
	Id        uint   `json:"id"`
	Uid       uint   `json:"uid"`
	Aid       uint   `json:"aid"`
	Pid       uint   `json:"pid"`
	Ppid      uint   `json:"ppid"`
	AtUid     uint   `json:"at_uid"`
	Content   string `json:"content"`
	CreatedAt uint   `json:"created_at"`
}
