package routers

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	SetupUserRoutes(r)
	SetupPostRoutes(r)
	SetupCommentRoutes(r)
}
