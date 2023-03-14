package service

import (
	"log"

	"github.com/google/uuid"
	"github.com/tharun-d/blog/models"
	"github.com/tharun-d/blog/repository"
)

type Service struct {
	blogRepository repository.IRepository
}

//go:mockery --name IService --disable-version-string
type IService interface {
	InsertBlogDetails(blogData models.Blog) (string, error)
	GetBlogByID(id string) (models.Blog, error)
	GetAllBlogDetails() ([]models.Blog, error)
}

func NewService(blogRepository repository.IRepository) IService {
	return &Service{
		blogRepository: blogRepository,
	}
}

func (s *Service) InsertBlogDetails(blogData models.Blog) (string, error) {

	blogData.Id = uuid.NewString()

	err := s.blogRepository.SaveBlog(blogData)
	if err != nil {
		log.Println(err.Error())
		return blogData.Id, err
	}
	return blogData.Id, nil
}

func (s *Service) GetBlogByID(id string) (models.Blog, error) {

	blogData, err := s.blogRepository.GetBlogByID(id)
	if err != nil {
		log.Println(err.Error())
		return blogData, err
	}

	return blogData, nil
}


func (s *Service) GetAllBlogDetails() ([]models.Blog, error) {

	blogData, err := s.blogRepository.GetAllBlogDetails()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return blogData, nil
}

