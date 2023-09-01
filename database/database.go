package database

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"project/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}

	return client
}

func GetCollection(db string, coll string, client *mongo.Client) *mongo.Collection {
	return client.Database(db).Collection(coll)
}

func GetDB(db string, client *mongo.Client) *mongo.Database {
	return client.Database(db)
}

func GetBucket(db *mongo.Database) *gridfs.Bucket {
	bucket, err := gridfs.NewBucket(db)

	if err != nil {
		panic(err)
	}

	return bucket
}

func DownloadFile(bucket gridfs.Bucket, idFile string, name string) {
	id, err := primitive.ObjectIDFromHex(idFile)
	fileBuffer := bytes.NewBuffer(nil)
	if _, err := bucket.DownloadToStream(id, fileBuffer); err != nil {
		panic(err)
	}

	fileContent := fileBuffer.Bytes()

	err = os.WriteFile(name, fileContent, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successful download.")
}

func FindDoc(coll *mongo.Collection, filter bson.D) (models.Media, error) {

	var result models.Media

	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
      fmt.Println("Not found.")
		}
	}

  return result, err
}
