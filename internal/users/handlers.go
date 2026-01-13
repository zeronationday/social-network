package users

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	repo "github.com/zeronationday/social-network/internal/adapters/postgresql/sqlc"
	"github.com/zeronationday/social-network/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers(r.Context())
	if err != nil {
		log.Printf("failed to list users: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, users)
}

func (h *handler) FindUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("failed to convert id to int: %v", err)
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, err := h.service.FindUserByID(r.Context(), int32(id))
	if err != nil {
		log.Printf("failed to find user: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, user)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var tmpUser repo.CreateUserParams
	if err := json.Read(r, &tmpUser); err != nil {
		log.Printf("failed to read user: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(r.Context(), tmpUser)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, user)
}
