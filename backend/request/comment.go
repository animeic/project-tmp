package request

type CommentPost struct {
	Aid     uint   `form:"aid" json:"aid" binding:"required"`
	Pid     uint   `form:"pid" json:"pid" binding:"required"`
	Ppid    uint   `form:"ppid" json:"ppid" binding:"required"`
	AtUid   uint   `form:"at_uid" json:"at_uid" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

func (ctp CommentPost) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"aid.required":     "文章id不能为空！",
		"pid.required":     "回复对象id不能为空！",
		"ppid.required":    "二级回复对象id不能为空！",
		"at_uid.required":  "回复@uid不能为空！",
		"content.required": "文章内容不能为空！",
	}
}

type CommentList struct {
	Aid uint `form:"aid" json:"aid" binding:"required"`
}

func (ctp CommentList) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"aid.required": "文章id不能为空！",
	}
}
