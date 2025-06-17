package routes

import (
	"webapi/controllers"
	"webapi/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// 初始化控制器
	userController := controllers.NewUserController(db)
	productController := controllers.NewProductController(db)
	authController := controllers.NewAuthController(db)

	// 认证路由（不需要JWT）
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
	}

	// API 版本 v1
	v1 := r.Group("/api/v1")
	{
		// 公开的用户路由（需要管理员权限）
		users := v1.Group("/users")
		users.Use(middleware.AuthMiddleware())
		{
			users.GET("/profile", userController.GetProfile)
			users.GET("", middleware.AdminMiddleware(), userController.GetUsers)
			users.GET("/:id", middleware.AdminMiddleware(), userController.GetUser)
			users.PUT("/:id", userController.UpdateUser) // 用户可以更新自己的信息
			users.DELETE("/:id", middleware.AdminMiddleware(), userController.DeleteUser)
		}

		// 产品路由（需要认证）
		products := v1.Group("/products")
		products.Use(middleware.AuthMiddleware())
		{
			products.POST("", productController.CreateProduct)
			products.GET("", productController.GetProducts)
			products.GET("/:id", productController.GetProduct)
			products.PUT("/:id", productController.UpdateProduct)
			products.DELETE("/:id", productController.DeleteProduct)
		}

		// 管理员路由
		admin := v1.Group("/admin")
		admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		{
			admin.POST("/users", userController.CreateUser) // 管理员创建用户
		}
	}

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "API is running",
		})
	})
}
