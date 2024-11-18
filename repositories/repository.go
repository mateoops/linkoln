package repositories

import (
	"context"

	"github.com/mateoops/linkoln/models"

	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShortRepo interface {
	CreateShort(ctx context.Context, short models.Short) (string, error)
	GetByShortUrl(ctx context.Context, shortUrl string) models.Short
}

type MongoShortRepo struct {
	collection mongo.Collection
}

func NewMongoShortRepo() *MongoShortRepo {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("LINKOLN_DB_URI")
	databaseName := os.Getenv("LINKOLN_DB_NAME")
	collectionName := os.Getenv("LINKOLN_DB_COLLECTION")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(databaseName).Collection(collectionName)

	return &MongoShortRepo{
		collection: *collection,
	}
}

func (mr *MongoShortRepo) CreateShort(ctx context.Context, short models.Short) (string, error) {
	_, err := mr.collection.InsertOne(context.TODO(), short)
	return short.ShortUrl, err
}

func (mr *MongoShortRepo) GetByShortUrl(ctx context.Context, shortUrl string) models.Short {

	var result models.Short
	mr.collection.FindOne(context.TODO(), bson.D{{"shorturl", shortUrl}}).Decode(&result)

	return result
}
