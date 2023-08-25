package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Media struct {
	UUID             primitive.ObjectID  `bson:"_id,omitempty"`
	Data_type        string              `bson:"data_type"`
	Code_book        string              `bson:"code_book"`
	Description_page string              `bson:"description_page"`
	Data             string              `bson:"data"`
	Metadata         string              `bson:"metadata,omitempty"`
	Datetime         primitive.Timestamp `bson:"datetime,omitempty"`
}

func connectDB() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}

	return client
}

func getCollection(db string, coll string, client *mongo.Client) *mongo.Collection {
	return client.Database(db).Collection(coll)
}

func getDB(db string, client *mongo.Client) *mongo.Database {
	return client.Database(db)
}

func getBucket(db *mongo.Database) *gridfs.Bucket {
	bucket, err := gridfs.NewBucket(db)

	if err != nil {
		panic(err)
	}

	return bucket
}

func insetColl(coll *mongo.Collection, media Media) *mongo.InsertOneResult {
	result, err := coll.InsertOne(context.Background(), media)

	if err != nil {
		panic(err)
	}

	return result
}

func uploadFile(bucket *gridfs.Bucket, path string, name string) primitive.ObjectID {
	file, err := os.Open(path)

	ObjectID, err := bucket.UploadFromStream(name, io.Reader(file))

	if err != nil {
		panic(err)
	}

	return ObjectID
}

func downloadFile(bucket gridfs.Bucket, idFile string, name string) {
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

	fmt.Println("Download realizado com sucesso!")
}

func main() {

}