package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rest-api/internal/comment"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// NewHandler - return a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// Response - an object to store responses from our API
type Response struct {
	Message string
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods(http.MethodPost)
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods(http.MethodPut)
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods(http.MethodDelete)

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am Alive"}); err != nil {
			panic(err)
		}
	})
}

// GetComment - retrievea comment by ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse UINT from ID")
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		fmt.Fprintf(w, "Error Retrieving comment by ID")
		return
	}

	w.Header().Set("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

// GetAllComments - retrieves all comments from the comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.GetAllComments()

	if err != nil {
		fmt.Fprintf(w, "Failed to retrieve all comments")
	}

	w.Header().Set("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

// PostComment - adds a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintf(w, "Failed to decode JSON Body")
	}

	comment, err := h.Service.PostComment(comment)

	if err != nil {
		fmt.Fprintf(w, "Failed to post new comment")
	}

	w.Header().Set("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

// UpdateComment - updates a comment by ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintf(w, "Failed to decode JSON Body")
	}

	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse UINT from ID")
	}

	comment, err = h.Service.UpdateComment(uint(commentID), comment)
	if err != nil {
		fmt.Fprintf(w, "Failed to update comment")
	}

	w.Header().Set("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

// DeleteComment - deletes a comment by ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse UINT from ID")
	}

	err = h.Service.DeleteComment(uint(commentID))

	if err != nil {
		fmt.Fprintf(w, "Failed to delete comment by comment ID")
	}

	w.Header().Set("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response{Message: "Comment successfuly deleted"}); err != nil {
		panic(err)
	}
}
