package parser

import (
	"context"
	"os"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	connectionUrl string
	credential    options.Credential
	clientOps     *options.ClientOptions
	connection    *mongo.Client
}

func NewClient() (*Client, error) {

	connectionUrl := "mongodb://localhost:27017"

	credential := options.Credential{
		Username: os.Getenv("MONGODB_USERNAME"),
		Password: os.Getenv("MONGODB_PASS"),
		//PasswordSet: true,
	}

	clientOpts := options.Client().ApplyURI(connectionUrl).SetAuth(credential)

	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		return nil, err
	}

	c := &Client{
		connectionUrl: connectionUrl,
		credential:    credential,
		clientOps:     clientOpts,
		connection:    client,
	}

	return c, nil
}

func (c Client) Database(dbName string) *mongo.Database {
	return c.connection.Database(dbName)
}
