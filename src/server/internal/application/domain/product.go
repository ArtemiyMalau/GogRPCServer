package domain

import (
	"grpc_server/proto"
	"time"
)

type PriceHistory struct {
	Price           uint32    `json:"price" bson:"price"`
	DatePriceChange time.Time `json:"date_price_change" bson:"date_price_change"`
}

type Product struct {
	ID                  string         `json:"id" bson:"_id,omitempty"`
	Name                string         `json:"name" bson:"name"`
	Price               uint32         `json:"price" bson:"price"`
	PriceChangedCount   uint32         `json:"price_changed_count" bson:"price_changed_count"`
	DateLastPriceChange time.Time      `json:"date_last_price_change" bson:"date_last_price_change"`
	PriceHistory        []PriceHistory `json:"price_history" bson:"price_history"`
}

type UpsertProductDTO struct {
	Name  string `json:"name" csv:"product_name"`
	Price uint32 `json:"price" csv:"price"`
}

type SelectProductDTO struct {
	Offset uint
	Limit  uint
	Sort   []int
}

const (
	SORT_PRICE_ASC         = int(proto.ListRequest_SortingParam_PRICE_ASC)
	SORT_PRICE_DESC        = int(proto.ListRequest_SortingParam_PRICE_DESC)
	SORT_NAME_ASC          = int(proto.ListRequest_SortingParam_NAME_ASC)
	SORT_NAME_DESC         = int(proto.ListRequest_SortingParam_NAME_DESC)
	SORT_CHANGE_COUNT_ASC  = int(proto.ListRequest_SortingParam_CHANGE_COUNT_ASC)
	SORT_CHANGE_COUNT_DESC = int(proto.ListRequest_SortingParam_CHANGE_COUNT_DESC)
	SORT_LAST_CHANGED_ASC  = int(proto.ListRequest_SortingParam_LAST_CHANGED_ASC)
	SORT_LAST_CHANGED_DESC = int(proto.ListRequest_SortingParam_LAST_CHANGED_DESC)
)
