// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUserPermission = "user_permissions"

// UserPermission 关联用户和一般权限
type UserPermission struct {
	UserID       int32 `gorm:"column:user_id;type:int;primaryKey;comment:用户ID" json:"user_id"`                                            // 用户ID
	PermissionID int32 `gorm:"column:permission_id;type:int;primaryKey;index:permission_id,priority:1;comment:权限ID" json:"permission_id"` // 权限ID
}

// TableName UserPermission's table name
func (*UserPermission) TableName() string {
	return TableNameUserPermission
}