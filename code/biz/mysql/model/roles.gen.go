// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRole = "roles"

// Role 存储权限角色类别
type Role struct {
	ID   int32  `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:角色ID，主键" json:"id"` // 角色ID，主键
	Name string `gorm:"column:name;type:varchar(50);not null;comment:角色名称" json:"name"`             // 角色名称
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}