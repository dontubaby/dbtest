// sf db task
package postgres

import (
	models "Skillfactory/31-DBpractice/pkg/storage/models"
	"context"
	"fmt"

	//"database/sql"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	Db *pgxpool.Pool
}

// Вспомогательный метод
func (s *Storage) NewTable(ctx context.Context) error {
	_, err := s.Db.Exec(ctx, `DROP TABLE IF EXISTS articles;

	CREATE TABLE articles( 
		id SERIAL NOT NULL UNIQUE PRIMARY KEY,
		author TEXT,
		title TEXT,
		description TEXT,
		url TEXT,
		urlToImage TEXT,
		publishedAt BIGINT,
		content TEXT
	);`)
	if err != nil {
		log.Fatalf("Error!Cant create new table:  %v\n", err)
		return err
	}
	return nil
}

// вспомогательный метод
func NewDb(ctx context.Context, connString string) (*Storage, error) {
	db, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		log.Fatalf("Cant create new instance of DB: %v\n", err)
	}
	s := Storage{
		Db: db,
	}
	return &s, nil
}

func (s *Storage) Articles(ctx context.Context) ([]models.Article, error) {
	rows, err := s.Db.Query(ctx, `SELECT * FROM articles`)
	if err != nil {
		log.Fatalf("Cant read data from database: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	articles := []models.Article{}

	for rows.Next() {
		article := models.Article{}
		err = rows.Scan(
			&article.ID,
			&article.Author,
			&article.Title,
			&article.Description,
			&article.URL,
			&article.URLToImage,
			&article.PublishedAt,
			&article.Content)
		if err != nil {
			return nil, fmt.Errorf("Unable scan row: %w", err)
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (s *Storage) AddArticle(ctx context.Context, a models.Article) error {
	_, err := s.Db.Exec(ctx, `INSERT INTO articles 
	(id,author,title,description,url,urlToImage,publishedAt,content) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8);`,
		a.ID, a.Author, a.Title, a.Description, a.URL, a.URLToImage, a.PublishedAt, a.Content)
	if err != nil {
		log.Fatalf("Cant add data in database! %v\n", err)
		return err
	}
	return nil
}

func (s *Storage) UpdateArticle(ctx context.Context, a models.Article) error {
	_, err := s.Db.Exec(ctx, `UPDATE articles SET 
	author=$2,title=$3,description=$4,url=$5,urlToImage=$6,publishedAt=$7,content=$8 
	WHERE id=$1;`,
		a.ID, a.Author, a.Title, a.Description, a.URL, a.URLToImage, a.PublishedAt, a.Content)
	if err != nil {
		log.Fatalf("Cant update data in database: %v\n", err)
		return err
	}
	return nil
}

func (s *Storage) DeleteArticle(ctx context.Context, a models.Article) error {
	_, err := s.Db.Exec(ctx, `DELETE FROM articles WHERE id=$1;`, a.ID)
	if err != nil {
		log.Fatalf("Error!Cant write new Article: %v\n", err)
		return err
	}
	return nil
}
