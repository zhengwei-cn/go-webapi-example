package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint           `gorm:"primarykey" json:"id" example:"1"`                                                      // 用户ID
	CreatedAt time.Time      `json:"created_at" example:"2023-01-01T00:00:00Z"`                                             // 创建时间
	UpdatedAt time.Time      `json:"updated_at" example:"2023-01-01T00:00:00Z"`                                             // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                                                                        // 删除时间（软删除）
	Name      string         `gorm:"not null" json:"name" binding:"required" example:"张三"`                                  // 用户姓名
	Email     string         `gorm:"uniqueIndex;not null" json:"email" binding:"required,email" example:"user@example.com"` // 邮箱地址
	Password  string         `gorm:"not null" json:"-"`                                                                     // 密码（不在JSON中显示）
	Age       int            `json:"age" binding:"min=0" example:"25"`                                                      // 年龄
	Role      string         `gorm:"default:'user'" json:"role" example:"user"`                                             // 用户角色（user/admin/superadmin）
	IsActive  bool           `gorm:"default:true" json:"is_active" example:"true"`                                          // 是否激活
}

// Product 产品模型
type Product struct {
	ID          uint           `gorm:"primarykey" json:"id" example:"1"`                                // 产品ID
	CreatedAt   time.Time      `json:"created_at" example:"2023-01-01T00:00:00Z"`                       // 创建时间
	UpdatedAt   time.Time      `json:"updated_at" example:"2023-01-01T00:00:00Z"`                       // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`                                                  // 删除时间（软删除）
	Name        string         `gorm:"not null" json:"name" binding:"required" example:"iPhone 15"`     // 产品名称
	Description string         `json:"description" example:"最新款智能手机"`                                   // 产品描述
	Price       float64        `gorm:"not null" json:"price" binding:"required,min=0" example:"999.99"` // 产品价格
	Stock       int            `gorm:"default:0" json:"stock" binding:"min=0" example:"100"`            // 库存数量
	UserID      uint           `json:"user_id" example:"1"`                                             // 创建用户ID
	User        User           `json:"user,omitempty"`                                                  // 关联用户信息
}

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required" example:"张三"`                      // 用户姓名
	Email    string `json:"email" binding:"required,email" example:"user@example.com"` // 邮箱地址
	Password string `json:"password" binding:"required,min=6" example:"password123"`   // 密码（至少6位）
	Age      int    `json:"age" binding:"min=0" example:"25"`                          // 年龄
	Role     string `json:"role,omitempty" example:"user"`                             // 用户角色（可选）
}

// LoginRequest 登录请求
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"admin@example.com"` // 邮箱地址
	Password string `json:"password" binding:"required" example:"admin123456"`          // 密码
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string      `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."` // JWT访问令牌
	User  UserProfile `json:"user"`                                                    // 用户信息
}

// UserProfile 用户资料（不包含敏感信息）
type UserProfile struct {
	ID       uint   `json:"id" example:"1"`                   // 用户ID
	Name     string `json:"name" example:"张三"`                // 用户姓名
	Email    string `json:"email" example:"user@example.com"` // 邮箱地址
	Age      int    `json:"age" example:"25"`                 // 年龄
	Role     string `json:"role" example:"user"`              // 用户角色
	IsActive bool   `json:"is_active" example:"true"`         // 是否激活
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Name string `json:"name,omitempty" example:"李四"`                // 用户姓名（可选）
	Age  int    `json:"age,omitempty" binding:"min=0" example:"30"` // 年龄（可选）
}

// CreateProductRequest 创建产品请求
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required" example:"iPhone 15"`     // 产品名称
	Description string  `json:"description" example:"最新款智能手机"`                   // 产品描述
	Price       float64 `json:"price" binding:"required,min=0" example:"999.99"` // 产品价格
	Stock       int     `json:"stock" binding:"min=0" example:"100"`             // 库存数量
	UserID      uint    `json:"user_id" binding:"required" example:"1"`          // 创建用户ID
}

// UpdateProductRequest 更新产品请求
type UpdateProductRequest struct {
	Name        string  `json:"name,omitempty" example:"iPhone 15 Pro"`            // 产品名称（可选）
	Description string  `json:"description,omitempty" example:"升级版智能手机"`           // 产品描述（可选）
	Price       float64 `json:"price,omitempty" binding:"min=0" example:"1299.99"` // 产品价格（可选）
	Stock       int     `json:"stock,omitempty" binding:"min=0" example:"50"`      // 库存数量（可选）
}

// ErrorResponse 错误响应模型
type ErrorResponse struct {
	Error string `json:"error" example:"错误信息描述"` // 错误信息
}

// SuccessResponse 成功响应模型
type SuccessResponse struct {
	Message string `json:"message" example:"操作成功"` // 成功信息
	Data    any    `json:"data,omitempty"`         // 返回数据（可选）
}
