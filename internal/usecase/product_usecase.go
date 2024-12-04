package usecase

import (
	"context"

	"github.com/adepuu/tokopaedi-backend/internal/model"
	"github.com/adepuu/tokopaedi-backend/internal/model/converter"
	"github.com/adepuu/tokopaedi-backend/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, req *model.CreateProductRequest) (*model.CreateProductResponse, error)
	GetProductByID(ctx context.Context, id int64) (*model.CreateProductResponse, error)
	UpdateProduct(ctx context.Context, req *model.CreateProductRequest, id int64) (*model.CreateProductResponse, error)
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

// DeleteProduct implements ProductUsecase.
func (p *productUsecase) DeleteProduct(ctx context.Context, id int64) (*model.DeleteProductResponse, error) {
	panic("unimplemented")
}

// GetProductByID implements ProductUsecase.
func (p *productUsecase) GetProductByID(ctx context.Context, id int64) (*model.CreateProductResponse, error) {
	panic("unimplemented")
}

// UpdateProduct implements ProductUsecase.
func (p *productUsecase) UpdateProduct(ctx context.Context, req *model.CreateProductRequest, id int64) (*model.CreateProductResponse, error) {
	panic("unimplemented")
}
