package services

import (
	"webapi/models"

	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db: db}
}

func (s *ProductService) CreateProduct(req *models.CreateProductRequest) (*models.Product, error) {
	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		UserID:      req.UserID,
	}

	if err := s.db.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := s.db.Preload("User").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := s.db.Preload("User").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) UpdateProduct(id uint, req *models.UpdateProductRequest) (*models.Product, error) {
	product, err := s.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price > 0 {
		product.Price = req.Price
	}
	if req.Stock >= 0 {
		product.Stock = req.Stock
	}

	if err := s.db.Save(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.db.Delete(&models.Product{}, id).Error
}
