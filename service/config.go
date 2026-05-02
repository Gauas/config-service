package service

import (
	"context"
	"net/http"

	"github.com/gauas/config-service/model"
	"github.com/gauas/config-service/packages/response"
	"github.com/google/uuid"
)

func (s *Service) NewConfig(ctx context.Context, req *model.Config) (*model.Config, error) {
	if req.Service == "" {
		return nil, response.NewError(http.StatusBadRequest, "service is required")
	}

	if ok, err := s.repository.Config.Exists(ctx, "service = ? AND environment = ?", req.Service, req.Environment); err != nil {
		return nil, err
	} else if ok {
		return nil, response.NewError(http.StatusConflict, "config already exists")
	}

	return s.repository.Config.Create(ctx, &model.Config{
		ID:          uuid.New(),
		Service:     req.Service,
		Environment: req.Environment,
		Config:      req.Config,
	})
}

func (s *Service) UpdateConfig(ctx context.Context, req *model.Config) (*model.Config, error) {
	existing, err := s.repository.Config.FindOne(ctx, "service = ? AND environment = ?", req.Service, req.Environment)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, response.ErrorNotFound
	}

	if err := existing.Merge(req.Config); err != nil {
		return nil, err
	}

	return s.repository.Config.Update(ctx, existing)
}

func (s *Service) GetConfig(ctx context.Context, req *model.Config) (*model.Config, error) {
	if req.Service == "" || req.Environment == "" {
		return nil, response.NewError(http.StatusBadRequest, "service and environment are required")
	}

	cfg, err := s.repository.Config.FindOne(ctx, req)
	if err != nil {
		return nil, err
	}
	if cfg == nil {
		return nil, response.ErrorNotFound
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
		return response.ErrorNotFound
	}

	return s.repository.Config.Delete(ctx, req)
}
