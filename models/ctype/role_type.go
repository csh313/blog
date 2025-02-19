package ctype

import "encoding/json"

// 角色类型，用于权限控制
type Role int

const (
	PermissionAdmin       Role = 1 //管理员
	PermissionUser        Role = 2 //普通用户
	PermissionVisitor     Role = 3 //游客
	PermissionDisableUser Role = 4 //封禁的用户
)

// 角色json序列化
func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// 角色匹配
func (r Role) String() string {
	switch r {
	case PermissionAdmin:
		return "管理员"
	case PermissionUser:
		return "用户"
	case PermissionVisitor:
		return "游客"
	case PermissionDisableUser:
		return "被禁言"
	default:
		return "其他"
	}
}
