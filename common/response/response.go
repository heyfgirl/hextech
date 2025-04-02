package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	// 业务成功
	SUCCESS_CODE = iota
	// 业务错误
	ERROR_CODE
)

// 返回结果 业务逻辑
func Result(code int, data any, msg string, c *gin.Context) {
	// 开始时间
	c.AbortWithStatusJSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS_CODE, map[string]any{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS_CODE, map[string]any{}, message, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS_CODE, data, "成功", c)
}

func OkWithDetailed(data any, message string, c *gin.Context) {
	Result(SUCCESS_CODE, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR_CODE, nil, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR_CODE, nil, message, c)
}

func FailWithDetailed(data any, message string, c *gin.Context) {
	Result(ERROR_CODE, data, message, c)
}

// 401 未验证
func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		ERROR_CODE,
		nil,
		message,
	})
	c.Abort()
}
