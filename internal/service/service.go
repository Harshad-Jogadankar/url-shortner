package service

import (
	"context"
	"url-shortner/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Shorten(url string) (string, error) {
	err := isValidUrl(url)

	if err != nil {
		return "", err
	}

	id, err = s.repo.NextId(url)
	if err != nil {
		return "", err
	}

	err = s.repo.Save(context.Background, id, url)
	if err != nil {
		return "", err
	}

	return nil
}

func isValidUrl(url string) error {

}
