package mongo

import (
	"fmt"

	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AnyIDToObjectID(id any) (primitive.ObjectID, error) {
	if id == nil {
		return primitive.NilObjectID, nil
	}

	oid, ok := id.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not cast to oid (in=%+v)", id))
	}

	return oid, nil
}

func appendResult(result *BulkWriteResult, item *mongo.BulkWriteResult) *BulkWriteResult {
	newResult := &BulkWriteResult{}
	newResult.InsertedCount = result.InsertedCount + item.InsertedCount
	newResult.MatchedCount = result.MatchedCount + item.MatchedCount
	newResult.ModifiedCount = result.ModifiedCount + item.ModifiedCount
	newResult.DeletedCount = result.DeletedCount + item.DeletedCount
	newResult.UpsertedCount = result.UpsertedCount + item.UpsertedCount
	newResult.UpsertedIDs = make(map[int64]interface{})

	for key, value := range result.UpsertedIDs {
		if _, ok := newResult.UpsertedIDs[key]; ok {
			continue
		}
		newResult.UpsertedIDs[key] = value
	}
	for ia, va := range item.UpsertedIDs {
		if _, ok := newResult.UpsertedIDs[ia]; ok {
			continue
		}
		newResult.UpsertedIDs[ia] = va
	}

	return newResult
}
