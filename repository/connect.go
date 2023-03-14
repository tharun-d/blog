package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/tharun-d/blog/core"
	"github.com/tharun-d/blog/models"
)

type Repository struct {
	conn *sql.DB
}

//go:mockery --name IRepository --disable-version-string
type IRepository interface {
	SaveBlog(blogData models.Blog) error
	GetBlogByID(id string) (models.Blog, error)
	GetAllBlogDetails() ([]models.Blog, error)
}

func NewRepository(config core.Config) (IRepository, error) {
	time.Sleep(10 * time.Second)
	var err error

	dbSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	conn, err := sql.Open("postgres", dbSource)

	if err != nil {
		return nil, err
	}
	if err = conn.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connection established")
	return &Repository{conn: conn}, nil
}
