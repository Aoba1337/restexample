package repository

import "github.com/aoba1337/restexample/models"

type DummyRepository struct {
}

func (m *DummyRepository) GetUsers() (*[]models.User, error) {
	users := &[]models.User{{
		Id:    1,
		Age:   28,
		Name:  "sanya",
		Debts: []models.Debt{{Id: 1, UserId: 1, Value: 100}},
	}}

	return users, nil
}
