// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "users"

// User 存储用户基本信息
type User struct {
	ID             int32      `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:用户ID，主键" json:"id"`                                   // 用户ID，主键
	Username       string     `gorm:"column:username;type:varchar(50);not null;uniqueIndex:username_unique,priority:1;comment:用户名" json:"username"` // 用户名
	Nickname       string     `gorm:"column:nickname;type:varchar(50);not null;comment:用户昵称" json:"nickname"`                                       // 用户昵称
	PasswordHash   string     `gorm:"column:password_hash;type:varchar(255);not null;comment:加密后的密码" json:"password_hash"`                          // 加密后的密码
	RoleID         *int32     `gorm:"column:role_id;type:int;index:role_id,priority:1;comment:关联角色表的外键" json:"role_id"`                             // 关联角色表的外键
	AdditionalInfo *string    `gorm:"column:additional_info;type:json;comment:存储JSON格式的额外信息" json:"additional_info"`                                // 存储JSON格式的额外信息
	CreatedAt      *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:记录创建时间" json:"created_at"`                  // 记录创建时间
	UpdatedAt      *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:记录更新时间" json:"updated_at"`                  // 记录更新时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
