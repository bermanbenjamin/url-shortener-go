package shortener

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ShortenerRepository interface {
	Create(shortenUrl ShortenURL) (string, error)
	Get(code string) (string, error)
}

type shortenerRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewShortenerRepository(client *mongo.Client) ShortenerRepository {
	collection := client.Database("Shortener").Collection("urls")
	return &shortenerRepository{client: client, collection: collection}
}

func (r *shortenerRepository) Create(shortenUrl ShortenURL) (string, error) {
	_, err := r.collection.InsertOne(context.Background(), shortenUrl)
	if err != nil {
		return "", err
	}
	return shortenUrl.Code, nil
}

func (r *shortenerRepository) Get(code string) (string, error) {
	filter := bson.M{"code": code}
	var shortUrl ShortenURL
	err := r.collection.FindOne(context.Background(), filter).Decode(&shortUrl)
	if err != nil {
		return "", err
	}
	return shortUrl.URL, nil
}
