package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tharun-d/blog/models"
	"github.com/tharun-d/blog/response"
	"github.com/tharun-d/blog/service"
)

type handler struct {
	svc service.IService
}

func NewHandlers(svc service.IService) handler {
	return handler{svc: svc}
}

func (h *handler) SaveBlog(w http.ResponseWriter, r *http.Request) {

	var blogData models.Blog

	err := json.NewDecoder(r.Body).Decode(&blogData)
	if err != nil {
		log.Println(err)
		response.Response(w, http.StatusBadRequest, "invalid request body", nil)
		return
	}
	if blogData.Title == "" {
		response.Response(w, http.StatusBadRequest, "title is empty", nil)
		return
	}
	if blogData.Content == "" {
		response.Response(w, http.StatusBadRequest, "content is empty", nil)
		return
	}
	if blogData.Author == "" {
		response.Response(w, http.StatusBadRequest, "author is empty", nil)
		return
	}
	blogID, err := h.svc.InsertBlogDetails(blogData)

	if err != nil {
		log.Println(err)
		response.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.Response(w, http.StatusOK, "Success", models.Blog{
		Id: blogID,
	})
}

func (h *handler) GetBlogByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	blogData, err := h.svc.GetBlogByID(id)

	if err != nil {
		log.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			response.Response(w, http.StatusNotFound, "no data found", nil)
		}
		response.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.Response(w, http.StatusOK, "Success", blogData)

}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {

	blogData, err := h.svc.GetAllBlogDetails()

	if err != nil {
		log.Println(err)
		response.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.Response(w, http.StatusOK, "Success", blogData)
}
