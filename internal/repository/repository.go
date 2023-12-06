package repository

import (
	"log"
	"os"
)

type repository struct{}

func New() *repository {
	return &repository{}
}

func (r *repository) Download(FileId string) error {
	log.Println(os.Getenv("TOKEN") + "/getFile?file_id=" + FileId)
	return nil
}

func (r *repository) Upload(FileId string) error { return nil }
