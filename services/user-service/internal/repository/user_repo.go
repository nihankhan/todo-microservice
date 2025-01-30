package repository

import (
	"fmt"
	"user-service/internal/domain"
)

type UserRepository interface {
	GetUserByID(userID string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	DeleteUser(userID string) error
}

type InMUserRepo struct {
	users map[string]*domain.User
}

func NewInMUserRepo() *InMUserRepo {
	return &InMUserRepo{
		users: make(map[string]*domain.User),
	}
}

func (repo *InMUserRepo) GetUserByID(userID string) (*domain.User, error) {
	user, exists := repo.users[userID]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (repo *InMUserRepo) GetUserByEmail(email string) (*domain.User, error) {
	for _, user := range repo.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}

func (repo *InMUserRepo) CreateUser(user *domain.User) (*domain.User, error) {
	repo.users[user.ID] = user

	return user, nil
}

func (repo *InMUserRepo) UpdateUser(user *domain.User) (*domain.User, error) {
	repo.users[user.ID] = user

	return user, nil
}

func (repo *InMUserRepo) DeleteUser(userID string) error {
	delete(repo.users, userID)

	return nil
}
