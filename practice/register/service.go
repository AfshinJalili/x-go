package register

import (
	"context"
	"errors"
	"strings"
	"database/sql"
	"fmt"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
	ErrEmailTaken   = errors.New("email already registered")
)

type User struct {
	ID    int64
	Email string
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (User, error)
	Create(ctx context.Context, email string) (User, error)
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) Register(ctx context.Context, email string) (User, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	if !validEmail(email) {
		return User{}, ErrInvalidEmail
	}

	_, err := s.repo.FindByEmail(ctx, email)
	if err == nil {
		return User{}, ErrEmailTaken 
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return User{}, fmt.Errorf("find user by email: %w", err)
	}

	user, err := s.repo.Create(ctx, email)
	if err != nil {
		return User{}, fmt.Errorf("create user: %w", err) 
	}

	return user, nil
}

func validEmail(email string) bool {
	at := strings.IndexByte(email, '@')
	if at <= 0 || at == len(email)-1 {
		return false
	}

	return strings.IndexByte(email[at+1:], '.') > 0
}