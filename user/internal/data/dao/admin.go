package dao

import (
	"user/internal/data/database"
	"user/internal/data/model"

	"gorm.io/gorm"
)

type AdminDao struct {
	db *gorm.DB
}

// NewAdminDao 创建管理员DAO实例
func NewAdminDao() *AdminDao {
	return &AdminDao{
		db: database.DB,
	}
}

// GetUserList 获取用户列表
func (d *AdminDao) GetUserList(page, pageSize int, keyword string, status int) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64

	query := d.db.Model(&model.User{})

	// 添加筛选条件
	if keyword != "" {
		query = query.Where("username LIKE ? OR email LIKE ? OR nickname LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	if status != 0 {
		query = query.Where("status = ?", status)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (page - 1) * pageSize
	if err := query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
