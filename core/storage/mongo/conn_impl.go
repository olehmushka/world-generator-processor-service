package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type conn struct {
	client *mongo.Client

	maxBulkItemsSize int
}

func NewConn(username, password, uri string, maxBulkItemsSize int) (Connection, error) {
	ctx := context.Background()

	creds := options.Credential{
		Username: username,
		Password: password,
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(creds))
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not create mongo connection")
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not ping mongo")
	}

	if maxBulkItemsSize == 0 {
		maxBulkItemsSize = DefaultMaxBulkItemsSize
	}

	return &conn{
		client:           client,
		maxBulkItemsSize: maxBulkItemsSize,
	}, nil
}

func (c *conn) Ping(ctx context.Context) error {
	if err := c.client.Ping(ctx, nil); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not ping mongodb server")
	}

	return nil
}

func (c *conn) Disconnect(ctx context.Context) error {
	if err := c.client.Disconnect(ctx); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not disconnect from mongodb server")
	}

	return nil
}

func (c *conn) InsertOne(ctx context.Context, dbName, collName string, doc any, opts ...*options.InsertOneOptions) (primitive.ObjectID, error) {
	coll := c.client.Database(dbName).Collection(collName)
	insertRes, err := coll.InsertOne(ctx, doc, opts...)
	if err != nil {
		return primitive.NilObjectID, wrapped_error.NewInternalServerError(err, "can not insert doc into mongodb")
	}

	oid, err := AnyIDToObjectID(insertRes.InsertedID)
	if err != nil {
		return primitive.NilObjectID, wrapped_error.NewInternalServerError(err, "can not cast inserting result to oid")
	}

	return oid, nil
}

func (c *conn) InsertMany(
	ctx context.Context,
	dbName,
	collName string,
	docs []any,
	opts ...*options.InsertManyOptions,
) ([]primitive.ObjectID, error) {
	coll := c.client.Database(dbName).Collection(collName)

	insertResults, err := coll.InsertMany(ctx, docs, opts...)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not insert many docs")
	}

	var resultIDs []primitive.ObjectID
	for _, insertedID := range insertResults.InsertedIDs {
		oid, err := AnyIDToObjectID(insertedID)
		if err != nil {
			return nil, err
		}

		resultIDs = append(resultIDs, oid)
	}

	return resultIDs, nil
}

func (c *conn) CountDocuments(ctx context.Context, dbName, collName string, filter any, opts ...*options.CountOptions) (int, error) {
	coll := c.client.Database(dbName).Collection(collName)

	count, err := coll.CountDocuments(ctx, filter, opts...)
	if err != nil {
		return 0, wrapped_error.NewInternalServerError(err, "can not count documents")
	}

	return int(count), nil
}

func (c *conn) Find(ctx context.Context, dbName, collName string, filter any, opts ...*options.FindOptions) (Cursor, error) {
	coll := c.client.Database(dbName).Collection(collName)

	cursor, err := coll.Find(ctx, filter, opts...)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not find docs via mongodb")
	}

	return cursor, nil
}

func (c *conn) FindOne(ctx context.Context, dbName, collName string, filter any, opts ...*options.FindOneOptions) (SingleResult, error) {
	coll := c.client.Database(dbName).Collection(collName)

	singleRes := coll.FindOne(ctx, filter, opts...)
	if errors.Is(singleRes.Err(), mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err := singleRes.Err(); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not find one doc via mongodb")
	}

	return singleRes, nil
}

func (c *conn) BulkWrite(
	ctx context.Context,
	dbName,
	collName string,
	operations []WriteModel,
	opts ...*options.BulkWriteOptions,
) (*BulkWriteResult, error) {
	coll := c.client.Database(dbName).Collection(collName)

	var mongoOperations []mongo.WriteModel
	result := &BulkWriteResult{}

	var num int
	for _, op := range operations {
		num++
		mongoOperations = append(mongoOperations, op.(mongo.WriteModel))

		if num == len(operations) {
			bulkWriteResult, err := c.bulkWrite(ctx, coll, mongoOperations, opts...)
			if err != nil {
				return nil, err
			}

			return appendResult(result, bulkWriteResult), nil
		}

		if len(mongoOperations) == c.maxBulkItemsSize {
			bulkWriteResult, err := c.bulkWrite(ctx, coll, mongoOperations, opts...)
			if err != nil {
				return nil, err
			}
			result = appendResult(result, bulkWriteResult)
			mongoOperations = []mongo.WriteModel{}
		}
	}

	return result, nil
}

func (c *conn) bulkWrite(
	ctx context.Context,
	coll *mongo.Collection,
	writeModels []mongo.WriteModel,
	opts ...*options.BulkWriteOptions,
) (*mongo.BulkWriteResult, error) {
	result, err := coll.BulkWrite(ctx, writeModels, opts...)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not bulk write")
	}

	return result, nil
}

func (c *conn) UpdateOne(ctx context.Context, dbName string, collName string,
	filter, update any, opts ...*options.UpdateOptions) (*UpdateResult, error) {
	coll := c.client.Database(dbName).Collection(collName)

	updateResult, err := coll.UpdateOne(ctx, filter, update, opts...)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not update one doc")
	}

	oid, err := AnyIDToObjectID(updateResult.UpsertedID)
	if err != nil {
		return nil, err
	}

	return &UpdateResult{
		MatchedCount:  updateResult.MatchedCount,
		ModifiedCount: updateResult.ModifiedCount,
		UpsertedCount: updateResult.UpsertedCount,
		UpsertedID:    oid,
	}, nil
}

func (c *conn) Truncate(ctx context.Context, dbName, collName string) error {
	coll := c.client.Database(dbName).Collection(collName)
	if err := coll.Drop(ctx); err != nil {
		return wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not trancate collection (db_name=%s, coll_name=%s)", dbName, collName))
	}

	return nil
}
func (c *conn) DeleteOne(ctx context.Context, dbName string, collName string,
	filter any, opts ...*options.DeleteOptions) (int, error) {
	coll := c.client.Database(dbName).Collection(collName)

	deleteResult, err := coll.DeleteOne(ctx, filter, opts...)
	if err != nil {
		return 0, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not delete one by filter (db_name=%s, coll_name=%s, filter=%+v)", dbName, collName, filter))
	}

	return int(deleteResult.DeletedCount), nil
}

func (c *conn) DeleteMany(ctx context.Context, dbName string, collName string,
	filter any, opts ...*options.DeleteOptions) (int, error) {
	coll := c.client.Database(dbName).Collection(collName)

	deleteResult, err := coll.DeleteMany(ctx, filter, opts...)
	if err != nil {
		return 0, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not delete many by filter (db_name=%s, coll_name=%s, filter=%+v)", dbName, collName, filter))
	}

	return int(deleteResult.DeletedCount), nil
}
