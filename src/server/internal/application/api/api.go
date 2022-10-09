package api

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
	"grpc_server/internal/application/apperror"
	"grpc_server/internal/application/domain"
	"grpc_server/internal/ports"
	"io"
	"log"
	"net/http"
)

type Application struct {
	db ports.DBPort
}

var _ ports.APIPort = (*Application)(nil)

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a Application) ProductSelectAll(ctx context.Context, dto domain.SelectProductDTO) ([]domain.Product, error) {
	return a.db.ProductFind(ctx, dto)
}

func (a Application) FetchCsv(ctx context.Context, url string) (csvBytes []byte, err error) {
	log.Printf("Fetch csv with url: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("cannot access to resource located at %s due to error %v", url, err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("cannot load file: status is %d", resp.StatusCode)
	}

	_, err = a.db.DocumentFindOne(ctx, url)
	if err != nil {
		if !errors.Is(err, apperror.ErrNotFound) {
			fmt.Println(err)
			return nil, err
		}

		// Fetched document isn't stored in database yet
		if _, err := a.db.DocumentInsert(ctx, domain.InsertDocumentDTO{Url: url}); err != nil {
			return nil, err
		}

		var products []domain.UpsertProductDTO
		var unmarshalBuf bytes.Buffer
		returnBytesTee := io.TeeReader(resp.Body, &unmarshalBuf)
		csvBytes, err = io.ReadAll(returnBytesTee)
		if err != nil {
			return nil, fmt.Errorf("error read resp.Body using ReadAll via TeeReader due to error %v", err)
		}
		if err := gocsv.Unmarshal(&unmarshalBuf, &products); err != nil {
			return nil, fmt.Errorf("error Unmarshal resp.Body using gocsv.ReadAll due to error %v", err)
		}
		for _, product := range products {
			log.Printf("Product is %+v\n", product)
			if _, err := a.db.ProductUpsert(ctx, product); err != nil {
				return nil, err
			}
		}

		return csvBytes, nil
	}

	csvBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error read resp.Body using ReadAll due to error %v", err)
	}
	return csvBytes, nil
}
