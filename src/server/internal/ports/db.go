package ports

import (
	"context"
	"grpc_server/internal/application/domain"
)

type DBPort interface {
	DocumentInsert(ctx context.Context, dto domain.InsertDocumentDTO) (string, error)
	DocumentFindOne(ctx context.Context, url string) (domain.Document, error)
	ProductUpsert(ctx context.Context, dto domain.UpsertProductDTO) (string, error)
	ProductFind(ctx context.Context, dto domain.SelectProductDTO) ([]domain.Product, error)
}
