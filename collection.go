package mgo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

type CollectionOperator struct {
	*mongo.Collection
}

// Bucket 获取与当前 collection 同名的 bucket
func (o *CollectionOperator) Bucket() (*gridfs.Bucket, error) {
	opt := options.GridFSBucket().SetName(o.Name())
	return gridfs.NewBucket(o.Database(), opt)
}

// InsertMany 批量插入
func (o *CollectionOperator) InsertMany(documents interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	docs := toSlice(documents)
	return o.Collection.InsertMany(nil, docs, opts...)
}

// DeleteID 通过ID删除记录, 数据库中的ID必须是 _id 且类型是 ObjectID
func (o *CollectionOperator) DeleteID(id interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	oid := parseID(id)
	return o.Collection.DeleteOne(nil, bson.M{"_id": oid}, opts...)
}

// DeleteBatchID 通过ID批量删除记录, 数据库中的ID必须是 _id 且类型是 ObjectID
func (o *CollectionOperator) DeleteBatchID(ids []string, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	oid := parseIDs(ids)
	return o.Collection.DeleteMany(nil, bson.M{"_id": bson.M{"$in": oid}}, opts...)
}

// Upsert 如果存在则修改, 不存在则创建
func (o *CollectionOperator) Upsert(filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	opt := options.Replace().SetUpsert(true)
	opts = append(opts, opt)
	return o.Collection.ReplaceOne(nil, filter, replacement, opts...)
}

// UpdateID 通过ID更新记录
func (o *CollectionOperator) UpdateID(id interface{}, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	oid := parseID(id)
	return o.Collection.UpdateOne(nil, bson.M{"_id": oid}, update, opts...)
}

func (o CollectionOperator) Exist(filter interface{}, opts ...*options.CountOptions) bool {
	count, err := o.Collection.CountDocuments(nil, filter, opts...)
	return err == nil && count > 0
}

// FindID 通过ID查找
func (o *CollectionOperator) FindID(id, result interface{}, opts ...*options.FindOneOptions) error {
	oid := parseID(id)
	return o.Collection.FindOne(nil, bson.M{"_id": oid}, opts...).Decode(result)
}

// FindBatchID 通过ID批量查找
func (o CollectionOperator) FindBatchID(ids []string, result interface{}, opts ...*options.FindOptions) error {
	oid := parseIDs(ids)
	cursor, err := o.Collection.Find(nil, bson.M{"_id": bson.M{"$in": oid}}, opts...)
	if err != nil {
		return err
	}
	return cursor.All(nil, result)
}

// FindMany 查找多条记录
func (o *CollectionOperator) FindMany(filter, result interface{}, opts ...*options.FindOptions) error {
	cursor, err := o.Collection.Find(nil, filter, opts...)
	if err != nil {
		return err
	}
	return cursor.All(nil, result)
}

// FindPage 分页查找
func (o *CollectionOperator) FindPage(filter interface{}, page, size int64, result interface{}, opts ...*options.FindOptions) (int64, error) {
	count, err := o.Collection.CountDocuments(nil, filter)
	if err != nil || count < 1 {
		return count, err
	}

	opt := options.Find().SetSkip((page - 1) * size).SetLimit(size)
	opts = append(opts, opt)

	cursor, err := o.Collection.Find(nil, filter, opts...)
	if err != nil {
		return count, err
	}
	return count, cursor.All(nil, result)
}

// FindAggregate 聚合查找
func (o *CollectionOperator) FindAggregate(pipeline interface{}, result interface{}, opts ...*options.AggregateOptions) error {
	cursor, err := o.Collection.Aggregate(nil, pipeline, opts...)
	if err != nil {
		return err
	}
	return cursor.All(nil, result)
}

func parseID(v interface{}) primitive.ObjectID {
	switch id := v.(type) {
	case primitive.ObjectID:
		return id
	case string:
		oid, _ := primitive.ObjectIDFromHex(id) // 如果错误会返回 primitive.NilObjectID
		return oid
	}
	return primitive.NilObjectID
}

func parseIDs(ids []string) []primitive.ObjectID {
	oid := make([]primitive.ObjectID, 0, len(ids))
	for _, id := range ids {
		hex, _ := primitive.ObjectIDFromHex(id)
		oid = append(oid, hex)
	}
	return oid
}

func toSlice(docs interface{}) (res []interface{}) {
	if docs == nil {
		return nil
	}
	v := reflect.ValueOf(docs)
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		res = append(res, docs)
	}
	size := v.Len()
	for i := 0; i < size; i++ {
		res = append(res, v.Index(i).Interface())
	}
	return
}
