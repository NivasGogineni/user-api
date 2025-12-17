package repository

import (
	"context"
	"time"

	db "user-api/db/sqlc"
)

type UserRepository struct {
	q *db.Queries
}

func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{q: q}
}

// Create user
func (r *UserRepository) Create(ctx context.Context, name, dob string) (db.User, error) {
	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return db.User{}, err
	}

	return r.q.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob:  parsedDOB,
	})
}

// Get user by ID
func (r *UserRepository) GetByID(ctx context.Context, id int32) (db.User, error) {
	return r.q.GetUserByID(ctx, id)
}

// List all users
func (r *UserRepository) List(ctx context.Context) ([]db.User, error) {
	return r.q.ListUsers(ctx)
}

// Update user
func (r *UserRepository) Update(ctx context.Context, id int32, name, dob string) (db.User, error) {
	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return db.User{}, err
	}

	return r.q.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  parsedDOB,
	})
}

// Delete user
func (r *UserRepository) Delete(ctx context.Context, id int32) error {
	return r.q.DeleteUser(ctx, id)
}
