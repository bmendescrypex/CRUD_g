package routes	

import "github.com/gin-gonic/gin"
import "github.com/bmendescrypex/crud_go/routes"

func UserRoute(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})
}