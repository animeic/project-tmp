package request

type ArticlePost struct {
	Content string `form:"content" json:"content" binding:"required"`
	Html    string `form:"html" json:"html" binding:"required"`
}

func (arc ArticlePost) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"content.required": "文章内容不能为空！",
		"html.required":    "文章内容不能为空h！",
	}
}

type ArticleDetail struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

func (arc ArticleDetail) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"id.required": "文章id不存在！",
	}
}

type ArticleList struct {
	CurPage  uint `form:"cur_page" json:"cur_page" binding:"required"`
	PageSize uint `form:"page_size" json:"page_size" binding:"required"`
}

func (arc ArticleList) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"cur_page.required":  "缺失cur_page！",
		"page_size.required": "缺失page_size！",
	}
}

// article/edit [admin]
type ArticleEdit struct {
	Id      uint   `form:"id" json:"id" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Html    string `form:"html" json:"html" binding:"required"`
}

func (arc ArticleEdit) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"id.required":      "文章id不能为空！",
		"content.required": "文章内容不能为空！",
		"html.required":    "文章内容不能为空h！",
	}
}
