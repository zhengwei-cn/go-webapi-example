package controllers

import (
	"net/http"
	"strconv"
	"webapi/models"
	"webapi/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		productService: services.NewProductService(db),
	}
}

// CreateProduct godoc
// @Summary 创建新产品
// @Description 创建一个新的产品记录，需要提供产品基本信息
// @Tags products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product body models.CreateProductRequest true "产品创建信息"
// @Success 201 {object} models.Product "创建成功，返回产品详细信息"
// @Failure 400 {object} map[string]string "请求参数错误"
// @Failure 401 {object} map[string]string "未授权访问"
// @Failure 500 {object} map[string]string "服务器内部错误"
// @Router /products [post]
func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var req models.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := c.productService.CreateProduct(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

// GetProduct godoc
// @Summary 获取单个产品详情
// @Description 根据产品ID获取产品的详细信息，包括关联的用户信息
// @Tags products
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "产品ID" minimum(1) example(1)
// @Success 200 {object} models.Product "获取成功，返回产品详细信息"
// @Failure 400 {object} map[string]string "请求参数错误"
// @Failure 401 {object} map[string]string "未授权访问"
// @Failure 404 {object} map[string]string "产品不存在"
// @Router /products/{id} [get]
func (c *ProductController) GetProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := c.productService.GetProductByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// GetProducts godoc
// @Summary 获取产品列表
// @Description 获取所有产品的列表，包括产品基本信息和关联用户信息
// @Tags products
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.Product "获取成功，返回产品列表"
// @Failure 401 {object} map[string]string "未授权访问"
// @Failure 500 {object} map[string]string "服务器内部错误"
// @Router /products [get]
func (c *ProductController) GetProducts(ctx *gin.Context) {
	products, err := c.productService.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// UpdateProduct godoc
// @Summary 更新产品信息
// @Description 根据产品ID更新产品的详细信息，支持部分字段更新
// @Tags products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "产品ID" minimum(1) example(1)
// @Param product body models.UpdateProductRequest true "产品更新信息"
// @Success 200 {object} models.Product "更新成功，返回更新后的产品信息"
// @Failure 400 {object} map[string]string "请求参数错误"
// @Failure 401 {object} map[string]string "未授权访问"
// @Failure 404 {object} map[string]string "产品不存在"
// @Failure 500 {object} map[string]string "服务器内部错误"
// @Router /products/{id} [put]
func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var req models.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := c.productService.UpdateProduct(uint(id), &req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary 删除产品
// @Description 根据产品ID删除指定产品（软删除）
// @Tags products
// @Security ApiKeyAuth
// @Param id path int true "产品ID" minimum(1) example(1)
// @Success 204 "删除成功，无返回内容"
// @Failure 400 {object} map[string]string "请求参数错误"
// @Failure 401 {object} map[string]string "未授权访问"
// @Failure 500 {object} map[string]string "服务器内部错误"
// @Router /products/{id} [delete]
func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := c.productService.DeleteProduct(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
