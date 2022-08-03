package request

// 注册表单结构
type RegisterForm struct {
	Ecode    string `form:"ecode" json:"ecode" binding:"required,min=6,max=6"`
	Account  string `form:"account" json:"account" binding:"required,account"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,password"`
}

func (register RegisterForm) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"ecode.required":    "昵称不能为空",
		"account.required":  "账号不能为空",
		"account.account":   "账号4-16字母数字下划线中划线组成",
		"email.required":    "邮箱不能为空",
		"email.email":       "邮箱格式不正确",
		"password.required": "用户密码不能为空",
		"password.password": "密码是6-20位字母数字",
	}
}

// 注册结构
type Register struct {
	Account  string `json:"account"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginForm struct {
	Account  string `form:"account" json:"account" binding:"required,account"`
	Password string `form:"password" json:"password" binding:"required,password"`
}

func (login LoginForm) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"account.required":  "账号不能为空",
		"account.account":   "账号4-16字母数字下划线中划线组成",
		"password.required": "用户密码不能为空",
		"password.password": "密码是6-20位字母数字",
	}
}

type RepassForm struct {
	Ecode    string `form:"ecode" json:"ecode" binding:"required,min=6,max=6"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,password"`
}

func (rp RepassForm) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"ecode.required":    "验证码不能为空",
		"email.required":    "邮箱不能为空",
		"email.email":       "邮箱格式不正确",
		"password.required": "用户密码不能为空",
		"password.password": "密码是6-20位字母数字",
	}
}
