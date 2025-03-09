package repository

import (
	"errors"
	"time"

	"user/internal/data/dao"
	"user/internal/data/model"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

// 定义错误
var (
	ErrUserNotFound     = errors.New("用户不存在")
	ErrUsernameExists   = errors.New("用户名已存在")
	ErrEmailExists      = errors.New("邮箱已存在")
	ErrPasswordMismatch = errors.New("密码错误")
)

// UserRepository 用户仓储接口
type UserRepository interface {
	Register(username, password, email string) (*model.User, error)
	Login(usernameOrEmail, password string) (*model.User, error)
	GetUserByID(id int64) (*model.User, error)
	UpdateProfile(id int64, nickname, avatar, introduction string) error
	UpdatePassword(id int64, oldPassword, newPassword string) error
	UpdateStatus(id int64, status int) error
}

// userRepository 用户仓储实现
type userRepository struct {
	userDao *dao.UserDao
}

// NewUserRepository 创建用户仓储实例
func NewUserRepository() UserRepository {
	return &userRepository{
		userDao: dao.NewUserDao(),
	}
}

// Register 用户注册
func (r *userRepository) Register(username, password, email string) (*model.User, error) {
	// 检查用户名是否已存在
	exists, err := r.userDao.ExistsByUsername(username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrUsernameExists
	}

	// 检查邮箱是否已存在
	exists, err = r.userDao.ExistsByEmail(email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrEmailExists
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 创建用户
	now := time.Now()

	user := &model.User{
		ID:        int64(uuid.New().ID()),
		Username:  username,
		Password:  string(hashedPassword),
		Email:     email,
		Status:    model.StatusActive,
		Role:      model.RoleUser,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := r.userDao.Create(user); err != nil {
		return nil, err
	}

	logx.Infof("用户注册成功: %s", username)

	return user, nil
}

// Login 用户登录
func (r *userRepository) Login(usernameOrEmail, password string) (*model.User, error) {
	var user *model.User
	var err error

	// 尝试通过用户名查找
	user, err = r.userDao.GetByUsername(usernameOrEmail)
	if err != nil {
		return nil, err
	}

	// 如果没找到，尝试通过邮箱查找
	if user == nil {
		user, err = r.userDao.GetByEmail(usernameOrEmail)
		if err != nil {
			return nil, err
		}
	}

	// 用户不存在
	if user == nil {
		return nil, ErrUserNotFound
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, ErrPasswordMismatch
	}

	// 更新最后登录时间
	now := time.Now()
	user.LastLoginAt = &now
	if err := r.userDao.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID 根据ID获取用户
func (r *userRepository) GetUserByID(id int64) (*model.User, error) {
	user, err := r.userDao.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

// UpdateProfile 更新用户资料
func (r *userRepository) UpdateProfile(id int64, nickname, avatar, introduction string) error {
	// 获取用户
	user, err := r.userDao.GetByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	// 更新字段
	fields := make(map[string]interface{})
	if nickname != "" {
		fields["nickname"] = nickname
	}
	if avatar != "" {
		fields["avatar"] = avatar
	}
	if introduction != "" {
		fields["introduction"] = introduction
	}
	fields["updated_at"] = time.Now()

	return r.userDao.UpdateFields(id, fields)
}

// UpdatePassword 更新密码
func (r *userRepository) UpdatePassword(id int64, oldPassword, newPassword string) error {
	// 获取用户
	user, err := r.userDao.GetByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	// 验证旧密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return ErrPasswordMismatch
	}

	// 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 更新密码
	fields := map[string]interface{}{
		"password":   string(hashedPassword),
		"updated_at": time.Now(),
	}
	return r.userDao.UpdateFields(id, fields)
}

// UpdateStatus 更新用户状态
func (r *userRepository) UpdateStatus(id int64, status int) error {
	// 获取用户
	user, err := r.userDao.GetByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	// 更新状态
	fields := map[string]interface{}{
		"status":     status,
		"updated_at": time.Now(),
	}
	return r.userDao.UpdateFields(id, fields)
}
