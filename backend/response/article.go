package response

import (
	"encoding/json"
)

// post
type ArticlePost struct {
	Id  uint `json:"id"`
	Uid uint `json:"uid"`
	// Summary   string `json:"summary"`
	// Views     uint   `json:"views"`
	// Likes     uint   `json:"likes"`
	// Comments  uint   `json:"comments"`
	// CreatedAt uint   `json:"created_at"`
}

// list
type AtInfo struct {
	Id        uint   `json:"id"`
	Uid       uint   `json:"uid"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Views     uint   `json:"views"`
	Likes     uint   `json:"likes"`
	Comments  uint   `json:"comments"`
	CreatedAt uint   `json:"created_at"`
	UserInfo  `json:"user_info"`
}

type ArticleList struct {
	Lists []AtInfo `json:"lists"`
	Total uint     `json:"total"`
}

// detail
type ArticleDetail struct {
	Id        uint   `json:"id"`
	Uid       uint   `json:"uid"`
	Html      string `json:"html"`
	Views     uint   `json:"views"`
	Likes     uint   `json:"likes"`
	Comments  uint   `json:"comments"`
	CreatedAt uint   `json:"created_at"`
	UserInfo  `json:"user_info"`
}

func (ard *ArticleDetail) MarshalBinary() ([]byte, error) {
	return json.Marshal(ard)
}
func (ard *ArticleDetail) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, ard)
}

//  article/editDetail
type AtEditDetail struct {
	Id      uint   `json:"id"`
	Uid     uint   `json:"uid"`
	Content string `json:"content"`
}
