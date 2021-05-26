package mgo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseOperator struct {
	*mongo.Database
}

// Collection 创建 CollectionOperator
func (o *DatabaseOperator) Collection(name string, opts ...*options.CollectionOptions) *CollectionOperator {
	return &CollectionOperator{Collection: o.Database.Collection(name, opts...)}
}
