package repository

import (
	"errors"

	"github.com/abdinep/rent-a-car-grpc/user-service/internal/service"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(username, password, email string) error
	FindUserByUsername(username string) (*service.User, error)
	// FindUserByID(userID string) (*service.User, error)
}

type userRepository struct {
	users map[string]*service.User
}

func NewUserRepository() UserRepository {
	return &userRepository{users: make(map[string]*service.User)}
}
func (r *userRepository) CreateUser(username, password, email string) error {
	for _, user := range r.users {
		if user.Email == email {
			// log.Fatal("user already exist")
			return errors.New("username already exists")
		}
	}
	userID := uuid.New().String()
	user := &service.User{
		ID:       userID,
		Username: username,
		Password: password,
		Email:    email,
	}
	r.users[userID] = user
	return nil
}
func (r *userRepository) FindUserByUsername(username string) (*service.User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
