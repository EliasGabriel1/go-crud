package usecase

import (
	"github.com/google/uuid"
	"go-crud/internal/domain"
	"go-crud/internal/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: r}
}

func (u *UserUsecase) Create(email, password string) error {
	user := domain.User{
		ID:       uuid.New().String(),
		Email:    email,
		Password: password,
	}
	return u.repo.Create(user)
}

func (u *UserUsecase) List() ([]domain.User, error) {
	return u.repo.FindAll()
}