package response

import (
	"basic_framework/web/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseResponse struct {
	Code int `json:"code"`
}

type Response struct {
	BaseResponse
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func result(code int, msg string, data interface{}, c *gin.Context) {
	var b BaseResponse
	b.Code = code
	c.JSON(http.StatusOK, Response{
		BaseResponse: b,
		Msg:          msg,
		Data:         data,
	})
}

// 简单成功返回
func SuccessSimple(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, data)
}

// 简单成功返回
func SuccessSimpleNone(c *gin.Context) {
	var b BaseResponse
	b.Code = common.SUCCESS
	c.JSON(http.StatusOK, b)
}

// 带数据成功返回
func Success(data interface{}, c *gin.Context) {
	result(common.SUCCESS, common.GetCodeMsg(common.SUCCESS, c.GetString("languages")), data, c)
}

// 简单成功返回
func SuccessNone(c *gin.Context) {
	result(common.SUCCESS, common.GetCodeMsg(common.SUCCESS, c.GetString("languages")), map[string]interface{}{}, c)
}

// 带信息成功返回
func SuccessMessage(msg string, c *gin.Context) {
	result(common.SUCCESS, msg, map[string]interface{}{}, c)
}

// 详细成功返回
func SuccessDetailed(data interface{}, msg string, c *gin.Context) {
	result(common.SUCCESS, msg, data, c)
}

// 带数据错误返回
func Error(data interface{}, c *gin.Context) {
	result(common.ERROR, common.GetCodeMsg(common.ERROR, c.GetString("languages")), data, c)
}

// 简单错误返回
func ErrorNone(c *gin.Context) {
	result(common.ERROR, common.GetCodeMsg(common.ERROR, c.GetString("languages")), map[string]interface{}{}, c)
}

// 错误码错误返回
func ErrorCode(code int, c *gin.Context) {
	result(code, common.GetCodeMsg(code, c.GetString("languages")), map[string]interface{}{}, c)
}

// 信息错误返回
func ErrorMessage(msg string, c *gin.Context) {
	result(common.ERROR, msg, map[string]interface{}{}, c)
}

// 信息错误返回
func ErrorDetailed(code int, msg string, c *gin.Context) {
	result(code, common.GetCodeMsg(code, c.GetString("languages"))+msg, map[string]interface{}{}, c)
}
