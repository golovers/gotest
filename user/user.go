package user

import "context"

type (
	repository interface {
		Get(ctx context.Context, id string) (*User, error)
		Create(ctx context.Context, user *User) error
		Update(ctx context.Context, user *User) error
	}

	Service struct {
		repo repository
	}
)

func NewService(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Register(ctx context.Context, user *User) error {
	_, err := s.repo.Get(ctx, user.ID)
	if err != nil && err != ErrNotFound {
		return err
	}
	if err == ErrNotFound {
		return s.repo.Create(ctx, user)
	}
	return s.repo.Update(ctx, user)
}
