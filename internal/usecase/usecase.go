package usecase

import (
	"tgbot/internal/handler"
	"tgbot/internal/models"
)

type usecase struct {
	repo RepositoryInterfaces
}

type RepositoryInterfaces interface {
	Download(FileId string) error
	Upload(FileId string) error
}

func New(repo RepositoryInterfaces) handler.UsecaseInterfaces {
	return &usecase{repo: repo}
}

func (u *usecase) Download(photoInfo []models.Photo) error {
	for i := 0; i < len(photoInfo); i++ {
		u.repo.Download(photoInfo[i].FileId)
	}
	return nil
}

func (u *usecase) Upload(photoInfo models.Photo) error {
	return nil
}
