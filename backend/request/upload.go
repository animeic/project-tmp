package request

import "mime/multipart"

type ImageUpload struct {
	Name  string                `form:"name" json:"name" binding:"required"`
	Image *multipart.FileHeader `json:"image"`
}

func (iu ImageUpload) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required": "名字不能为空！",
		// "image.required": "图片不能为",
	}
}
