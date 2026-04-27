package service

import (
	"context"
	"errors"

	"github.com/gauas/config-service/model"
	"github.com/google/uuid"
)

func (s *Service) NewConfig(ctx context.Context, req *model.Config) (*model.Config, error) {
	if ok, err := s.repository.Config.Exists(ctx, req); err != nil {
		return nil, err
	} else if ok {
		return nil, errors.New("config already exists")
	}

	if req.Service == "" {
		return nil, errors.New("service is required")
	}

	return s.repository.Config.Create(ctx, &model.Config{
		ID:          uuid.New(),
		Service:     req.Service,
		Environment: req.Environment,
		Config:      req.Config,
	})
}

func (s *Service) UpdateConfig(ctx context.Context, req *model.Config) (*model.Config, error) {
	return s.repository.Config.Update(ctx, req)
}

func (s *Service) GetConfig(ctx context.Context, req *model.Config) (*model.Config, error) {
	if req.Service == "" || req.Environment == "" {
		return nil, errors.New("service and environment are required")
	}

	cfg, err := s.repository.Config.FindOne(ctx, req)
	if err != nil {
		return nil, err
	}
	if cfg == nil {
		return nil, errors.New("config does not exist")
	}

	return cfg, nil
}

func (s *Service) AllConfigs(ctx context.Context, req *model.Config) ([]model.Config, error) {
	items, err := s.repository.Config.GetAll(ctx, req)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) DeleteConfig(ctx context.Context, req *model.Config) error {
	if ok, err := s.repository.Config.Exists(ctx, req); err != nil {
		return err
	} else if !ok {
		return errors.New("config does not exist")
	}

	return s.repository.Config.Delete(ctx, req)
}
