package dao

import (
	"errors"
	"pet-adoption-platform/internal/model"
	"time"

	"gorm.io/gorm"
)

// UserDAO 用户数据访问对象
type UserDAO struct {
	db *gorm.DB
}

// NewUserDAO 创建UserDAO实例
func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

// Create 创建用户
func (dao *UserDAO) Create(user *model.User) error {
	return dao.db.Create(user).Error
}

// GetByID 根据ID查询用户
func (dao *UserDAO) GetByID(id int64) (*model.User, error) {
	var user model.User
	err := dao.db.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetByUsername 根据用户名查询
func (dao *UserDAO) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := dao.db.Where("username = ? AND deleted_at IS NULL", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetByPhone 根据手机号查询
func (dao *UserDAO) GetByPhone(phone string) (*model.User, error) {
	var user model.User
	err := dao.db.Where("phone = ? AND deleted_at IS NULL", phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetByEmail 根据邮箱查询
func (dao *UserDAO) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := dao.db.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetByLoginIdentifier 根据用户名/手机号/邮箱查询（合并查询，优化登录性能）
func (dao *UserDAO) GetByLoginIdentifier(identifier string) (*model.User, error) {
	var user model.User
	err := dao.db.Where(
		"(username = ? OR phone = ? OR email = ?) AND deleted_at IS NULL",
		identifier, identifier, identifier,
	).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Update 更新用户信息
func (dao *UserDAO) Update(user *model.User) error {
	return dao.db.Save(user).Error
}

// UpdateFields 更新指定字段
func (dao *UserDAO) UpdateFields(id int64, fields map[string]interface{}) error {
	return dao.db.Model(&model.User{}).Where("id = ?", id).Updates(fields).Error
}

// UpdatePassword 更新密码
func (dao *UserDAO) UpdatePassword(id int64, newPassword string) error {
	return dao.db.Model(&model.User{}).
		Where("id = ?", id).
		Update("password", newPassword).Error
}

// UpdateLastLogin 更新最后登录时间和IP
func (dao *UserDAO) UpdateLastLogin(id int64, ip string) error {
	return dao.db.Model(&model.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"last_login_at": time.Now(),
			"last_login_ip": ip,
		}).Error
}

// UpdateStatus 更新用户状态
func (dao *UserDAO) UpdateStatus(id int64, status model.UserStatus) error {
	return dao.db.Model(&model.User{}).
		Where("id = ?", id).
		Update("status", status).Error
}

// Delete 软删除用户
func (dao *UserDAO) Delete(id int64) error {
	now := time.Now()
	return dao.db.Model(&model.User{}).
		Where("id = ?", id).
		Update("deleted_at", now).Error
}

// List 分页查询用户列表
func (dao *UserDAO) List(page, pageSize int, role model.UserRole, status *model.UserStatus) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64

	query := dao.db.Model(&model.User{}).Where("deleted_at IS NULL")

	// 角色过滤
	if role != "" {
		query = query.Where("role = ?", role)
	}

	// 状态过滤
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	// 查询总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// Search 搜索用户（根据用户名、手机号、邮箱）
func (dao *UserDAO) Search(keyword string, page, pageSize int) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64

	query := dao.db.Model(&model.User{}).Where("deleted_at IS NULL")

	// 模糊搜索
	if keyword != "" {
		query = query.Where(
			"username LIKE ? OR phone LIKE ? OR email LIKE ? OR real_name LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%",
		)
	}

	// 查询总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// ExistsByUsername 检查用户名是否存在
func (dao *UserDAO) ExistsByUsername(username string) (bool, error) {
	var count int64
	err := dao.db.Model(&model.User{}).
		Where("username = ? AND deleted_at IS NULL", username).
		Count(&count).Error
	return count > 0, err
}

// ExistsByPhone 检查手机号是否存在
func (dao *UserDAO) ExistsByPhone(phone string) (bool, error) {
	var count int64
	err := dao.db.Model(&model.User{}).
		Where("phone = ? AND deleted_at IS NULL", phone).
		Count(&count).Error
	return count > 0, err
}

// ExistsByEmail 检查邮箱是否存在
func (dao *UserDAO) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := dao.db.Model(&model.User{}).
		Where("email = ? AND deleted_at IS NULL", email).
		Count(&count).Error
	return count > 0, err
}

// CountByRole 统计指定角色的用户数
func (dao *UserDAO) CountByRole(role model.UserRole) (int64, error) {
	var count int64
	err := dao.db.Model(&model.User{}).
		Where("role = ? AND deleted_at IS NULL", role).
		Count(&count).Error
	return count, err
}

// CountByStatus 统计指定状态的用户数
func (dao *UserDAO) CountByStatus(status model.UserStatus) (int64, error) {
	var count int64
	err := dao.db.Model(&model.User{}).
		Where("status = ? AND deleted_at IS NULL", status).
		Count(&count).Error
	return count, err
}

// GetRecentUsers 获取最近注册的用户
func (dao *UserDAO) GetRecentUsers(limit int) ([]*model.User, error) {
	var users []*model.User
	err := dao.db.Where("deleted_at IS NULL").
		Order("created_at DESC").
		Limit(limit).
		Find(&users).Error
	return users, err
}
