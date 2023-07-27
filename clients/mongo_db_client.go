package clients

import (
	"CurrencyExchangeService/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const clientURI = "mongodb+srv://darryljiang:C2hFE6VfssGmnxKn@cluster0.bpmmfdz.mongodb.net/?retryWrites=true&w=majority"

func NewMongoDBClient(logger *logger.ServiceLogger) *mongo.Client {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(clientURI).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	logger.Info("Create MongoDB Connection")
	dbClient, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	//defer func() {
	//	err = dbClient.Disconnect(context.TODO())
	//	if err != nil {
	//		panic(err)
	//	}
	//}()
	// Send a ping to confirm a successful connection
	if err := dbClient.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	logger.Info("Pinged your deployment. You successfully connected to MongoDB!")
	return dbClient
}
