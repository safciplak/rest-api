package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rest-api/internal/comment"
	log "github.com/sirupsen/logrus"
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

// LoggingMiddleware - adds middleware around endpoints
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"Method": r.Method,
				"Path":   r.URL.Path,
			}).Info("Request Received")
		next.ServeHTTP(w, r)
	})
}

// Response - an object to store responses from our API
type Response struct {
	Message string
	Error   string
}

// BasicAuth - a handy middleware function that will provide basic auth around specific endpoints
func BasicAuth(original func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("basic auth endpoint hit")
		user, pass, ok := r.BasicAuth()
		if user == "admin" && pass == "pass" && ok {
			original(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			sendErrorResponse(w, "not auhorized", errors.New("not authorized"))
		}
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	log.Info("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/comment", BasicAuth(h.PostComment)).Methods(http.MethodPost)
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/comment/{id}", BasicAuth(h.UpdateComment)).Methods(http.MethodPut)
	h.Router.HandleFunc("/api/comment/{id}", BasicAuth(h.DeleteComment)).Methods(http.MethodDelete)

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		if err := sendOkResponse(w, Response{Message: "I am Alive"}); err != nil {
			panic(err)
		}
	})
}

func sendOkResponse(w http.ResponseWriter, res interface{}) error {
	w.Header().Set("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
