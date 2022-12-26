package repository

import "github.com/aoba1337/restexample/models"

type IRepository interface {
	GetUsers() (*[]models.User, error)
}

func CreateRepository() (IRepository, error) {
	return &DummyRepository{}, nil
}
