package repository

import (
	"user/internal/data/dao"
	"user/internal/data/model"
)

type AdminRepository interface {
	// 添加获取用户列表的方法
	GetUserList(page, pageSize int, keyword string, status int) ([]*model.User, int64, error)
	// 包含全部 User 相关的方法
	UserRepository // 嵌入 UserRepository 接口
}

// adminRepository 管理员仓储实现
type adminRepository struct {
	UserRepository
	userDao  *dao.UserDao // 直接访问 DAO 层
	adminDao *dao.AdminDao
}

// NewAdminRepository 创建管理员仓储实例
func NewAdminRepository() AdminRepository {
	userDao := dao.NewUserDao()
	adminDao := dao.NewAdminDao()

	return &adminRepository{
		userDao:  userDao,
		adminDao: adminDao,
	}
}

// GetUserList 获取用户列表
func (r *adminRepository) GetUserList(page, pageSize int, keyword string, status int) ([]*model.User, int64, error) {
	// 直接使用 DAO 层实现列表功能
	return r.adminDao.GetUserList(page, pageSize, keyword, status)
}
