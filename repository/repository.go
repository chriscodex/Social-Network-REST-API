package repository

import (
	"context"

	"github.com/ChrisCodeX/REST-API-Go/models"
)

type Repository interface {
	// Table User
	ValidateUserAlreadyRegistered(ctx context.Context, user *models.User) error
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	// Table Posts
	InsertPost(ctx context.Context, post *models.Post) error
	GetPostById(ctx context.Context, id string) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) error
	DeletePost(ctx context.Context, id string, userId string) error
	ListPost(ctx context.Context, page uint64, size uint64) ([]*models.Post, error)
	Close() error
}

// Assign Repository

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

// Table User Operations

func ValidateUserAlreadyRegistered(ctx context.Context, user *models.User) error {
	return implementation.ValidateUserAlreadyRegistered(ctx, user)
}

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

func GetPostById(ctx context.Context, id string) (*models.Post, error) {
	return implementation.GetPostById(ctx, id)
}

func UpdatePost(ctx context.Context, post *models.Post) error {
	return implementation.UpdatePost(ctx, post)
}

func DeletePost(ctx context.Context, id string, userId string) error {
	return implementation.DeletePost(ctx, id, userId)
}

func ListPost(ctx context.Context, page uint64, size uint64) ([]*models.Post, error) {
	return implementation.ListPost(ctx, page, size)
}
