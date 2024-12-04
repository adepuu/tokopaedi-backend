package usecase

import (
	"context"
	"errors"

	"github.com/adepuu/tokopaedi-backend/internal/model"
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
	return nil, errors.New("create product not implemented")
}

// GetProductByID implements ProductUsecase.
func (p *productUsecase) GetProductByID(ctx context.Context, id int64) (*model.CreateProductResponse, error) {
	return nil, errors.New("get product by id not implemented")
}

// UpdateProduct implements ProductUsecase.
func (p *productUsecase) UpdateProduct(ctx context.Context, req *model.UpdateProductRequest, id int64) (*model.CreateProductResponse, error) {
	return nil, errors.New("update product not implemented")
}

// DeleteProduct implements ProductUsecase.
func (p *productUsecase) DeleteProduct(ctx context.Context, id int64) (*model.DeleteProductResponse, error) {
	return nil, errors.New("delete product not implemented")
}
