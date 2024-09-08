package application

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jibaru/localshops/internal/shops/domain"
)

var (
	ErrCanNotGetAllShops = errors.New("can not get all shops")
)

type GetAllServ struct {
	repo domain.ShopRepo
}

type GetAllResponse struct {
	Data []ShopDTO `json:"data"`
}

func NewGetAllServ(
	repo domain.ShopRepo,
) *GetAllServ {
	return &GetAllServ{
		repo: repo,
	}
}

func (s *GetAllServ) Execute(ctx context.Context) (*GetAllResponse, error) {
	slog.InfoContext(ctx, "exec get all shops")
	shops, err := s.repo.Get(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "get all shops failed", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrCanNotGetAllShops, err)
	}

	resp := GetAllResponse{}
	for _, shop := range shops {
		resp.Data = append(resp.Data, ShopDTO{
			ID: shop.ID(),
			ShopFieldsDTO: ShopFieldsDTO{
				Name:        shop.Name(),
				Description: shop.Description(),
				Latitude:    shop.Latitude(),
				Longitude:   shop.Longitude(),
			},
		})
	}

	return &resp, nil
}
