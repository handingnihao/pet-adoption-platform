package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"pet-adoption-platform/config"
	"pet-adoption-platform/pkg/utils"
)

// UploadService 文件上传服务
type UploadService struct {
	config *config.UploadConfig
}

// NewUploadService 创建上传服务实例
func NewUploadService() *UploadService {
	return &UploadService{
		config: &config.AppConfig.Upload,
	}
}

// UploadResult 上传结果
type UploadResult struct {
	Filename    string `json:"filename"`
	OriginalName string `json:"original_name"`
	URL         string `json:"url"`
	Size        int64  `json:"size"`
	MimeType    string `json:"mime_type"`
}

// UploadImage 上传图片
func (s *UploadService) UploadImage(file *multipart.FileHeader, category string) (*UploadResult, error) {
	// 验证文件大小
	maxSize := s.config.MaxSize * 1024 * 1024 // 转换为字节
	if file.Size > maxSize {
		return nil, fmt.Errorf("文件大小超过限制，最大允许 %dMB", s.config.MaxSize)
	}

	// 获取文件扩展名
	ext := strings.ToLower(filepath.Ext(file.Filename))
	
	// 验证文件类型
	if !s.isAllowedExt(ext) {
		return nil, fmt.Errorf("不支持的文件类型: %s，允许的类型: %v", ext, s.config.AllowedExts)
	}

	// 生成唯一文件名
	newFilename := s.generateFilename(ext)

	// 构建存储路径: uploads/{category}/{year-month}/{filename}
	datePath := time.Now().Format("2006-01")
	relativePath := filepath.Join(category, datePath)
	fullDir := filepath.Join(s.config.LocalPath, relativePath)

	// 确保目录存在
	if err := os.MkdirAll(fullDir, 0755); err != nil {
		return nil, fmt.Errorf("创建目录失败: %w", err)
	}

	// 保存文件
	filePath := filepath.Join(fullDir, newFilename)
	if err := s.saveFile(file, filePath); err != nil {
		return nil, fmt.Errorf("保存文件失败: %w", err)
	}

	// 构建访问URL
	url := fmt.Sprintf("%s/%s/%s/%s", 
		strings.TrimSuffix(s.config.ServerURL, "/"),
		category,
		datePath,
		newFilename,
	)

	return &UploadResult{
		Filename:     newFilename,
		OriginalName: file.Filename,
		URL:          url,
		Size:         file.Size,
		MimeType:     file.Header.Get("Content-Type"),
	}, nil
}

// UploadMultipleImages 批量上传图片
func (s *UploadService) UploadMultipleImages(files []*multipart.FileHeader, category string) ([]*UploadResult, error) {
	results := make([]*UploadResult, 0, len(files))
	
	for _, file := range files {
		result, err := s.UploadImage(file, category)
		if err != nil {
			return results, fmt.Errorf("上传文件 %s 失败: %w", file.Filename, err)
		}
		results = append(results, result)
	}
	
	return results, nil
}

// DeleteFile 删除文件
func (s *UploadService) DeleteFile(url string) error {
	// 从URL提取相对路径
	serverURL := strings.TrimSuffix(s.config.ServerURL, "/")
	if !strings.HasPrefix(url, serverURL) {
		return fmt.Errorf("无效的文件URL")
	}
	
	relativePath := strings.TrimPrefix(url, serverURL)
	filePath := filepath.Join(s.config.LocalPath, relativePath)
	
	// 删除文件
	if err := os.Remove(filePath); err != nil {
		if os.IsNotExist(err) {
			return nil // 文件不存在，视为成功
		}
		return fmt.Errorf("删除文件失败: %w", err)
	}
	
	return nil
}

// isAllowedExt 检查是否是允许的扩展名
func (s *UploadService) isAllowedExt(ext string) bool {
	for _, allowed := range s.config.AllowedExts {
		if strings.EqualFold(ext, allowed) {
			return true
		}
	}
	return false
}

// generateFilename 生成唯一文件名
func (s *UploadService) generateFilename(ext string) string {
	// 使用时间戳+随机字符串生成唯一文件名
	timestamp := time.Now().UnixNano()
	randomStr := utils.RandomString(8)
	return fmt.Sprintf("%d_%s%s", timestamp, randomStr, ext)
}

// saveFile 保存文件到指定路径
func (s *UploadService) saveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
