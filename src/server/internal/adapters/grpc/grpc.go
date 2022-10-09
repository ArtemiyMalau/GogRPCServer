package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc_server/internal/application/domain"
	"grpc_server/proto"
	"log"
)

func (a Adapter) Fetch(request *proto.FetchRequest, stream proto.ProductsExplorer_FetchServer) error {
	csvBytes, err := a.api.FetchCsv(context.Background(), request.GetUrl())
	if err != nil {
		return err
	}

	err = stream.Send(&proto.FetchResponse{FileChunk: csvBytes})
	if err != nil {
		log.Println("error while sending chunk:", err)
		return err
	}

	return nil
}

func (a Adapter) List(ctx context.Context, request *proto.ListRequest) (*proto.ListResponse, error) {
	var (
		sort   []int
		limit  = uint(request.GetPaging().GetCount())
		offset = uint((request.GetPaging().GetPage() - 1) * request.GetPaging().GetCount())
	)
	for _, sortParam := range request.Sorting.GetSort() {
		sort = append(sort, int(sortParam))
	}

	products, err := a.api.ProductSelectAll(ctx, domain.SelectProductDTO{
		Offset: offset,
		Limit:  limit,
		Sort:   sort,
	})
	if err != nil {
		return nil, err
	}

	response := &proto.ListResponse{}
	for _, p := range products {
		response.Products = append(response.Products, &proto.Product{
			Name:        p.Name,
			Price:       p.Price,
			ChangeCount: p.PriceChangedCount,
			LastChanged: timestamppb.New(p.DateLastPriceChange),
		})
	}

	return response, nil
}
