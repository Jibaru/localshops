package orm

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jibaru/localshops/internal/shops/domain"
)

type ShopRepo struct {
	db *gorm.DB
}

type ShopRow struct {
	ID          string `gorm:"id"`
	Name        string `gorm:"name"`
	Description string `gorm:"description"`
	Latitude    string `gorm:"latitude"`
	Longitude   string `gorm:"longitude"`
}

func NewShopRepo(db *gorm.DB) *ShopRepo {
	return &ShopRepo{
		db: db,
	}
}

func (r *ShopRow) TableName() string {
	return "shops"
}

func mapShopToShopRow(shop *domain.Shop) *ShopRow {
	return &ShopRow{
		ID:          shop.ID,
		Name:        shop.Name,
		Description: shop.Description,
		Latitude:    fmt.Sprintf("%v", shop.Location.Latitude),
		Longitude:   fmt.Sprintf("%v", shop.Location.Longitude),
	}
}

func (r *ShopRepo) NextID() string {
	return uuid.NewString()
}

func (r *ShopRepo) Save(ctx context.Context, shop *domain.Shop) error {
	res := r.db.WithContext(ctx).Create(mapShopToShopRow(shop))
	if res.Error != nil {
		return fmt.Errorf("%w: %v", domain.ErrShopRepoCanNotSave, res.Error)
	}

	return nil
}
