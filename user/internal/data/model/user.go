package model

import (
	"time"

	"gorm.io/gorm"
)

// 用户角色
const (
	RoleUser string = "user" // 普通用户
)

// 用户状态
const (
	StatusActive   = 1 // 正常
	StatusInactive = 2 // 未激活
	StatusDisabled = 3 // 已禁用
)

// User 用户模型
type User struct {
	ID           int64          `gorm:"primaryKey;column:id;autoIncrement"`
	Username     string         `gorm:"type:varchar(50);uniqueIndex;not null"`
	Password     string         `gorm:"type:varchar(100);not null"` // 加密后的密码
	Email        string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	Phone        string         `gorm:"type:varchar(20);default:null"`
	Avatar       string         `gorm:"type:varchar(255);default:null"`
	Nickname     string         `gorm:"type:varchar(50);default:null"`
	Introduction string         `gorm:"type:text;default:null"`
	Role         string         `gorm:"type:varchar(10);default:'user';not null"` // user: 普通用户 admin: 管理员
	Status       int            `gorm:"type:smallint;default:1;not null"`         // 1: 正常, 2: 未激活, 3: 已禁用
	LastLoginAt  *time.Time     `gorm:"default:null"`                             // 最后登录时间
	CreatedAt    time.Time      `gorm:"not null"`                                 // 创建时间
	UpdatedAt    time.Time      `gorm:"not null"`                                 // 更新时间
	DeletedAt    gorm.DeletedAt `gorm:"index"`                                    // 软删除时间
}
