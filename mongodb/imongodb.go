package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type IMongoDB interface {
	Collection(name string) *mongo.Collection
	Close() error
}
