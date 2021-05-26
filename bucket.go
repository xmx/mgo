package mgo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type BucketFile struct {
	ID         interface{} `bson:"_id"`
	Length     int64       `bson:"length"`
	ChunkSize  int32       `bson:"chunkSize"`
	UploadDate time.Time   `bson:"uploadDate"`
	Name       string      `bson:"filename"`
	Metadata   bson.Raw    `bson:"metadata"`
}

type BucketOperator struct {
	*gridfs.Bucket
}

// File 根据文件ID获取文件的信息
func (o *BucketOperator) File(id interface{}, opts ...*options.FindOneOptions) (BucketFile, error) {
	oid := parseID(id)
	var bf BucketFile
	err := o.Bucket.GetFilesCollection().FindOne(nil, oid, opts...).Decode(&bf)
	return bf, err
}
