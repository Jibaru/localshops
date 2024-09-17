package application

import (
	"context"
	"errors"
	"fmt"

	"github.com/jibaru/localshops/internal/shops/domain"
)

var (
	ErrCanNotGetShopByID = errors.New("can not get shop by id")
)

type GetResponse struct {
	ShopDTO
}

type getServ struct {
	repo domain.ShopRepo
}

func NewGetServ(
	repo domain.ShopRepo,
) *getServ {
	return &getServ{
		repo: repo,
	}
}

func (s *getServ) Execute(ctx context.Context, id string) (*GetResponse, error) {
	shop, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCanNotGetShopByID, err)
	}

	if shop == nil {
		return nil, fmt.Errorf("%w: %v", ErrCanNotGetShopByID, "there is no shop for the given id")
	}

	return &GetResponse{
		ShopDTO: ShopDTO{
			ID: shop.ID(),
			ShopFieldsDTO: ShopFieldsDTO{
				Name:        shop.Name(),
				Description: shop.Description(),
				Latitude:    shop.Latitude(),
				Longitude:   shop.Longitude(),
			},
		},
	}, nil
}
