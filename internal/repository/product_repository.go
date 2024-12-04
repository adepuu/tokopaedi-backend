package repository

import (
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
	err := db.Create(product).Error
	if err != nil {
		p.Log.Error(err)
		return nil, err
	}

	return product, nil
}

// Delete implements ProductRepository.
func (p *ProductRepository) Delete(db *gorm.DB, id int64) error {
	err := db.Where("id = ?", id).Delete(&entity.Product{}).Error
	if err != nil {
		p.Log.Error(err)
		return err
	}
	return nil
}

// Update implements ProductRepository.
func (p *ProductRepository) Update(db *gorm.DB, product *entity.Product) (*entity.Product, error) {
	err := db.Save(product).Error
	if err != nil {
		p.Log.Error(err)
		return nil, err
	}
	return product, nil
}

// GetByID implements ProductRepository.
func (p *ProductRepository) GetByID(db *gorm.DB, id int64) (*entity.Product, error) {
	var product entity.Product
	err := db.Where("id = ?", id).First(&product).Error
	if err != nil {
		p.Log.Error(err)
		return nil, err
	}
	return &product, nil
}
