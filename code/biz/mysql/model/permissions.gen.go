// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePermission = "permissions"

// Permission 存储一般的权限
type Permission struct {
	ID   int32  `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:权限ID，主键" json:"id"` // 权限ID，主键
	Name string `gorm:"column:name;type:varchar(100);not null;comment:权限名称" json:"name"`            // 权限名称
}

// TableName Permission's table name
func (*Permission) TableName() string {
	return TableNamePermission
}