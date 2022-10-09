package ports

import (
	"context"
	"grpc_server/internal/application/domain"
)

type APIPort interface {
	ProductSelectAll(ctx context.Context, dto domain.SelectProductDTO) ([]domain.Product, error)
	FetchCsv(ctx context.Context, url string) ([]byte, error)
}
