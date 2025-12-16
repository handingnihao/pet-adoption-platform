package controller

import (
	"pet-adoption-platform/internal/service"
	"pet-adoption-platform/pkg/response"

	"github.com/gin-gonic/gin"
)

// UploadController 文件上传控制器
type UploadController struct {
	uploadService *service.UploadService
}

// NewUploadController 创建上传控制器实例
func NewUploadController() *UploadController {
	return &UploadController{
		uploadService: service.NewUploadService(),
	}
}

// UploadImage 上传单张图片
// @Summary 上传图片
// @Description 上传单张图片，支持jpg/jpeg/png/gif格式
// @Tags 文件上传
// @Accept multipart/form-data
// @Produce json
// @Param file formance file true "图片文件"
// @Param category query string false "分类目录" default(general)
// @Success 200 {object} response.Response{data=service.UploadResult}
// @Router /api/v1/files/image [post]
func (c *UploadController) UploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.ParamError(ctx, "请选择要上传的文件")
		return
	}

	// 获取分类，默认为general
	category := ctx.DefaultQuery("category", "general")
	
	// 验证分类名称，只允许字母、数字和下划线
	if !isValidCategory(category) {
		response.ParamError(ctx, "无效的分类名称")
		return
	}

	result, err := c.uploadService.UploadImage(file, category)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, result)
}

// UploadImages 批量上传图片
// @Summary 批量上传图片
// @Description 批量上传多张图片，最多10张
// @Tags 文件上传
// @Accept multipart/form-data
// @Produce json
// @Param files formData file true "图片文件（可多选）"
// @Param category query string false "分类目录" default(general)
// @Success 200 {object} response.Response{data=[]service.UploadResult}
// @Router /api/v1/files/images [post]
func (c *UploadController) UploadImages(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		response.ParamError(ctx, "请选择要上传的文件")
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		response.ParamError(ctx, "请选择要上传的文件")
		return
	}

	// 限制最多上传10张图片
	if len(files) > 10 {
		response.ParamError(ctx, "一次最多上传10张图片")
		return
	}

	category := ctx.DefaultQuery("category", "general")
	if !isValidCategory(category) {
		response.ParamError(ctx, "无效的分类名称")
		return
	}

	results, err := c.uploadService.UploadMultipleImages(files, category)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, gin.H{
		"count": len(results),
		"files": results,
	})
}

// UploadCredential 上传机构认证图片
// @Summary 上传机构认证图片
// @Description 上传机构认证相关图片（营业执照、资质证书等）
// @Tags 文件上传
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "认证图片"
// @Success 200 {object} response.Response{data=service.UploadResult}
// @Router /api/v1/files/credential [post]
func (c *UploadController) UploadCredential(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.ParamError(ctx, "请选择要上传的认证图片")
		return
	}

	result, err := c.uploadService.UploadImage(file, "credentials")
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, result)
}

// UploadAvatar 上传用户头像
// @Summary 上传用户头像
// @Description 上传用户头像图片
// @Tags 文件上传
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "头像图片"
// @Success 200 {object} response.Response{data=service.UploadResult}
// @Router /api/v1/files/avatar [post]
func (c *UploadController) UploadAvatar(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.ParamError(ctx, "请选择要上传的头像")
		return
	}

	result, err := c.uploadService.UploadImage(file, "avatars")
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, result)
}

// UploadPetPhoto 上传宠物图片
// @Summary 上传宠物图片
// @Description 上传宠物相关图片
// @Tags 文件上传
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "宠物图片"
// @Success 200 {object} response.Response{data=service.UploadResult}
// @Router /api/v1/files/pet [post]
func (c *UploadController) UploadPetPhoto(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.ParamError(ctx, "请选择要上传的宠物图片")
		return
	}

	result, err := c.uploadService.UploadImage(file, "pets")
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, result)
}

// DeleteFile 删除文件
// @Summary 删除文件
// @Description 删除已上传的文件
// @Tags 文件上传
// @Accept json
// @Produce json
// @Param url body string true "文件URL"
// @Success 200 {object} response.Response
// @Router /api/v1/files [delete]
func (c *UploadController) DeleteFile(ctx *gin.Context) {
	var req struct {
		URL string `json:"url" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ParamError(ctx, "请提供文件URL")
		return
	}

	if err := c.uploadService.DeleteFile(req.URL); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, gin.H{"message": "文件删除成功"})
}

// isValidCategory 验证分类名称是否有效
func isValidCategory(category string) bool {
	if len(category) == 0 || len(category) > 50 {
		return false
	}
	for _, c := range category {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_' || c == '-') {
			return false
		}
	}
	return true
}
