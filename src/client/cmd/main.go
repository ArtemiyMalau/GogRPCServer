package main

import (
	"flag"
	"grpc_client/config"
	"grpc_client/internal/adapters/grpc"
	"log"
	"os"
	"strconv"
)

type intArrayFlag []int32

func (i *intArrayFlag) String() string {
	return ""
}

func (i *intArrayFlag) Set(value string) error {
	integerValue, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	*i = append(*i, int32(integerValue))
	return nil
}

func main() {
	var fetchCmd = flag.NewFlagSet("fetch", flag.ExitOnError)
	var fetchUrl = fetchCmd.String("url", "", "Csv document's url")
	var fetchOutFile = fetchCmd.String("out_file", "./result.csv", "The csv documents path")

	var listCmd = flag.NewFlagSet("list", flag.ExitOnError)
	var listPagingPage = listCmd.Uint("paging_page", 1, "Requested page")
	var listPagingCount = listCmd.Uint("paging_count", 10, "Returned items count per page")
	var listSortingSort intArrayFlag
	listCmd.Var(&listSortingSort, "sorting_sort", "Sort settings passed as int constants")

	if len(os.Args) < 2 {
		log.Fatalln("expected 'fetch' or 'list' subcommands")
	}

	cfg := config.GetConfig()
	grpcClient, endConn := grpc.NewGRPCClient(cfg.GRPC.Host, cfg.GRPC.Port)
	adapter := grpc.NewAdapter(grpcClient)

	switch os.Args[1] {
	case "fetch":
		fetchCmd.Parse(os.Args[2:])
		adapter.Fetch(*fetchUrl, *fetchOutFile)
		endConn <- 1

	case "list":
		listCmd.Parse(os.Args[2:])
		adapter.List(uint32(*listPagingPage), uint32(*listPagingCount), listSortingSort)
		endConn <- 1
	}
}
