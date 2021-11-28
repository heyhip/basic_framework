package routes

import (
	"basic_framework/web/controllers"
	"github.com/gin-gonic/gin"
)

func testRoute(r *gin.Engine) {

	g := r.Group("/test")
	{
		g.GET("/index", controllers.Index)
		g.POST("/index", controllers.Index)
	}

}
