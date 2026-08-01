package service

import (
	"context"
	"fmt"
	pkgurl "net/url"
	"url-shortner/internal/encoder"
	"url-shortner/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Shorten(url string) (string, error) {
	ok, err := isValidUrl(url)

	if !ok {
		return "", err
	}

	id, err = s.repo.NextId(url)
	if err != nil {
		return "", err
	}

	shortenedUrl, err := encoder.Encode(id)
	if err != nil {
		return "", err
	}

	err = s.repo.Save(context.Background, url, shortenedUrl)
	if err != nil {
		return "", err
	}

	return shortenedUrl, nil
}

func isValidUrl(url string) (bol, error) {
	u, err := pkgurl.ParseRequestURI(url)
	if err != nil {
		return false, err
	}

	if u.Host == "" {
		return false, fmt.Errorf("Empty host, enter valid url")
	}

	return true, nil
}
