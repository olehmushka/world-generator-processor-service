package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type WriteModel interface {
	mongo.WriteModel
}
