package application

import "context"

type CreateServ interface {
	Execute(ctx context.Context, req CreateRequest) (*CreateResponse, error)
}
