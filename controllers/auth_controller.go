package controllers

import (
	"go-webapi-example/models"
	"go-webapi-example/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	userService *services.UserService
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		userService: services.NewUserService(db),
	}
}

// Login godoc
// @Summary 用户登录
// @Description 用户使用邮箱和密码登录系统，成功后返回JWT访问令牌
// @Tags auth
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "登录凭据"
// @Success 200 {object} models.LoginResponse "登录成功，返回访问令牌和用户信息"
// @Failure 400 {object} map[string]string "请求参数错误"
// @Failure 401 {object} map[string]string "登录凭据无效"
// @Router /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResponse, err := c.userService.Login(&req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loginResponse)
}

// Register godoc
// @Summary 用户注册
// @Description 注册新用户账户，创建成功后需要使用登录接口获取访问令牌
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "用户注册信息"
// @Success 201 {object} object{message=string,user_id=int} "注册成功，返回用户ID"
// @Failure 400 {object} map[string]string "请求参数错误或邮箱已存在"
// @Failure 500 {object} map[string]string "服务器内部错误"
// @Router /auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 普通注册只能创建普通用户
	req.Role = "user"

	user, err := c.userService.CreateUser(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user_id": user.ID,
	})
}
