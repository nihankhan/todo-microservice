package service

import (
	"errors"
	"fmt"
	"time"
	"user-service/internal/domain"
	"user-service/internal/repository"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

type UserService struct {
	repo        repository.UserRepository
	jwtSecret   string
	tokenExpiry time.Duration
}

func NewUserService(repo repository.UserRepository, jwtSecret string, tokenExpiry time.Duration) *UserService {
	return &UserService{
		repo:        repo,
		jwtSecret:   jwtSecret,
		tokenExpiry: tokenExpiry,
	}
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	if !checkPasswordHash(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := s.generateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("could not generate token %w", err)
	}

	return token, nil
}

func (s *UserService) GetUser(userID string) (*domain.User, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch user: %w", err)
	}

	return user, nil
}

func (s *UserService) CreateUser(name, username, email, password string) (*domain.User, error) {
	hashedPass, err := hashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("could not hash password: %w", err)
	}
	user := &domain.User{
		ID:           generateUserID(),
		Name:         name,
		Username:     username,
		Email:        email,
		Password:     hashedPass,
		ProfileImage: "",
	}

	createUser, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %w", err)
	}

	return createUser, nil
}

func (s *UserService) UpdateUser(userID, name, username, email, password, profileImage string) (*domain.User, error) {
	user := &domain.User{
		ID:           userID,
		Name:         name,
		Username:     username,
		Email:        email,
		Password:     password,
		ProfileImage: profileImage,
	}

	updateUser, err := s.repo.UpdateUser(user)
	if err != nil {
		return nil, fmt.Errorf("could not update user: %w", err)
	}

	return updateUser, nil
}

func (s *UserService) DeleteUser(userID string) error {
	err := s.repo.DeleteUser(userID)
	if err != nil {
		return fmt.Errorf("could not delete user: %w", err)
	}

	return nil
}

func generateUserID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func (s *UserService) generateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(s.tokenExpiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.jwtSecret))
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
