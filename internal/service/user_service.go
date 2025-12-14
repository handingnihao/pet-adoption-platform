package service

import (
	"context"
	"errors"
	"fmt"
	"pet-adoption-platform/internal/dao"
	"pet-adoption-platform/internal/model"
	"pet-adoption-platform/pkg/logger"
	"pet-adoption-platform/pkg/utils"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ErrUserExists         = errors.New("用户已存在")
	ErrUserNotFound       = errors.New("用户不存在")
	ErrInvalidPassword    = errors.New("密码错误")
	ErrUserDisabled       = errors.New("用户已被禁用")
	ErrInvalidCode        = errors.New("验证码错误或已过期")
	ErrPhoneExists        = errors.New("手机号已被注册")
	ErrEmailExists        = errors.New("邮箱已被注册")
	ErrOldPasswordWrong   = errors.New("原密码错误")
)

// UserService 用户服务
type UserService struct {
	userDAO *dao.UserDAO
	cache   *redis.Client
}

// NewUserService 创建用户服务实例
func NewUserService(db *gorm.DB, rdb *redis.Client) *UserService {
	return &UserService{
		userDAO: dao.NewUserDAO(db),
		cache:   rdb,
	}
}

// Register 用户注册
func (s *UserService) Register(req *model.UserRegisterRequest) (*model.User, error) {
	ctx := context.Background()

	// 2. 检查用户名是否存在
	exists, err := s.userDAO.ExistsByUsername(req.Username)
	if err != nil {
		logger.Error("检查用户名失败", zap.Error(err))
		return nil, err
	}
	if exists {
		return nil, ErrUserExists
	}

	// 3. 检查手机号是否存在
	exists, err = s.userDAO.ExistsByPhone(req.Phone)
	if err != nil {
		logger.Error("检查手机号失败", zap.Error(err))
		return nil, err
	}
	if exists {
		return nil, ErrPhoneExists
	}

	// 4. 检查邮箱是否存在
	exists, err = s.userDAO.ExistsByEmail(req.Email)
	if err != nil {
		logger.Error("检查邮箱失败", zap.Error(err))
		return nil, err
	}
	if exists {
		return nil, ErrEmailExists
	}

	// 5. 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		logger.Error("密码加密失败", zap.Error(err))
		return nil, err
	}

	// 6. 创建用户
	user := &model.User{
		Username: req.Username,
		Password: hashedPassword,
		Phone:    req.Phone,
		Email:    req.Email,
		Role:     model.RoleUser,
		Status:   model.UserStatusNormal,
	}

	if err := s.userDAO.Create(user); err != nil {
		logger.Error("创建用户失败", zap.Error(err))
		return nil, err
	}


	logger.Info("用户注册成功",
		zap.Int64("user_id", user.ID),
		zap.String("username", user.Username))

	return user, nil
}

// Login 用户登录
func (s *UserService) Login(req *model.UserLoginRequest, ip string) (*model.UserLoginResponse, error) {
	// 1. 查询用户
	var user *model.User
	var err error

	// 支持用户名、手机号、邮箱登录
	user, err = s.userDAO.GetByUsername(req.Username)
	if err != nil {
		logger.Error("查询用户失败", zap.Error(err))
		return nil, err
	}

	if user == nil {
		user, err = s.userDAO.GetByPhone(req.Username)
		if err != nil {
			return nil, err
		}
	}

	if user == nil {
		user, err = s.userDAO.GetByEmail(req.Username)
		if err != nil {
			return nil, err
		}
	}

	if user == nil {
		logger.Warn("用户不存在", zap.String("username", req.Username))
		return nil, ErrUserNotFound
	}

	// 2. 检查用户状态
	if !user.IsActive() {
		logger.Warn("用户已被禁用",
			zap.Int64("user_id", user.ID),
			zap.String("username", user.Username))
		return nil, ErrUserDisabled
	}

	// 3. 验证密码
	if !utils.CheckPassword(req.Password, user.Password) {
		logger.Warn("密码错误",
			zap.Int64("user_id", user.ID),
			zap.String("username", user.Username))
		return nil, ErrInvalidPassword
	}

	// 4. 生成Token
	token, expiresAt, err := utils.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		logger.Error("生成Token失败", zap.Error(err))
		return nil, err
	}

	// 5. 更新最后登录时间
	_ = s.userDAO.UpdateLastLogin(user.ID, ip)

	// 6. 缓存用户信息（可选，用于快速访问）
	userCacheKey := fmt.Sprintf("user:%d", user.ID)
	ctx := context.Background()
	_ = s.cache.Set(ctx, userCacheKey, user.ToUserInfo(), 24*time.Hour).Err()

	logger.Info("用户登录成功",
		zap.Int64("user_id", user.ID),
		zap.String("username", user.Username),
		zap.String("ip", ip))

	return &model.UserLoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User:      user.ToUserInfo(),
	}, nil
}

