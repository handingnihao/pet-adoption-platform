package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构体
type Response struct {
	Code      int         `json:"code"`      // 业务状态码
	Message   string      `json:"message"`   // 提示信息
	Data      interface{} `json:"data"`      // 数据
	Timestamp int64       `json:"timestamp"` // 时间戳
}

// 业务状态码
const (
	CodeSuccess      = 200   // 成功
	CodeError        = 500   // 服务器错误
	CodeInvalidParam = 400   // 参数错误
	CodeUnauthorized = 401   // 未授权
	CodeForbidden    = 403   // 禁止访问
	CodeNotFound     = 404   // 资源不存在
	CodeConflict     = 409   // 资源冲突
	CodeTooMany      = 429   // 请求过多
)

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:      CodeSuccess,
		Message:   "success",
		Data:      data,
		Timestamp: time.Now().Unix(),
	})
}

// SuccessWithMessage 成功响应（自定义消息）
func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:      CodeSuccess,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Unix(),
	})
}

// Error 错误响应
func Error(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{
		Code:      CodeError,
		Message:   message,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	})
}

// ErrorWithCode 错误响应（自定义状态码）
func ErrorWithCode(c *gin.Context, code int, message string) {
	httpStatus := http.StatusOK
	if code == CodeUnauthorized {
		httpStatus = http.StatusUnauthorized
	} else if code == CodeForbidden {
		httpStatus = http.StatusForbidden
	} else if code == CodeNotFound {
		httpStatus = http.StatusNotFound
	}

	c.JSON(httpStatus, Response{
		Code:      code,
		Message:   message,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	})
}

// InvalidParam 参数错误
func InvalidParam(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Response{
		Code:      CodeInvalidParam,
		Message:   message,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	})
}

// ParamError 参数错误（InvalidParam的别名）
func ParamError(c *gin.Context, message string) {
	InvalidParam(c, message)
}

// BadRequest 错误请求
func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Response{
		Code:      CodeInvalidParam,
		Message:   message,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	})
}

// Unauthorized 未授权
func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, Response{
		Code:      CodeUnauthorized,
		Message:   message,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	})
}

// Forbidden 禁止访问
func Forbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, Response{
		Code:      CodeForbidden,
		Message:   message,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	})
}

// NotFound 资源不存在
func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Code:      CodeNotFound,
		Message:   message,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	})
}

// TooManyRequests 请求过多
func TooManyRequests(c *gin.Context, message string) {
	c.JSON(http.StatusTooManyRequests, Response{
		Code:      CodeTooMany,
		Message:   message,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	})
}

// PageData 分页数据结构
type PageData struct {
	List     interface{} `json:"list"`      // 列表数据
	Total    int64       `json:"total"`     // 总数
	Page     int         `json:"page"`      // 当前页
	PageSize int         `json:"page_size"` // 每页数量
}

// SuccessWithPage 分页成功响应
func SuccessWithPage(c *gin.Context, list interface{}, total int64, page, pageSize int) {
	Success(c, PageData{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}

// SuccessWithPagination 分页成功响应（SuccessWithPage的别名）
func SuccessWithPagination(c *gin.Context, list interface{}, page, pageSize int, total int64) {
	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}
	
	Success(c, gin.H{
		"list": list,
		"pagination": gin.H{
			"page":        page,
			"page_size":   pageSize,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}
