# Exercise: Idiomatic Go — service-layer error handling

## Goal

Implement `UserService.Register` with clear error classification and wrapped unexpected failures.

## Spec

```go
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

func (s *UserService) Register(ctx context.Context, email string) (User, error)
```

## Requirements

- Trim whitespace and lowercase the email.
- Return `ErrInvalidEmail` for an obviously invalid address.
- If the repository finds a user, return `ErrEmailTaken`.
- Treat `sql.ErrNoRows` as “email is available.”
- Wrap unexpected repository errors with `%w` and useful context.
- Do not log inside the service.
- Table-driven (or focused) tests using a small repository fake.

## Files

- `service.go` — domain types + `Register`
- `service_test.go` — `fakeRepo` + tests

## Run

```bash
go test ./practice/register/ -v
```

## Rubric

- Distinguishes found, not-found (`sql.ErrNoRows`), and unexpected repo failures.
- Uses explicit sentinel errors and `%w` without losing the cause (`errors.Is`).
- Keeps `UserRepository` narrow to service needs.
- Covers: success (normalize), invalid email, duplicate, lookup failure, create failure.
