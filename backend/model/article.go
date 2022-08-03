package model

type Article struct {
	ID        uint   `gorm:"primary_key"`
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"not null"`
	Summary   string `gorm:"not null;type:text"`
	Content   string `gorm:"not null;type:text;index:unique;unique"`
	Html      string `gorm:"not null;type:text"`
	Views     uint   `gorm:"not null"`
	Likes     uint   `gorm:"not null"`
	Comments  uint   `gorm:"not null"`
	CreatedAt int64  `gorm:"not null;autoCreateTime"`
	UpdatedAt int64  `gorm:"not null;autoUpdateTime"`
	DeletedAt int64  `gorm:"not null;default:0"`
}
