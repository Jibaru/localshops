package application

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jibaru/localshops/internal/products/domain"
)

var (
	ErrCanNotCreateProduct = errors.New("can not create product")
)

type createServ struct {
	repository  domain.ProductRepo
	productServ domain.ProductServ
}

type CreateRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	ShopID        string `json:"shop_id"`
	PriceCurrency string `json:"price_currency"`
	PriceAmount   int    `json:"price_amount"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

func NewCreateServ(
	repo domain.ProductRepo,
	productServ domain.ProductServ,
) *createServ {
	return &createServ{
		repository:  repo,
		productServ: productServ,
	}
}

func (s *createServ) Execute(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	slog.InfoContext(ctx, "exec create product", "req", req)
	product, err := s.productServ.New(
		ctx,
		s.repository.NextID(),
		req.Name,
		req.Description,
		req.ShopID,
		req.PriceAmount,
		req.PriceCurrency,
	)
	if err != nil {
		slog.ErrorContext(ctx, "create product failed", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrCanNotCreateProduct, err)
	}

	err = s.repository.Save(ctx, product)
	if err != nil {
		slog.ErrorContext(ctx, "save product failed", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrCanNotCreateProduct, err)
	}

	slog.InfoContext(ctx, "product created", "id", product.ID())

	return &CreateResponse{
		ID: product.ID(),
	}, nil
}
