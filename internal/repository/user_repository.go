package repository

import "go-crud/internal/domain"

type UserRepository interface {
	Create(user domain.User) error
	FindAll() ([]domain.User, error)
	FindByID(id string) (domain.User, error)
	Update(user domain.User) error
	Delete(id string) error
}