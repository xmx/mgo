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
	file := &gridfs.File{}
	err := o.Bucket.GetFilesCollection().FindOne(nil, bson.M{"_id": oid}).Decode(file)
	return file, err
}
