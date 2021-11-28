package request

import (
	"basic_framework/web/common"
	"github.com/gin-gonic/gin"
)

// demo
type UserParam struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
	Age      *int   `form:"age" validate:"required"`
}

func (this *UserParam) GetInfo(c *gin.Context) error {
	c.ShouldBind(this)
	return common.ValidationPara(this)
}
