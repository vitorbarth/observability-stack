package service

import (
	"github.com/google/uuid"
	"github.com/vitormbarth/observability-stack/services/account-service/internal/model"
	"github.com/vitormbarth/observability-stack/services/account-service/internal/repository"
)

type AccountService struct {
	repo *repository.MemoryRepository
}

func NewAccountService(r *repository.MemoryRepository) *AccountService {
	return &AccountService{repo: r}
}

func (s *AccountService) Create(name, email string) (model.Account, error) {
	acc := model.Account{
		ID:    uuid.NewString(),
		Name:  name,
		Email: email,
	}
	err := s.repo.Create(acc)
	return acc, err
}

func (s *AccountService) Get(id string) (model.Account, error) {
	return s.repo.Get(id)
}

func (s *AccountService) GetAll() []model.Account {
	return s.repo.GetAll()
}

func (s *AccountService) Delete(id string) error {
	return s.repo.Delete(id)
}
