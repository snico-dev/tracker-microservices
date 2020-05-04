package dbcontext

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DbContext database context
type DbContext struct {
	client *mongo.Client
}

//NewContext database context
func NewContext() DbContext {
	return DbContext{}
}

//Connect method to connect in database
func (ctx *DbContext) Connect() error {

	connectionString := os.Getenv("MONGO_CONNECTION_STRING")

	if connectionString == "" {
		connectionString = "mongodb://localhost:27017"
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	ctx.client = client
	fmt.Println("Connected to MongoDB!")

	return err
}

func getCollectionName(myvar interface{}) string {
	t := reflect.TypeOf(myvar)

	if t.Kind() == reflect.Ptr {
		return strings.ToLower(t.Elem().Name() + "s")
	}

	return strings.ToLower(t.Name() + "s")
}

//GetCollection get collection
func (ctx *DbContext) GetCollection(structInstance interface{}) (*mongo.Collection, error) {
	collection := ctx.client.Database("trips").Collection(getCollectionName(structInstance))

	return collection, nil
}

//GetCtx context
func (ctx *DbContext) GetCtx(structInstance interface{}) (*mongo.Client, error) {

	return ctx.client, nil
}