// GetProfile 获取用户信息
func (s *UserService) GetProfile(userID int64) (*model.UserInfo, error) {
	// 先从缓存获取（简化版，实际需要反序列化）
	cacheKey := fmt.Sprintf("user:%d", userID)
	ctx := context.Background()
	_, err := s.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		// 缓存命中，但需要反序列化，这里简化处理，直接查数据库
	}

	// 缓存未命中，从数据库获取
	user, err := s.userDAO.GetByID(userID)
	if err != nil {
		logger.Error("查询用户失败", zap.Error(err))
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	userInfo := user.ToUserInfo()

	// 更新缓存
	_ = s.cache.Set(ctx, cacheKey, userInfo, 24*time.Hour).Err()

	return userInfo, nil
}

// UpdateProfile 更新用户信息
func (s *UserService) UpdateProfile(userID int64, req *model.UserUpdateRequest) error {
	// 1. 查询用户
	user, err := s.userDAO.GetByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	// 2. 更新字段
	updates := make(map[string]interface{})
	if req.RealName != "" {
		updates["real_name"] = req.RealName
	}
	if req.Avatar != "" {
		updates["avatar"] = req.Avatar
	}
	if req.Gender > 0 {
		updates["gender"] = req.Gender
	}
	if req.Birthday != nil {
		updates["birthday"] = req.Birthday
	}
	if req.Address != "" {
		updates["address"] = req.Address
	}

	if len(updates) == 0 {
		return nil
	}

	// 3. 更新数据库
	if err := s.userDAO.UpdateFields(userID, updates); err != nil {
		logger.Error("更新用户信息失败", zap.Error(err))
		return err
	}

	// 4. 清除缓存
	cacheKey := fmt.Sprintf("user:%d", userID)
	ctx := context.Background()
	_ = s.cache.Del(ctx, cacheKey).Err()

	logger.Info("用户信息更新成功", zap.Int64("user_id", userID))

	return nil
}

// UpdatePassword 修改密码
func (s *UserService) UpdatePassword(userID int64, req *model.UpdatePasswordRequest) error {
	// 1. 查询用户
	user, err := s.userDAO.GetByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	// 2. 验证原密码
	if !utils.CheckPassword(req.OldPassword, user.Password) {
		logger.Warn("原密码错误", zap.Int64("user_id", userID))
		return ErrOldPasswordWrong
	}

	// 3. 加密新密码
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		logger.Error("密码加密失败", zap.Error(err))
		return err
	}

	// 4. 更新密码
	if err := s.userDAO.UpdatePassword(userID, hashedPassword); err != nil {
		logger.Error("更新密码失败", zap.Error(err))
		return err
	}

	logger.Info("密码修改成功", zap.Int64("user_id", userID))

	return nil
}

// GetUserList 获取用户列表（管理员功能）
func (s *UserService) GetUserList(page, pageSize int, role model.UserRole, status *model.UserStatus) ([]*model.UserInfo, int64, error) {
	users, total, err := s.userDAO.List(page, pageSize, role, status)
	if err != nil {
		return nil, 0, err
	}

	userInfos := make([]*model.UserInfo, 0, len(users))
	for _, user := range users {
		userInfos = append(userInfos, user.ToUserInfo())
	}

	return userInfos, total, nil
}

// SearchUsers 搜索用户（管理员功能）
func (s *UserService) SearchUsers(keyword string, page, pageSize int) ([]*model.UserInfo, int64, error) {
	users, total, err := s.userDAO.Search(keyword, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	userInfos := make([]*model.UserInfo, 0, len(users))
	for _, user := range users {
		userInfos = append(userInfos, user.ToUserInfo())
	}

	return userInfos, total, nil
}

// DisableUser 禁用用户（管理员功能）
func (s *UserService) DisableUser(userID int64) error {
	if err := s.userDAO.UpdateStatus(userID, model.UserStatusDisabled); err != nil {
		return err
	}

	// 清除缓存
	cacheKey := fmt.Sprintf("user:%d", userID)
	ctx := context.Background()
	_ = s.cache.Del(ctx, cacheKey).Err()

	logger.Info("用户已禁用", zap.Int64("user_id", userID))
	return nil
}

// EnableUser 启用用户（管理员功能）
func (s *UserService) EnableUser(userID int64) error {
	if err := s.userDAO.UpdateStatus(userID, model.UserStatusNormal); err != nil {
		return err
	}

	// 清除缓存
	cacheKey := fmt.Sprintf("user:%d", userID)
	ctx := context.Background()
	_ = s.cache.Del(ctx, cacheKey).Err()

	logger.Info("用户已启用", zap.Int64("user_id", userID))
	return nil
}
