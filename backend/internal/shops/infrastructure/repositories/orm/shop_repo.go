package orm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

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
		ID:          shop.ID(),
		Name:        shop.Name(),
		Description: shop.Description(),
		Latitude:    fmt.Sprintf("%v", shop.Latitude()),
		Longitude:   fmt.Sprintf("%v", shop.Longitude()),
	}
}

func mapShopRowToShop(row *ShopRow) *domain.Shop {
	latitude, _ := strconv.ParseFloat(row.Latitude, 64)
	longitude, _ := strconv.ParseFloat(row.Longitude, 64)

	return domain.HydrateShop(
		row.ID,
		row.Name,
		row.Description,
		latitude,
		longitude,
	)
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

func (r *ShopRepo) Get(ctx context.Context) ([]*domain.Shop, error) {
	var rows []*ShopRow
	res := r.db.WithContext(ctx).Find(&rows)
	if res.Error != nil {
		if errors.Is(res.Error, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w: %v", domain.ErrShopRepoCanNotGet, res.Error)
	}

	shops := make([]*domain.Shop, len(rows))
	for i, row := range rows {
		shops[i] = mapShopRowToShop(row)
	}

	return shops, nil
}

func (r *ShopRepo) GetByID(ctx context.Context, id string) (*domain.Shop, error) {
	var row ShopRow
	res := r.db.WithContext(ctx).First(&row, "id = ?", id)
	if res.Error != nil {
		if errors.Is(res.Error, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w: %v", domain.ErrShopRepoCanNotGetByID, res.Error)
	}

	return mapShopRowToShop(&row), nil
}
