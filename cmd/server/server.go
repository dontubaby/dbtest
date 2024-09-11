package main

import (
	DB "Skillfactory/31-DBpractice/pkg/storage"
	models "Skillfactory/31-DBpractice/pkg/storage/models"
	"Skillfactory/31-DBpractice/pkg/storage/mongo"

	//"Skillfactory/31-DBpractice/pkg/storage/mongo"
	"Skillfactory/31-DBpractice/pkg/storage/postgres"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//тестовые данные для Postgres
	testArticle := models.Article{
		ID:          1,
		Author:      "TestAuthor",
		Title:       "Test",
		Description: "TestDescription",
		URL:         "https://www.theverge.com/2024/9/5/24235776/aqara-voice-mate-h1-voice-control-device",
		URLToImage:  "https://cdn.vox-cdn.com/thumbor/YreIeB9R6de-ProxNAVxd7c76Yg=/533x123:1920x1080/1200x628/filters:focal(1274x716:1275x717)/cdn.vox-cdn.com/uploads/chorus_asset/file/25603955/Aqara_Voice_Mate_H1.png",
		PublishedAt: 1725535011,
		Content:     "Test Content Here",
	}

	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	pwd := os.Getenv("DBPASSWORD")
	//mongoAdress := os.Getenv("MONGOADRESS")

	//**************************ТЕСТИРОВАНИЕ РАБОТЫ POSTGRES**************************

	connString := "postgres://postgres:" + pwd + "@localhost:5432/gotest"

	pool, err := postgres.NewDb(ctx, connString)
	if err != nil {
		log.Fatalf("Cant connect to DB: %v", err)
	}
	defer pool.Db.Close()

	//fmt.Printf("PORT NUMBER: %v\n", port)

	pool.NewTable(ctx)
	err = DB.Add(ctx, pool, testArticle)
	if err != nil {
		log.Fatalf("cant add test data in postgress db: %v", err)
	}

	pdata, err := DB.GetAll(ctx, pool)
	if err != nil {
		log.Fatalf("cant GET test data from postgress db: %v", err)
	}
	fmt.Println(pdata)

	DB.Update(ctx, pool, testArticle)
	//вывод данных после UPDATE
	pdata, err = DB.GetAll(ctx, pool)
	if err != nil {
		log.Fatalf("cant GET test data from postgress db: %v", err)
	}
	fmt.Println(pdata)

	DB.Delete(ctx, pool, testArticle)
	//вывод данных после DELETE
	pdata, err = DB.GetAll(ctx, pool)
	if err != nil {
		log.Fatalf("cant GET test data from postgress db: %v", err)
	}
	fmt.Println(pdata)

	//**************************ТЕСТИРОВАНИЕ РАБОТЫ  ИНТЕРФЕЙСА DB c MongoDB**************************
	mng, err := mongo.NewMongoDb(ctx, mongoAdress)
	if err != nil {
		log.Fatalf("cant connect to MongoDb: %v", err)
	}
	defer mng.Db.Disconnect(ctx)

	err = DB.Add(ctx, mng, testArticle)
	if err != nil {
		log.Fatalf("cant ADD test data in mongodb db: %v", err)
	}

	data, err := DB.GetAll(ctx, mng)
	if err != nil {
		log.Fatalf("cant GET test data in mongodb db: %v", err)
	}
	fmt.Println(data)

	DB.Update(ctx, mng, testArticle)
	//вывод данных после UPDATE
	data, err = DB.GetAll(ctx, mng)
	if err != nil {
		log.Fatalf("cant GET test data in mongodb db: %v", err)
	}
	fmt.Println(data)

	DB.Delete(ctx, mng, testArticle)
	//вывод данных после DELETE
	data, err = DB.GetAll(ctx, mng)
	if err != nil {
		log.Fatalf("cant GET test data in mongodb db: %v", err)
	}
	fmt.Println(data)

}
