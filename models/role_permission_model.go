package models

// RoleModel - Role
type RoleModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// RolePermissionModel - Connection Role and Permission Model
type RolePermissionModel struct {
	Permission PermissionModel `json:"permission"`
	Role       RoleModel       `json:"role"`
}

// PermissionModel - Permission
type PermissionModel struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	Method   string `json:"method"`
}
