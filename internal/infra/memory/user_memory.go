package memory

import (
	"errors"
	"sync"
	"go-crud/internal/domain"
)

type UserMemoryRepo struct {
	data map[string]domain.User
	mu   sync.Mutex
}

func NewUserMemoryRepo() *UserMemoryRepo {
	return &UserMemoryRepo{
		data: make(map[string]domain.User),
	}
}

func (r *UserMemoryRepo) Create(user domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[user.ID] = user
	return nil
}

func (r *UserMemoryRepo) FindAll() ([]domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var users []domain.User
	for _, u := range r.data {
		users = append(users, u)
	}
	return users, nil
}

func (r *UserMemoryRepo) FindByID(id string) (domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, ok := r.data[id]
	if !ok {
		return domain.User{}, errors.New("not found")
	}
	return user, nil
}

func (r *UserMemoryRepo) Update(user domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[user.ID]; !ok {
		return errors.New("not found")
	}

	r.data[user.ID] = user
	return nil
}

func (r *UserMemoryRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.data, id)
	return nil
}