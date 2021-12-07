package controllers

import (
	"basic_framework/core/log"
	"basic_framework/web/common"
	"basic_framework/web/request"
	"basic_framework/web/services"
	"github.com/gin-gonic/gin"
)

// @Tags Index
// @Summary 文档信息
// @Security ApiKeyAuth
// @accept x-www-form-urlencoded
// @Produce json
// @Param data formData request.UserParam true "请求参数"
// @Success 200 {object} response.UserResp {"code":200,"msg":"ok","data":{"username":"111","password":"222","age":333}}
// @Router /deviceKey/info [post]
func Index(c *gin.Context) {
	var req request.UserParam
	e := req.GetInfo(c)
	if e != nil {
		log.Error(e)
		common.ErrorCode(common.PARAM_ERROR, c)
		return
	}

	result := services.UserService{}.GetUser(req.Username, req.Password, *req.Age)

	common.Success(result, c)
}
