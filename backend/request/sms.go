package request

type RegisterSendForm struct {
	Account string `form:"account" json:"account" binding:"required,account"`
	Email   string `form:"email" json:"email" binding:"required,email"`
}

func (ssf RegisterSendForm) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"account.required": "账号不能为空！",
		"account.email":    "账号格式不正确！",
		"email.required":   "邮箱不能为空！",
		"email.email":      "邮箱格式不正确！",
	}
}

type RepassSendForm struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}

func (ssf RepassSendForm) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"email.required": "邮箱不能为空！",
		"email.email":    "邮箱格式不正确！",
	}
}
