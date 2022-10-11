package grpc

import (
	"context"
	"fmt"
	"grpc_client/pb"
	"io"
	"log"
	"os"
	"time"
)

type ClientAdapter struct {
	client pb.ProductsExplorerClient
}

func NewAdapter(client pb.ProductsExplorerClient) *ClientAdapter {
	return &ClientAdapter{client: client}
}

func (a ClientAdapter) Fetch(url string, outFile string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := a.client.Fetch(ctx, &pb.FetchRequest{Url: url})
	if err != nil {
		log.Fatalf("client.ListFeatures failed: %v", err)
	}

	file, err := os.Create(outFile)
	if err != nil {
		log.Fatalf("Cannot create file located at %s path due the error %v", outFile, err)
	}
	defer file.Close()

	for {
		fetchResponse, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("client.ListFeatures failed due the error %v", err)
		}
		numBytes, err := file.Write(fetchResponse.GetFileChunk())
		fmt.Printf("wrote %d bytes\n", numBytes)
	}
	if err := file.Sync(); err != nil {
		log.Fatalf("Cannot flush data to file due the error %v", err)
	}
}

func (a ClientAdapter) List(page uint32, count uint32, sort []int32) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var sortType []pb.ListRequest_SortingParam_SortType
	for _, s := range sort {
		sortType = append(sortType, pb.ListRequest_SortingParam_SortType(s))
	}

	listResponse, err := a.client.List(ctx, &pb.ListRequest{
		Paging: &pb.ListRequest_PagingParam{
			Page:  page,
			Count: count,
		},
		Sorting: &pb.ListRequest_SortingParam{
			Sort: sortType,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, product := range listResponse.GetProducts() {
		fmt.Println(product.LastChanged.AsTime())
		fmt.Printf("%+v\n", product)
	}
}
