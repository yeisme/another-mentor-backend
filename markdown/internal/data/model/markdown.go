package model

import (
	"time"

	"gorm.io/gorm"
)

// MarkdownFile 用于存储Markdown文件的元数据
type MarkdownFile struct {
	ID         int64     `gorm:"primaryKey;autoIncrement;column:id"`
	Filename   string    `gorm:"type:varchar(255);not null;index"`
	Size       int64     `gorm:"not null"`
	UploadTime time.Time `gorm:"not null;index"`
	UpdateTime time.Time `gorm:"not null;index"`
	OwnerID    int64     `gorm:"not null;index"`
	Hash       string    `gorm:"type:varchar(64);not null;uniqueIndex"`
	IsPublic   bool      `gorm:"not null;default:false;index"`
	IsDeleted  bool      `gorm:"not null;default:false;index"`
	Tags       string    `gorm:"type:jsonb;default:'[]'"`
	Path       string    `gorm:"type:varchar(512);not null;index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
