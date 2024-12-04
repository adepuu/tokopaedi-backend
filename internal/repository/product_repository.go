package repository

import (
	"errors"

	"github.com/adepuu/tokopaedi-backend/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Log *logrus.Logger
}

func NewProductRepository(log *logrus.Logger) *ProductRepository {
	return &ProductRepository{
		Log: log,
	}
}

func (p *ProductRepository) Save(db *gorm.DB, product *entity.Product) (*entity.Product, error) {
	return nil, errors.New("save product not implemented")
}

// GetByID implements ProductRepository.
func (p *ProductRepository) GetByID(db *gorm.DB, id int64) (*entity.Product, error) {
	return nil, errors.New("get product by id not implemented")
}

// Update implements ProductRepository.
func (p *ProductRepository) Update(db *gorm.DB, product *entity.Product) (*entity.Product, error) {
	return nil, errors.New("update product not implemented")
}

// Delete implements ProductRepository.
func (p *ProductRepository) Delete(db *gorm.DB, id int64) error {
	return errors.New("delete product not implemented")
}
