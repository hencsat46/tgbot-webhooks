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
	Create(string, string) error
	ReadPassword(string) (string, error)
	GetCount(int64) (int64, error)
}

func New(repo RepositoryInterfaces) handler.UsecaseInterfaces {
	return &usecase{repo: repo}
}

func (u *usecase) Download(photoInfo []models.Photo) error {
	u.repo.Download(photoInfo[0].FileId)

	return nil
}

func (u *usecase) Commands(command string) (int, error) {
	switch command {
	case "/profile":
		return 1, nil
	case "/addpicture":
		return 2, nil
	case "/getPictures":
		return 3, nil
	}

	return -1, nil
}

func (u *usecase) Upload(photoInfo models.Photo) error {
	return nil
}
