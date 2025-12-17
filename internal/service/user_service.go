package service

import (
	"context"
	"time"

	db "user-api/db/sqlc"
	"user-api/internal/models"
	"user-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

// Convert DB user → API user + age
func (s *UserService) mapUser(u db.User) models.User {
	dob := time.Time(u.Dob)

	return models.User{
		ID:   int(u.ID),
		Name: u.Name,
		DOB:  dob,
		Age:  CalculateAge(dob),
	}
}

// Create user
func (s *UserService) Create(ctx context.Context, name, dob string) (db.User, error) {
	return s.repo.Create(ctx, name, dob)
}

// Get user by ID (with age)
func (s *UserService) Get(ctx context.Context, id int32) (models.User, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	return s.mapUser(u), nil
}

// List users (with age)
func (s *UserService) List(ctx context.Context) ([]models.User, error) {
	users, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]models.User, 0, len(users))
	for _, u := range users {
		result = append(result, s.mapUser(u))
	}
	return result, nil
}

// Update user
func (s *UserService) Update(ctx context.Context, id int32, name, dob string) (db.User, error) {
	return s.repo.Update(ctx, id, name, dob)
}

// Delete user
func (s *UserService) Delete(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}
