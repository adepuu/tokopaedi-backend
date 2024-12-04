package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/adepuu/tokopaedi-backend/internal/model"
	"github.com/adepuu/tokopaedi-backend/internal/model/converter"
	"github.com/adepuu/tokopaedi-backend/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, req *model.CreateProductRequest) (*model.CreateProductResponse, error)
	GetProductByID(ctx context.Context, id int64) (*model.CreateProductResponse, error)
	UpdateProduct(ctx context.Context, req *model.UpdateProductRequest, id int64) (*model.CreateProductResponse, error)
	DeleteProduct(ctx context.Context, id int64) (*model.DeleteProductResponse, error)
}

type productUsecase struct {
	ProductRepository *repository.ProductRepository
	Log               *logrus.Logger
	DB                *gorm.DB
}

func NewProductUsecase(
	productRepository *repository.ProductRepository,
	log *logrus.Logger,
	db *gorm.DB,
) ProductUsecase {
	return &productUsecase{
		ProductRepository: productRepository,
		Log:               log,
		DB:                db,
	}
}

// CreateProduct implements ProductUsecase.
func (p *productUsecase) CreateProduct(ctx context.Context, req *model.CreateProductRequest) (*model.CreateProductResponse, error) {
	tx := p.DB.Begin()
	product := converter.ToProductEntity(*req)
	savedProduct, err := p.ProductRepository.Save(tx, &product)

	if err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to save product")
		return nil, err
	}
	response := converter.ToCreateProductResponse(*savedProduct)
	return &response, tx.Commit().Error
}

// GetProductByID implements ProductUsecase.
func (p *productUsecase) GetProductByID(ctx context.Context, id int64) (*model.CreateProductResponse, error) {
	product, err := p.ProductRepository.GetByID(p.DB, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		p.Log.WithError(err).Error("failed to get product")
		return nil, err
	}

	response := converter.ToCreateProductResponse(*product)
	return &response, nil
}

// UpdateProduct implements ProductUsecase.
func (p *productUsecase) UpdateProduct(ctx context.Context, req *model.UpdateProductRequest, id int64) (*model.CreateProductResponse, error) {
	tx := p.DB.Begin()

	// Get existing product
	existingProduct, err := p.ProductRepository.GetByID(tx, id)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		p.Log.WithError(err).Error("failed to get product")
		return nil, err
	}

	// Update fields if provided
	if req.Name != nil {
		existingProduct.Name = *req.Name
	}
	if req.Description != nil {
		existingProduct.Description = *req.Description
	}
	if req.Price != nil {
		existingProduct.Price = *req.Price
	}
	if req.Stock != nil {
		existingProduct.Stock = *req.Stock
	}
	if req.Category != nil {
		existingProduct.Category = *req.Category
	}
	if req.Discount != nil {
		existingProduct.Discount = req.Discount
	}
	existingProduct.UpdatedAt = time.Now()

	// Save updates
	updatedProduct, err := p.ProductRepository.Update(tx, existingProduct)
	if err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to update product")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	response := converter.ToCreateProductResponse(*updatedProduct)
	return &response, nil
}

// DeleteProduct implements ProductUsecase.
func (p *productUsecase) DeleteProduct(ctx context.Context, id int64) (*model.DeleteProductResponse, error) {
	tx := p.DB.Begin()

	// Check if product exists
	_, err := p.ProductRepository.GetByID(tx, id)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		p.Log.WithError(err).Error("failed to get product")
		return nil, err
	}

	// Perform soft delete
	err = p.ProductRepository.Delete(tx, id)
	if err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to delete product")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return &model.DeleteProductResponse{
		Message: "product deleted successfully",
	}, nil
}
