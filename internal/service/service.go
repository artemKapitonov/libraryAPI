package service

type Service struct{}

type Repository interface {
}

func New(repo Repository) *Service {
	return &Service{}
}
