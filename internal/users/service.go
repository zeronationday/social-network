package users

import (
	"context"
	"errors"

	repo "github.com/zeronationday/social-network/internal/adapters/postgresql/sqlc"
	"github.com/zeronationday/social-network/internal/crypto"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)

type Service interface {
	ListUsers(ctx context.Context) ([]repo.User, error)
	FindUserByID(ctx context.Context, id int32) (repo.User, error)
	CreateUser(ctx context.Context, user repo.CreateUserParams) (repo.CreateUserRow, error)
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

func (s *svc) FindUserByID(ctx context.Context, id int32) (repo.User, error) {
	user, err := s.repo.FindUserByID(ctx, id)
	if err != nil {
		return repo.User{}, ErrUserNotFound
	}

	return user, nil
}

func (s *svc) CreateUser(ctx context.Context, user repo.CreateUserParams) (repo.CreateUserRow, error) {
	_, err := s.repo.FindUserByEmail(ctx, user.Email)
	if err == nil {
		return repo.CreateUserRow{}, ErrUserAlreadyExists
	}

	err = crypto.ValidatePasswordStrength(user.Password)
	if err != nil {
		return repo.CreateUserRow{}, err
	}

	hashedPassword, err := crypto.HashPassword(user.Password)
	if err != nil {
		return repo.CreateUserRow{}, err
	}
	user.Password = hashedPassword

	return s.repo.CreateUser(ctx, user)
}
