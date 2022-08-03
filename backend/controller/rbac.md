# rbac权限设计
## 用户
id、nickname、deleted_at
## 角色
id name deleted_at
一个用户对多个角色
## 权限
id name path deleted_at
一个角色对应多个权限
## 用户角色
id uid rid
## 角色权限
id rid pid

## 如何找找到一个用户角色
## 如何找找到一个用户权限
```go
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

```

### 权限管理
1. 角色列表
2. 新增角色：选择对应权限列表加入到rp表
3. 删除角色

4. 权限列表
5. 新增权限
6. 删除权限

7. 新增用户：

