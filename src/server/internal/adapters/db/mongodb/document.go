package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"grpc_server/internal/application/domain"
)

func (a *Adapter) DocumentFindOne(ctx context.Context, url string) (document domain.Document, err error) {
	result := a.documentC.FindOne(ctx, bson.M{"url": url})

	if result.Err() != nil {
		// TODO implement ErrNoDocuments error handling
		return document, fmt.Errorf("failed to find one document by url: %s due to error %v", url, result.Err())
	}
	if err := result.Decode(&document); err != nil {
		return document, fmt.Errorf("failed to decode finded by url: %s document due to error %v", url, result.Err())
	}

	return document, err
}

func (a *Adapter) DocumentInsert(ctx context.Context, dto domain.InsertDocumentDTO) (string, error) {
	document := bson.D{
		{"url", dto.Url},
	}
	insertResult, err := a.documentC.InsertOne(ctx, document)
	if err != nil {
		return "", fmt.Errorf("failed to Insert %s collection at Insert method due to error %v", a.documentC.Name(), err)
	}

	oid, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to cast operationID to ObjectID")
	}
	return oid.Hex(), nil
}
