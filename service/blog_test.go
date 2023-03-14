package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tharun-d/blog/models"
	repositoryMocks "github.com/tharun-d/blog/repository/mocks"
)

func TestInsertBlogDetails(t *testing.T) {
	tests := map[string]struct {
		blogRepository *repositoryMocks.IRepository
		wantError      error
	}{
		"success": {
			blogRepository: func() *repositoryMocks.IRepository {
				var m repositoryMocks.IRepository
				m.On("SaveBlog", mock.Anything).Return(nil)
				return &m
			}(),
			wantError: nil,
		},
		"failure": {
			blogRepository: func() *repositoryMocks.IRepository {
				var m repositoryMocks.IRepository
				m.On("SaveBlog", mock.Anything).Return(errors.New("system error"))
				return &m
			}(),
			wantError: errors.New("system error"),
		},
	}

	for name, tt := range tests {
		svc := Service{
			blogRepository: tt.blogRepository,
		}

		_, err := svc.InsertBlogDetails(models.Blog{})
		assert.Equal(t, tt.wantError, err, name)
	}
}

func TestGetAllBlogDetails(t *testing.T) {
	tests := map[string]struct {
		blogRepository *repositoryMocks.IRepository
		want           []models.Blog
		wantError      error
	}{
		"success": {
			blogRepository: func() *repositoryMocks.IRepository {
				var m repositoryMocks.IRepository
				m.On("GetAllBlogDetails", mock.Anything).Return([]models.Blog{}, nil)
				return &m
			}(),
			want:      []models.Blog{},
			wantError: nil,
		},
		"failure": {
			blogRepository: func() *repositoryMocks.IRepository {
				var m repositoryMocks.IRepository
				m.On("GetAllBlogDetails", mock.Anything).Return([]models.Blog{}, errors.New("system error"))
				return &m
			}(),
			wantError: errors.New("system error"),
		},
	}

	for name, tt := range tests {
		svc := Service{
			blogRepository: tt.blogRepository,
		}

		got, err := svc.GetAllBlogDetails()
		assert.Equal(t, tt.want, got, name)
		assert.Equal(t, tt.wantError, err, name)
	}
}

func TestGetBlogByID(t *testing.T) {
	tests := map[string]struct {
		blogRepository *repositoryMocks.IRepository
		want           models.Blog
		wantError      error
	}{
		"success": {
			blogRepository: func() *repositoryMocks.IRepository {
				var m repositoryMocks.IRepository
				m.On("GetBlogByID", mock.Anything).Return(models.Blog{}, nil)
				return &m
			}(),
			want:      models.Blog{},
			wantError: nil,
		},
		"failure": {
			blogRepository: func() *repositoryMocks.IRepository {
				var m repositoryMocks.IRepository
				m.On("GetBlogByID", mock.Anything).Return(models.Blog{}, errors.New("system error"))
				return &m
			}(),
			wantError: errors.New("system error"),
		},
	}

	for name, tt := range tests {
		svc := Service{
			blogRepository: tt.blogRepository,
		}

		got, err := svc.GetBlogByID("dda8121e-5f4e-4ebe-a0a1-598768508e2b")
		assert.Equal(t, tt.want, got, name)
		assert.Equal(t, tt.wantError, err, name)
	}
}

func TestNewService(t *testing.T) {
	tests := map[string]struct {
		want           IService
		blogRepository *repositoryMocks.IRepository
	}{
		"success": {
			want: &Service{
				blogRepository: func() *repositoryMocks.IRepository {
					var m repositoryMocks.IRepository
					return &m
				}(),
			},
			blogRepository: func() *repositoryMocks.IRepository {
				var m repositoryMocks.IRepository
				return &m
			}(),
		},
	}

	for name, tt := range tests {

		got := NewService(tt.blogRepository)
		assert.Equal(t, tt.want, got, name)
	}
}
