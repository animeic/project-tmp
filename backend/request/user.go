package request

// todo 数据表接口-------------------------------------start
// /user/info
type UserInfoForm struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

func (ui UserInfoForm) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"id.required": "用户id不能为空！",
	}
}

// todo 数据表接口-------------------------------------end

// 验证获取用户信息
type Id struct {
	Id int64 `form:"id" json:"id" binding:"required,gte=0"`
}

func (id Id) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"username.required": "账号不能为空",
	}
}

type UserName struct {
	Username string `form:"username" json:"username" binding:"required,username,min=4,max=16"`
}

func (username UserName) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"username.required": "账号不能为空",
	}
}

type Email struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}

func (email Email) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"email.required": "账号不能为空！",
		"email.email":    "邮箱格式不正确！",
	}
}
