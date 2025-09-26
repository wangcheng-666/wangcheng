package routers

import (
	"web3-blog-backend/controllers"
	"web3-blog-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(r *gin.Engine) {
	// 所有 /posts 开头的接口都需要认证
	auth := r.Group("/posts")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("", controllers.CreatePost)
		auth.PUT("/:id", controllers.UpdatePost)
		auth.DELETE("/:id", controllers.DeletePost)
	}

	// 公开接口
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
}
