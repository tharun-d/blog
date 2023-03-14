package handlers

import (
	"database/sql"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tharun-d/blog/models"
	svcMocks "github.com/tharun-d/blog/service/mocks"
)

func TestSaveBlog(t *testing.T) {
	fakeURL, _ := url.Parse("localhost/articles")

	tests := map[string]struct {
		req        http.Request
		svc        svcMocks.IService
		wantStatus int
	}{
		"invalid body": {
			req: http.Request{
				URL:  fakeURL,
				Body: io.NopCloser(strings.NewReader("}")),
			},
			wantStatus: http.StatusBadRequest,
		},
		"title empty": {
			req: http.Request{
				URL:  fakeURL,
				Body: io.NopCloser(strings.NewReader("{\"content\":\"OSCAR\",\"author\":\"RajiMouli\"}")),
			},
			wantStatus: http.StatusBadRequest,
		},
		"content empty": {
			req: http.Request{
				URL:  fakeURL,
				Body: io.NopCloser(strings.NewReader("{\"title\":\"RRR\",\"author\":\"RajiMouli\"}")),
			},
			wantStatus: http.StatusBadRequest,
		},
		"author empty": {
			req: http.Request{
				URL:  fakeURL,
				Body: io.NopCloser(strings.NewReader("{\"title\":\"RRR\",\"content\":\"OSCAR\"}")),
			},
			wantStatus: http.StatusBadRequest,
		},
		"success": {
			req: http.Request{
				URL:  fakeURL,
				Body: io.NopCloser(strings.NewReader("{\"title\":\"RRR\", \"content\":\"OSCAR\",\"author\":\"RajiMouli\"}")),
			},
			svc: func() svcMocks.IService {
				var m svcMocks.IService
				m.On("InsertBlogDetails", mock.Anything).Return("id", nil)
				return m
			}(),
			wantStatus: http.StatusOK,
		},
		"server error": {
			req: http.Request{
				URL:  fakeURL,
				Body: io.NopCloser(strings.NewReader("{\"title\":\"RRR\", \"content\":\"OSCAR\",\"author\":\"RajiMouli\"}")),
			},
			svc: func() svcMocks.IService {
				var m svcMocks.IService
				m.On("InsertBlogDetails", mock.Anything).Return("", errors.New("server error"))
				return m
			}(),
			wantStatus: http.StatusInternalServerError,
		},
	}

	for name, test := range tests {

		ha := handler{
			svc: &test.svc,
		}
		rw := httptest.NewRecorder()
		ha.SaveBlog(rw, &test.req)

		assert.Equal(t, test.wantStatus, rw.Code, name)
	}
}

func TestGetBlogByID(t *testing.T) {
	fakeURL, _ := url.Parse("localhost/articles")

	tests := map[string]struct {
		req        http.Request
		svc        svcMocks.IService
		wantStatus int
	}{

		"success": {
			req: http.Request{
				URL: fakeURL,
			},
			svc: func() svcMocks.IService {
				var m svcMocks.IService
				m.On("GetBlogByID", mock.Anything).Return(models.Blog{}, nil)
				return m
			}(),
			wantStatus: http.StatusOK,
		},
		"no data error": {
			req: http.Request{
				URL: fakeURL,
			},
			svc: func() svcMocks.IService {
				var m svcMocks.IService
				m.On("GetBlogByID", mock.Anything).Return(models.Blog{}, sql.ErrNoRows)
				return m
			}(),
			wantStatus: http.StatusNotFound,
		},
		"server error": {
			req: http.Request{
				URL: fakeURL,
			},
			svc: func() svcMocks.IService {
				var m svcMocks.IService
				m.On("GetBlogByID", mock.Anything).Return(models.Blog{}, errors.New("server error"))
				return m
			}(),
			wantStatus: http.StatusInternalServerError,
		},
	}

	for name, test := range tests {

		ha := handler{
			svc: &test.svc,
		}
		rw := httptest.NewRecorder()
		ha.GetBlogByID(rw, &test.req)

		assert.Equal(t, test.wantStatus, rw.Code, name)
	}
}

func TestGetAll(t *testing.T) {
	fakeURL, _ := url.Parse("localhost/articles")

	tests := map[string]struct {
		req        http.Request
		svc        svcMocks.IService
		wantStatus int
	}{

		"success": {
			req: http.Request{
				URL: fakeURL,
			},
			svc: func() svcMocks.IService {
				var m svcMocks.IService
				m.On("GetAllBlogDetails").Return([]models.Blog{}, nil)
				return m
			}(),
			wantStatus: http.StatusOK,
		},
		"server error": {
			req: http.Request{
				URL: fakeURL,
			},
			svc: func() svcMocks.IService {
				var m svcMocks.IService
				m.On("GetAllBlogDetails").Return([]models.Blog{}, errors.New("server error"))
				return m
			}(),
			wantStatus: http.StatusInternalServerError,
		},
	}

	for name, test := range tests {

		ha := handler{
			svc: &test.svc,
		}
		rw := httptest.NewRecorder()
		ha.GetAll(rw, &test.req)

		assert.Equal(t, test.wantStatus, rw.Code, name)
	}
}

func TestNewService(t *testing.T) {
	tests := map[string]struct {
		want handler
		svc  *svcMocks.IService
	}{
		"success": {
			svc: func() *svcMocks.IService {
				var m svcMocks.IService
				return &m
			}(),
			want: handler{
				svc: func() *svcMocks.IService {
					var m svcMocks.IService
					return &m
				}(),
			},
		},
	}

	for name, tt := range tests {

		got := NewHandlers(tt.svc)
		assert.Equal(t, tt.want, got, name)
	}
}
