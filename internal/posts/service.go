package posts

import (
	"context"
	"errors"

	repo "github.com/zeronationday/social-network/internal/adapters/postgresql/sqlc"
)

var (
	ErrPostAlreadyExists = errors.New("post already exists")
	ErrPostNotFound      = errors.New("post not found")
)

type Service interface {
	ListPostsByUserID(ctx context.Context, userID int32) ([]repo.Post, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListPostsByUserID(ctx context.Context, userID int32) ([]repo.Post, error) {
	return s.repo.ListPostsByUserID(ctx, userID)
}
