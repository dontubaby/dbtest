package mongo

import (
	models "Skillfactory/31-DBpractice/pkg/storage/models"
	"context"
	"fmt"
	"log"

	//"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	Db *mongo.Client
}

const (
	dbname         = "gotest"
	CollectionName = "Articles"
)

func NewMongoDb(ctx context.Context, adress string) (*Storage, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + adress + ":27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Cant create mongo client!", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Not answer from MongoDb server!", err)
	}

	s := Storage{
		Db: client,
	}
	return &s, nil
}

func (s *Storage) Articles(ctx context.Context) ([]models.Article, error) {
	collection := s.Db.Database(dbname).Collection(CollectionName)
	cursor, err := collection.Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	var articles []models.Article
	for cursor.Next(ctx) {

		if err = cursor.All(ctx, &articles); err != nil {
			log.Fatalf("Decoding error MongoDb: %v", err)
			return nil, err
		}
		for _, article := range articles {
			fmt.Printf("%+v\n", article)
		}
	}
	return articles, nil
}

func (s *Storage) AddArticle(ctx context.Context, article models.Article) error {
	collection := s.Db.Database(dbname).Collection(CollectionName)
	_, err := collection.InsertOne(ctx, article)
	if err != nil {
		log.Fatalf("Cant add data in MongoDb: %v", err)
		return err
	}
	return nil
}

func (s *Storage) UpdateArticle(ctx context.Context, article models.Article) error {
	collection := s.Db.Database(dbname).Collection(CollectionName)

	filter := bson.D{{"title", article.Title}}

	update := bson.D{
		{"$set", bson.D{
			{"title", "testing update here"},
		}},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatalf("Cant update data in MongoDb: %v", err)
		return err
	}
	return nil
}

func (s *Storage) DeleteArticle(ctx context.Context, article models.Article) error {
	collection := s.Db.Database(dbname).Collection(CollectionName)
	filter := bson.M{}
	_, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatalf("Cant delete data from MongoDb: %v", err)
		return err
	}
	return nil
}
