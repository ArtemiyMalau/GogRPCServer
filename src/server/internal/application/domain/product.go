package domain

import (
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
