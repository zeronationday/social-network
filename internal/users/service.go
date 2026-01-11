package users

import (
	"context"

	repo "github.com/zeronationday/social-network/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListUsers(ctx context.Context) ([]repo.User, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListUsers(ctx context.Context) ([]repo.User, error) {
	return s.repo.ListUsers(ctx)
}
