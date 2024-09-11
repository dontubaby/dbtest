package DB

import (
	"Skillfactory/31-DBpractice/pkg/storage/models"
	"context"
	"log"
)

type DbInterface interface {
	Articles(context.Context) ([]models.Article, error)
	AddArticle(context.Context, models.Article) error
	UpdateArticle(context.Context, models.Article) error
	DeleteArticle(context.Context, models.Article) error
}

func GetAll(ctx context.Context, db DbInterface) ([]models.Article, error) {
	result, err := db.Articles(ctx)
	if err != nil {
		log.Fatalf("Error when GET articles from server: %v\n", err)
		return nil, err
	}

	return result, nil
}

func Add(ctx context.Context, db DbInterface, article models.Article) error {
	err := db.AddArticle(ctx, article)
	if err != nil {
		log.Fatalf("Error when ADD article to database: %v\n", err)
		return err
	}
	return nil
}

func Update(ctx context.Context, db DbInterface, article models.Article) error {
	err := db.UpdateArticle(ctx, article)
	if err != nil {
		log.Fatalf("Error when UPDATE article to database: %v\n", err)
		return err
	}
	return nil
}

func Delete(ctx context.Context, db DbInterface, article models.Article) error {
	err := db.DeleteArticle(ctx, article)
	if err != nil {
		log.Fatalf("Error when DELETE article to database: %v\n", err)
		return err
	}
	return nil
}
