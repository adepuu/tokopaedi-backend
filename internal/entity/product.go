package entity

import (
	"time"
)

type Product struct {
	ID          int64      `gorm:"column:id;primaryKey;autoIncrement"`
	Name        string     `gorm:"column:name"`
	Description string     `gorm:"column:description"`
	Price       float64    `gorm:"column:price;type:numeric"`
	Stock       int        `gorm:"column:stock"`
	Category    string     `gorm:"column:category"`
	Discount    *float64   `gorm:"column:discount;type:numeric"` // nullable/optional field
	CreatedAt   time.Time  `gorm:"column:created_at;type:timestamptz;default:now()"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;type:timestamptz;default:now()"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:timestamptz"` // nullable/optional field for soft delete
}

func (p *Product) TableName() string {
	return "products"
}
