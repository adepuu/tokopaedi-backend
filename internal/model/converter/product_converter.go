package converter

import (
	"time"

	"github.com/adepuu/tokopaedi-backend/internal/entity"
	"github.com/adepuu/tokopaedi-backend/internal/model"
)

func ToProductEntity(req model.CreateProductRequest) entity.Product {
	return entity.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		Discount:    req.Discount,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ToCreateProductResponse(product entity.Product) model.CreateProductResponse {
	return model.CreateProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Category:    product.Category,
		Discount:    product.Discount,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
	}
}

func ToUpdateProductResponse(product entity.Product) model.UpdateProductResponse {
	return model.UpdateProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Category:    product.Category,
		Discount:    product.Discount,
		UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
	}
}
