package dao

import (
	"errors"

	"user/internal/data/database"
	"user/internal/data/model"

	"gorm.io/gorm"
)

// UserDao 用户数据访问对象
type UserDao struct {
	db *gorm.DB
}

// NewUserDao 创建用户DAO实例
func NewUserDao() *UserDao {
	return &UserDao{
		db: database.DB,
	}
}

// Create 创建用户
func (d *UserDao) Create(user *model.User) error {
	return d.db.Create(user).Error
}

// GetByUsername 根据用户名获取用户信息
func (d *UserDao) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := d.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 返回nil表示未找到用户
		}
		return nil, err
	}
	return &user, nil
}

// GetByEmail 根据邮箱获取用户信息
func (d *UserDao) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := d.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetByID 根据用户ID获取用户信息
func (d *UserDao) GetByID(id int64) (*model.User, error) {
	var user model.User
	err := d.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Update 更新用户信息
func (d *UserDao) Update(user *model.User) error {
	return d.db.Save(user).Error
}

// UpdateFields 更新用户指定字段
func (d *UserDao) UpdateFields(id int64, fields map[string]interface{}) error {
	return d.db.Model(&model.User{}).Where("id = ?", id).Updates(fields).Error
}

// Delete 删除用户(软删除)
func (d *UserDao) Delete(id int64) error {
	return d.db.Delete(&model.User{}, id).Error
}

// ExistsByUsername 检查用户名是否存在
func (d *UserDao) ExistsByUsername(username string) (bool, error) {
	var count int64
	err := d.db.Model(&model.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

// ExistsByEmail 检查邮箱是否存在
func (d *UserDao) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := d.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

// List 获取用户列表(分页)
func (d *UserDao) List(page, pageSize int) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64

	// 获取总数
	if err := d.db.Model(&model.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (page - 1) * pageSize
	if err := d.db.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
