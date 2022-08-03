package model

type Comment struct {
	Id   uint `gorm:"primary_key"`
	Aid  uint `gorm:"not null;comment:回复的文章id"`
	Pid  uint `gorm:"not null;default:0;comment:二级评论时，为被评论的id"`
	Ppid uint `gorm:"not null;default:0;comment:二级评论下的评论@user，为被评论的id"`
	Uid  uint `gorm:"not null;comment:用户id"`
	// 冗余字段 用于@user功能，避免查表
	AtUid     uint   `gorm:"not null;comment:@uid,二级评论时@uid"`
	Content   string `gorm:"not null"`
	CreatedAt int64  `gorm:"not null;autoCreateTime"`
	UpdatedAt int64  `gorm:"not null;autoUpdateTime"`
	DeletedAt int64  `gorm:"not null;default:0"`
}
