package application

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jibaru/localshops/internal/shops/domain"
)

var (
	ErrCanNotCreateShop = errors.New("can not create shop")
)

type createServ struct {
	repository domain.ShopRepo
}

type CreateRequest struct {
	ShopFieldsDTO
}

type CreateResponse struct {
	ID string `json:"id"`
}

func NewCreateServ(
	repository domain.ShopRepo,
) *createServ {
	return &createServ{
		repository: repository,
	}
}

func (s *createServ) Execute(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	slog.InfoContext(ctx, "exec create shop", "req", req)
	shop, err := domain.NewShop(
		s.repository.NextID(),
		req.Name,
		req.Description,
		req.Latitude,
		req.Longitude,
	)
	if err != nil {
		slog.ErrorContext(ctx, "new shop failed", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrCanNotCreateShop, err)
	}

	err = s.repository.Save(ctx, shop)
	if err != nil {
		slog.ErrorContext(ctx, "save shop failed", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrCanNotCreateShop, err)
	}

	slog.InfoContext(ctx, "shop created", "shop", shop)

	return &CreateResponse{
		ID: shop.ID(),
	}, nil
}
