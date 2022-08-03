package model

type Role struct {
	Id        uint   `gorm:"primary_key"`
	Nick      string `gorm:"not null;index:unique;unique"`
	Name      string `gorm:"not null;index:unique;unique"`
	CreatedAt int64  `gorm:"not null;autoCreateTime"`
	UpdatedAt int64  `gorm:"not null;autoUpdateTime"`
	DeletedAt uint   `gorm:"not null;default:0"`
}

type Permission struct {
	Id        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null;index:unique;unique"`
	Path      string `gorm:"not null;index:unique;unique"`
	CreatedAt int64  `gorm:"not null;autoCreateTime"`
	UpdatedAt int64  `gorm:"not null;autoUpdateTime"`
	DeletedAt uint   `gorm:"not null;default:0"`
}

type UserRole struct {
	Id        uint  `gorm:"primary_key"`
	Uid       uint  `gorm:"not null;default:0"`
	Rid       uint  `gorm:"not null;default:0"`
	CreatedAt int64 `gorm:"not null;autoCreateTime"`
	UpdatedAt int64 `gorm:"not null;autoUpdateTime"`
	DeletedAt uint  `gorm:"not null;default:0"`
}

type RolePermission struct {
	Id        uint  `gorm:"primary_key"`
	Rid       uint  `gorm:"not null;default:0"`
	Pid       uint  `gorm:"not null;default:0"`
	CreatedAt int64 `gorm:"not null;autoCreateTime"`
	UpdatedAt int64 `gorm:"not null;autoUpdateTime"`
	DeletedAt uint  `gorm:"not null;default:0"`
}
