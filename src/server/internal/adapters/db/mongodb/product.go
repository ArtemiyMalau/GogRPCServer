package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"grpc_server/internal/application/domain"
	"time"
)

// ProductUpsert upsert product using product's dto
// Return product's _id ONLY in case when product was inserted not updated
func (a *Adapter) ProductUpsert(ctx context.Context, dto domain.UpsertProductDTO) (string, error) {
	opts := options.Update().SetUpsert(true)
	update := bson.D{
		{"$set", bson.M{"price": dto.Price}},
		{"$inc", bson.M{"price_changed_count": 1}},
		{"$set", bson.M{"date_last_price_change": time.Now()}},
		{"$push", bson.M{"date_price_change": domain.PriceHistory{
			Price:           dto.Price,
			DatePriceChange: time.Now(),
		}}},
	}
	updateResult, err := a.productC.UpdateOne(ctx, bson.M{"name": dto.Name}, update, opts)
	if err != nil {
		return "", fmt.Errorf("failed to Upsert %s collection at UpdateOne method due to error %v", a.productC.Name(), err)
	}

	if updateResult.UpsertedCount == 1 {
		oid, ok := updateResult.UpsertedID.(primitive.ObjectID)
		if !ok {
			return "", fmt.Errorf("failed to cast operationID to ObjectID")
		}
		return oid.Hex(), nil
	} else if updateResult.ModifiedCount == 1 {
		return "", nil
	} else {
		// TODO Handle case when product was neither created nor updated
		return "", fmt.Errorf("no objects was affected using dto %+v", dto)
	}

}

func (a *Adapter) ProductFind(ctx context.Context, dto domain.SelectProductDTO) ([]domain.Product, error) {
	sortParams := bson.D{}
	for _, s := range dto.Sort {
		var e bson.E
		switch s {
		case domain.SORT_PRICE_ASC:
			e = bson.E{Key: "price", Value: 1}
		case domain.SORT_PRICE_DESC:
			e = bson.E{Key: "price", Value: -1}
		case domain.SORT_NAME_ASC:
			e = bson.E{Key: "name", Value: 1}
		case domain.SORT_NAME_DESC:
			e = bson.E{Key: "name", Value: -1}
		case domain.SORT_CHANGE_COUNT_ASC:
			e = bson.E{Key: "price_changed_count", Value: 1}
		case domain.SORT_CHANGE_COUNT_DESC:
			e = bson.E{Key: "price_changed_count", Value: -1}
		case domain.SORT_LAST_CHANGED_ASC:
			e = bson.E{Key: "date_last_price_change", Value: 1}
		case domain.SORT_LAST_CHANGED_DESC:
			e = bson.E{Key: "date_last_price_change", Value: -1}
		}
		sortParams = append(sortParams, e)
	}

	opts := options.Find().
		SetSort(sortParams).
		SetSkip(int64(dto.Offset)).
		SetLimit(int64(dto.Limit)).
		SetProjection(bson.D{{"_id", 0}}) // Exclude unnecessary _id column
	cur, err := a.productC.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, fmt.Errorf("cannot obtain select cursor due to error %v", err)
	}

	products := make([]domain.Product, 0)
	if err := cur.All(ctx, &products); err != nil {
		return nil, fmt.Errorf("cannot retrieve products using cursor.All due to error %v", err)
	}

	return products, nil
}
