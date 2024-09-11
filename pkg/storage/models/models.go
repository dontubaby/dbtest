package models

import "context"

type DbInterface interface {
	Articles(context.Context) ([]Article, error)
	AddArticle(context.Context, Article) error
	UpdateArticle(context.Context, Article) error
	DeleteArticle(context.Context, Article) error
}

type Source struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Article struct {
	ID          int    `db:"id"`
	Author      string `db:"author"`
	Title       string `db:"title"`
	Description string `db:"description"`
	URL         string `db:"url"`
	URLToImage  string `db:"urlToImage"`
	PublishedAt int64  `db:"publishedAt"`
	Content     string `db:"content"`
}

type Results struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Search struct {
	SearchKey  string
	NextPage   int
	TotalPages int
	Results    Results
}
