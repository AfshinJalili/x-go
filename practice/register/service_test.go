package register

import (
	"context"
	"database/sql"
	"errors"
	"testing"
)

type fakeRepo struct {
	users     map[string]User
	nextID    int64
	findErr   error
	createErr error
}

func newFakeRepo() *fakeRepo {
	return &fakeRepo{
		users:  make(map[string]User),
		nextID: 1,
	}
}

func (r *fakeRepo) FindByEmail(ctx context.Context, email string) (User, error) {
	if r.findErr != nil {
		return User{}, r.findErr
	}
	u, ok := r.users[email]
	if !ok {
		return User{}, sql.ErrNoRows
	}
	return u, nil
}

func (r *fakeRepo) Create(ctx context.Context, email string) (User, error) {
	if r.createErr != nil {
		return User{}, r.createErr
	}
	u := User{ID: r.nextID, Email: email}
	r.nextID++
	r.users[email] = u
	return u, nil
}

func TestRegister_successNormalizesEmail(t *testing.T) {
	repo := newFakeRepo()
	svc := &UserService{repo: repo}

	got, err := svc.Register(context.Background(), "  a@b.com  ")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if got.Email != "a@b.com" {
		t.Fatalf("email =%q, want %q", got.Email, "a@b.com")
	}
}

func TestRegister_invalidEmail(t *testing.T) {
	svc := &UserService{repo: newFakeRepo()}

	_, err := svc.Register(context.Background(), "not-an-email")
	if !errors.Is(err, ErrInvalidEmail) {
		t.Fatalf("err = %v, want %v", err, ErrInvalidEmail)
	}
}

func TestResgister_emailTaken(t *testing.T) {
	repo := newFakeRepo()
	repo.users["a@b.com"] = User{ID: 9, Email: "a@b.com"}
	svc := &UserService{repo: repo}

	_, err := svc.Register(context.Background(), "a@b.com")
	if !errors.Is(err, ErrEmailTaken) {
		t.Fatalf("err = %v, want %v", err, ErrEmailTaken)
	}
}

func TestRegister_lookupFailure(t *testing.T) {
	boom := errors.New("db down")
	repo := newFakeRepo()
	repo.findErr = boom
	svc := &UserService{repo: repo}

	_, err := svc.Register(context.Background(), "a@b.com")
	if !errors.Is(err, boom) {
		t.Fatalf("err = %v, want %v", err, boom)
	}
}

func TestRegister_createFailure(t *testing.T) {
	boom := errors.New("db down")
	repo := newFakeRepo()
	repo.createErr = boom
	svc := &UserService{repo: repo}

	_, err := svc.Register(context.Background(), "a@b.com")
	if !errors.Is(err, boom) {
		t.Fatalf("err = %v, want %v", err, boom)
	}
}
