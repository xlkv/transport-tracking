package serivce

import "context"

type StatusRepository interface {
	Status(ctx context.Context) (string, error)
}

type StatusService struct {
	repo StatusRepository
}

func NewStatusService(repo StatusRepository) *StatusService {
	return &StatusService{
		repo: repo,
	}
}

func (s *StatusService) Status(ctx context.Context) (string, error) {
	return s.Status(ctx)
}
