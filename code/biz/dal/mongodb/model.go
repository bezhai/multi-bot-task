package mongodb

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/bezhai/multi-bot-task/biz/utils/langx/slicex"
)

type MongoCollection[T any] struct {
	c *mongo.Collection
}

func NewCollection[T any](c *mongo.Collection) *MongoCollection[T] {
	return &MongoCollection[T]{
		c: c,
	}
}

func (m *MongoCollection[T]) FindOne(ctx context.Context, filter interface{},
	opts ...*options.FindOneOptions) (*T, error) {
	resp := m.c.FindOne(ctx, filter, opts...)
	var t T
	err := resp.Err()
	if err != nil {
		return nil, err
	}

	err = resp.Decode(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (m *MongoCollection[T]) Find(ctx context.Context, filter interface{},
	opts ...*options.FindOptions) ([]*T, error) {
	cursor, err := m.c.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	var res []*T
	for cursor.Next(ctx) {
		var t T
		err = cursor.Decode(&t)
		if err != nil {
			return nil, err
		}
		res = append(res, &t)
	}

	return res, nil
}

func (m *MongoCollection[T]) InsertOne(ctx context.Context, doc *T,
	opts ...*options.InsertOneOptions) error {
	_, err := m.c.InsertOne(ctx, doc, opts...)
	return err
}

func (m *MongoCollection[T]) InsertMany(ctx context.Context, docs []*T,
	opts ...*options.InsertManyOptions) error {
	_, err := m.c.InsertMany(ctx, slicex.Map(docs, func(from *T) interface{} {
		return from
	}), opts...)
	return err
}

func (m *MongoCollection[T]) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {
	_, err := m.c.UpdateOne(ctx, filter, bson.M{"$set": update}, opts...)
	return err
}

func (m *MongoCollection[T]) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {
	_, err := m.c.UpdateMany(ctx, filter, bson.M{"$set": update}, opts...)
	return err
}

func (m *MongoCollection[T]) DeleteMany(ctx context.Context, filter bson.M, opts ...*options.DeleteOptions) error {
	deleteRes, err := m.c.DeleteMany(ctx, filter, opts...)
	if err != nil {
		hlog.CtxWarnf(ctx, "DeleteMany err:%v", err)
	} else {
		hlog.CtxInfof(ctx, "DeleteMany success, deleteNum:%d", deleteRes.DeletedCount)
	}
	return err
}

func (m *MongoCollection[T]) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	count, err := m.c.CountDocuments(ctx, filter, opts...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

type MongoCollectionAggregate[K any] struct {
	c        *mongo.Collection
	pipeLine []bson.M
}

func NewAggregate[T any, K any](mongoCollection *MongoCollection[T]) *MongoCollectionAggregate[K] {
	return &MongoCollectionAggregate[K]{
		c: mongoCollection.c,
	}
}

func (m *MongoCollectionAggregate[K]) Match(filter bson.M) *MongoCollectionAggregate[K] {
	m.pipeLine = append(m.pipeLine, bson.M{"$match": filter})
	return m
}

func (m *MongoCollectionAggregate[K]) Group(field string) *MongoCollectionAggregate[K] {
	m.pipeLine = append(m.pipeLine, bson.M{"$group": bson.M{"_id": field}})
	return m
}

func (m *MongoCollectionAggregate[K]) MultiGroup(fields []string) *MongoCollectionAggregate[K] {
	ids := bson.M{}
	for _, field := range fields {
		ids[field] = "$" + field
	}
	m.pipeLine = append(m.pipeLine, bson.M{"$group": bson.M{"_id": ids}})
	return m
}

func (m *MongoCollectionAggregate[K]) Sort(field string, asc bool) *MongoCollectionAggregate[K] {
	sort := 1
	if !asc {
		sort = -1
	}
	m.pipeLine = append(m.pipeLine, bson.M{"$sort": bson.M{field: sort}})
	return m
}

func (m *MongoCollectionAggregate[K]) Limit(limit int64) *MongoCollectionAggregate[K] {
	m.pipeLine = append(m.pipeLine, bson.M{"$limit": limit})
	return m
}

func (m *MongoCollectionAggregate[K]) Unwind(field string) *MongoCollectionAggregate[K] {
	m.pipeLine = append(m.pipeLine, bson.M{"$unwind": "$" + field})
	return m
}

func (m *MongoCollectionAggregate[K]) Project(fields bson.M) *MongoCollectionAggregate[K] {
	m.pipeLine = append(m.pipeLine, bson.M{"$project": fields})
	return m
}

func (m *MongoCollectionAggregate[K]) Lookup(from, localField, foreignField, as string) *MongoCollectionAggregate[K] {
	m.pipeLine = append(m.pipeLine, bson.M{"$lookup": bson.M{
		"from":         from,
		"localField":   localField,
		"foreignField": foreignField,
		"as":           as,
	}})
	return m
}

func (m *MongoCollectionAggregate[K]) Find(ctx context.Context) ([]*K, error) {
	cursor, err := m.c.Aggregate(ctx, m.pipeLine)
	if err != nil {
		hlog.CtxWarnf(ctx, "Aggregate err:%v", err)
		return nil, err
	}
	var res []*K
	for cursor.Next(ctx) {
		var t K
		err = cursor.Decode(&t)
		if err != nil {
			hlog.CtxWarnf(ctx, "Aggregate err:%v", err)
			return nil, err
		}
		res = append(res, &t)
	}

	return res, nil
}
