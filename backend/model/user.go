package model

type User struct {
	ID        uint   `gorm:"primary_key"`
	Account   string `gorm:"not null;index:unique;unique"`
	Email     string `gorm:"not null;index:unique;unique"`
	Password  string `gorm:"not null;type:char(64)" json:"-"`
	Avatar    string `gorm:"not null;default:ed4beb111b0b.jpg"`
	Nickname  string `gorm:"not null;default:ed4beb"`
	Age       uint   `gorm:"not null;default:1"`
	Gender    int8   `gorm:"not null;default:18"`
	CreatedAt int64  `gorm:"not null;autoCreateTime"`
	UpdatedAt int64  `gorm:"not null;autoUpdateTime"`
	DeletedAt int64  `gorm:"not null;default:0"`
}
