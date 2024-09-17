package remoteservices

import (
	"context"
	"fmt"

	shops_application "github.com/jibaru/localshops/internal/shops/application"

	"github.com/jibaru/localshops/internal/products/domain"
)

type ShopServ struct {
	getServ shops_application.GetServ
}

func NewShopServ(getServ shops_application.GetServ) *ShopServ {
	return &ShopServ{
		getServ: getServ,
	}
}

func (s *ShopServ) ExistsByID(ctx context.Context, id string) (bool, error) {
	_, err := s.getServ.Execute(ctx, id)
	if err != nil {
		return false, fmt.Errorf("%w: %v", domain.ErrRemoteShopServExistsByID, err)
	}

	return true, nil
}
