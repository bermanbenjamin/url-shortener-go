package config

import (
	"context"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Config struct {
	Client    *mongo.Client
	Variables *viper.Viper
}

func NewConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	uri := viper.GetString("database.uri")

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		Client:    client,
		Variables: viper.GetViper(),
	}
}

func (c *Config) Close() error {
	return c.Client.Disconnect(context.Background())
}
