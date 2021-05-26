package mgo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type BucketOperator struct {
	*gridfs.Bucket
}

// File 根据文件ID获取文件的信息
func (o *BucketOperator) File(id interface{}) (*gridfs.File, error) {
	oid := parseID(id)
	cursor, err := o.Bucket.GetFilesCollection().Find(nil, bson.E{Key: "_id", Value: oid})
	if err != nil {
		return nil, err
	}
	f := &gridfs.File{}
	err = cursor.Decode(f)
	return f, err
}
