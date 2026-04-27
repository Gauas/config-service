package service

import "github.com/gauas/config-service/repository"

type Service struct {
	repository *repository.Registry
}

func New(repo *repository.Registry) *Service {
	return &Service{
		repository: repo,
	}
}
