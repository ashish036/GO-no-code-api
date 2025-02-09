package utils

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	repoUtils "github.com/ashish036/GO-no-code-api/utils"
)

const (
	mongoDbConnectionKey = "MONGODB_CONNECTION_STRING"
	logModule            = "MONGODB_UTILS"
)

var (
	mongoClient     *mongo.Client = nil
	mongoCLientOnce sync.Once
)

func initMongoClient() {
	// Fetch the MongoDB connection string from the environment variables
	dB_URL, status := repoUtils.GetValueFromEnv(mongoDbConnectionKey)
	if status.IsError() {
		repoUtils.GetLogHelper(context.TODO(), logModule, http.StatusInternalServerError).Error("Failed to get MongoDB connection string from environment variables")
		return
	}

	// Create a new MongoClient
	tlsConfig := &tls.Config{InsecureSkipVerify: false}
	clientOptions := options.Client().ApplyURI(dB_URL).SetMaxPoolSize(50).SetTLSConfig(tlsConfig)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		repoUtils.GetLogHelper(context.TODO(), logModule, http.StatusInternalServerError).Error("Failed to connect to MongoDB")
		return
	}

	// Check the connection
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		repoUtils.GetLogHelper(context.TODO(), logModule, http.StatusInternalServerError).Error("Failed to ping MongoDB")
	}

	// Assign the value
	mongoClient = client
}

func GetDBCollectionConnection(dbName string, collection string) (*mongo.Collection, repoUtils.Status) {
	mongoCLientOnce.Do(func() {
		initMongoClient()
	})

	if mongoClient == nil {
		return nil, repoUtils.Status{
			Code:    http.StatusInternalServerError,
			Message: "Failed to initialize MongoDB client",
			Error:   errors.New("failed to initialize MongoDB client"),
		}
	}

	if dbName == "" || collection == "" {
		return nil, repoUtils.Status{
			Code:    http.StatusBadRequest,
			Message: "Database and collection name cannot be empty",
			Error:   errors.New("database and collection name cannot be empty"),
		}
	}

	return mongoClient.Database(dbName).Collection(collection), repoUtils.Status{}
}
