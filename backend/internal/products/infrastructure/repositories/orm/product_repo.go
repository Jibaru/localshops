package orm

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jibaru/localshops/internal/products/domain"
)

type ProductRepo struct {
	db *gorm.DB
}

type ProductRow struct {
	ID            string `gorm:"id"`
	Name          string `gorm:"name"`
	Description   string `gorm:"description"`
	PriceAmount   int    `gorm:"price_amount"`
	PriceCurrency string `gorm:"price_currency"`
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func mapProductToProductRow(product *domain.Product) *ProductRow {
	return &ProductRow{
		ID:            product.ID(),
		Name:          product.Name(),
		Description:   product.Description(),
		PriceAmount:   product.PriceAmount(),
		PriceCurrency: product.PriceCurrency(),
	}
}

func (r *ProductRow) TableName() string {
	return "products"
}

func (r *ProductRepo) NextID() string {
	return uuid.NewString()
}

func (r *ProductRepo) Save(ctx context.Context, product *domain.Product) error {
	res := r.db.WithContext(ctx).Create(mapProductToProductRow(product))
	if res.Error != nil {
		return fmt.Errorf("%w: %v", domain.ErrProductRepoCanNotSave, res.Error)
	}

	return nil
}
