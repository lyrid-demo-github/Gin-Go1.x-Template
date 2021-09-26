package entry

import (
	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {

	r := gin.Default()
	r.Static("/", "./dist")

	return r
}
