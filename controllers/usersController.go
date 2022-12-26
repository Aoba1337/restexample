package controllers

import (
	"github.com/aoba1337/restexample/repository"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	repository repository.IRepository
}

func (c *UsersController) GetUsers(ctx *gin.Context) error {
	users, err := c.repository.GetUsers()
	if err != nil {
		return err
	}

	ctx.JSON(200, users)
	return nil
}

func CreateUserController() (*UsersController, error) {
	repo, err := repository.CreateRepository()
	if err != nil {
		return nil, err
	}

	return &UsersController{repository: repo}, nil
}
