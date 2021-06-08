package YOUR_FUNCTION_NAME

import (
	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {

	r := gin.Default()
	r.Any("/*route", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"METHOD": c.Request.Method,
			"PATH": c.Param("route"),
		})
	})

	return r
}