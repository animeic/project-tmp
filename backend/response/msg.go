package response

type AnErr struct {
	Code    int
	Message string
}

var (

	// common
	DataNotFundErr AnErr = AnErr{Code: 4000, Message: "数据查找失败！"}

	// Auth //业务错误从1100开始 系统错误1000~1099
	IsAccountErr     AnErr = AnErr{Code: 1100, Message: "账号已注册！"}
	IsEmailErr       AnErr = AnErr{Code: 1101, Message: "邮箱已注册！"}
	IsNotRegisterErr AnErr = AnErr{Code: 1101, Message: "该邮箱未注册！"}
	GenTokenErr      AnErr = AnErr{Code: 1102, Message: "令牌生成失败！"}
	CheckAccountErr  AnErr = AnErr{Code: 1103, Message: "账号或密码错误！"}
	RepassErr        AnErr = AnErr{Code: 1103, Message: "更新密码失败！"}

	RegisterSucc string = "注册成功！"
	LoginSucc    string = "登录成功！"
	RepassSucc   string = "重置密码成功！"
	SendCodeSucc string = "验证码发送成功！"
	GetDataSucc  string = "获取数据成功！"
	CommentSucc  string = "评论成功！"

	// sms
	SendCodeErr  AnErr = AnErr{Code: 1104, Message: "发送验证码失败！"}
	EmailCodeErr AnErr = AnErr{Code: 1105, Message: "邮箱验证码错误！"}

	ValdateFormErrCode int = 1001
	WebSocketCode      int = 1002

	GetUserInfoSuccess string = "获取用户信息成功！"

	UploadImagSucc  string = "上传图片成功"
	ArticlePostSucc string = "文章保存成功！"
	ArticleGetSucc  string = "获取文章信息成功！"
	GetListSucc     string = "获取列表成功！"
	GetUserInfoSucc string = "获取用户信息成功！"

	ImageSuffixErr   AnErr = AnErr{Code: 2001, Message: "图片格式不正确！"}
	ImageSizeErr     AnErr = AnErr{Code: 2002, Message: "图片超过10M！"}
	ImageParamErr    AnErr = AnErr{Code: 2003, Message: "图片参数非法！"}
	UploadFileErr    AnErr = AnErr{Code: 2004, Message: "上传文件不合法！"}
	TextNotChangeErr AnErr = AnErr{Code: 2005, Message: "内容未发生变化！"}
	ParseTextErr     AnErr = AnErr{Code: 2006, Message: "解析数据失败！"}
	UpdataDataErr    AnErr = AnErr{Code: 2006, Message: "更新数据失败！"}

	// user
	// GetUserInfoErr     AnErr = AnErr{Code: 2006, Message: "用户数据"}

	ParamsErr        AnErr = AnErr{Code: 1010, Message: "参数错误！"}
	RegisterError    AnErr = AnErr{Code: 2001, Message: "注册失败！"}
	RegisterUserErr  AnErr = AnErr{Code: 2004, Message: "用户信息注册失败！"}
	GetUserInfoErr   AnErr = AnErr{Code: 2006, Message: "获取用户信息失败！"}
	ExistUserErr     AnErr = AnErr{Code: 2008, Message: "该用户已被注册！"}
	LoginPasswordErr AnErr = AnErr{Code: 2011, Message: "密码错误！"}
	RepssErr         AnErr = AnErr{Code: 2012, Message: "密码重置失败！"}
	TokenError       AnErr = AnErr{Code: 2002, Message: "无效令牌！"}

	// CommentPostErr comment
	CommentPostErr AnErr = AnErr{Code: 3001, Message: "评论失败！"}
)
