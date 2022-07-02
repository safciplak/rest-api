package http

import (
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
		fmt.Fprintf(w, "I am alive!")
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

	fmt.Fprintf(w, "Comment: %+v", comment)
}

// GetAllComments - retrieves all comments from the comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.GetAllComments()

	if err != nil {
		fmt.Fprintf(w, "Failed to retrieve all comments")
	}

	fmt.Fprintf(w, "Comment: %+v", comment)
}

// PostComment - adds a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.PostComment(comment.Comment{
		Slug: "/",
	})

	if err != nil {
		fmt.Fprintf(w, "Failed to post new comment")
	}

	fmt.Fprintf(w, "%+v", comment)
}

// UpdateComment - updates a comment by ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.UpdateComment(1, comment.Comment{
		Slug: "/new",
	})

	if err != nil {
		fmt.Fprintf(w, "Failed to updae comment")
	}

	fmt.Fprintf(w, "%+v", comment)
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

	fmt.Fprintf(w, "Succesfuly deleted Comment")
}
