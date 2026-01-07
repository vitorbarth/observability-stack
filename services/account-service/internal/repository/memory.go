package repository

import (
	"errors"
	"sync"

	"github.com/vitormbarth/observability-stack/services/account-service/internal/model"
)

var ErrNotFound = errors.New("account not found")

type MemoryRepository struct {
	mu       sync.RWMutex
	accounts map[string]model.Account
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		accounts: make(map[string]model.Account),
	}
}

func (r *MemoryRepository) Create(a model.Account) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.accounts[a.ID] = a
	return nil
}

func (r *MemoryRepository) Get(id string) (model.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	acc, ok := r.accounts[id]
	if !ok {
		return model.Account{}, ErrNotFound
	}

	return acc, nil
}

func (r *MemoryRepository) GetAll() []model.Account {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]model.Account, 0, len(r.accounts))
	for _, acc := range r.accounts {
		list = append(list, acc)
	}
	return list
}

func (r *MemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.accounts[id]; !ok {
		return ErrNotFound
	}
	delete(r.accounts, id)
	return nil
}
