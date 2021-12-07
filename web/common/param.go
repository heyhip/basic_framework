package common

import "github.com/gin-gonic/gin"

type Param interface {
	GetInfo(c *gin.Context) error
}
