package repository

import (
	"context"

	"github.com/ChrisCodeX/REST-API-Go/models"
)

type Repository interface {
	// Table User
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	// Table Posts
	InsertPost(ctx context.Context, post *models.Post) error
	Close() error
}

// Assign Repository

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

// Table User Operations

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

// Table Post Operations

func InsertPost(ctx context.Context, post *models.Post) error {
	return implementation.InsertPost(ctx, post)
}
