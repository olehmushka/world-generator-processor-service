package mongo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestAnyIDToObjectID(t *testing.T) {
	validOID := primitive.NewObjectID()

	tCases := map[string]struct {
		input          any
		expectedOutput primitive.ObjectID
		expectedErr    bool
	}{
		"should return nil object id for nil input": {
			input:          nil,
			expectedOutput: primitive.NilObjectID,
		},
		"should return the object id unchanged for a valid object id input": {
			input:          validOID,
			expectedOutput: validOID,
		},
		"should return an error for a non object id input": {
			input:       "not-an-object-id",
			expectedErr: true,
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			out, err := AnyIDToObjectID(tc.input)
			if tc.expectedErr {
				require.Error(tt, err)
				return
			}
			require.NoError(tt, err)
			assert.Equal(tt, tc.expectedOutput, out)
		})
	}
}

func TestAppendResult(t *testing.T) {
	oid1 := primitive.NewObjectID()
	oid2 := primitive.NewObjectID()

	tCases := map[string]struct {
		result         *BulkWriteResult
		item           *mongo.BulkWriteResult
		expectedOutput *BulkWriteResult
	}{
		"should sum counts and merge upserted ids of two empty results": {
			result: &BulkWriteResult{},
			item:   &mongo.BulkWriteResult{},
			expectedOutput: &BulkWriteResult{
				UpsertedIDs: map[int64]interface{}{},
			},
		},
		"should sum counts across both results": {
			result: &BulkWriteResult{
				InsertedCount: 1,
				MatchedCount:  2,
				ModifiedCount: 3,
				DeletedCount:  4,
				UpsertedCount: 5,
				UpsertedIDs:   map[int64]interface{}{0: oid1},
			},
			item: &mongo.BulkWriteResult{
				InsertedCount: 10,
				MatchedCount:  20,
				ModifiedCount: 30,
				DeletedCount:  40,
				UpsertedCount: 50,
				UpsertedIDs:   map[int64]interface{}{1: oid2},
			},
			expectedOutput: &BulkWriteResult{
				InsertedCount: 11,
				MatchedCount:  22,
				ModifiedCount: 33,
				DeletedCount:  44,
				UpsertedCount: 55,
				UpsertedIDs: map[int64]interface{}{
					0: oid1,
					1: oid2,
				},
			},
		},
		"should keep the existing upserted id when the same index is present in both": {
			result: &BulkWriteResult{
				UpsertedIDs: map[int64]interface{}{0: oid1},
			},
			item: &mongo.BulkWriteResult{
				UpsertedIDs: map[int64]interface{}{0: oid2},
			},
			expectedOutput: &BulkWriteResult{
				UpsertedIDs: map[int64]interface{}{0: oid1},
			},
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			out := appendResult(tc.result, tc.item)
			assert.Equal(tt, tc.expectedOutput, out)
		})
	}
}
