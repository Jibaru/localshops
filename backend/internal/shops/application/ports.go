package application

import "context"

type CreateServ interface {
	Execute(ctx context.Context, req CreateRequest) (*CreateResponse, error)
}

type GetAllServ interface {
	Execute(ctx context.Context) (*GetAllResponse, error)
}

type GetServ interface {
	Execute(ctx context.Context, id string) (*GetResponse, error)
}
