package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ChrisCodeX/REST-API-Go/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

// Function that closes the connection
func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}

// Constructor of Postgres Repository
func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

// Methods that makes PostgresRepository a Repository
// Table User Operations

/* Validate User Already Registered in Database */
func (repo *PostgresRepository) ValidateUserAlreadyRegistered(ctx context.Context, user *models.User) error {
	//Validation (Email already registered)
	// Query
	rows, err := repo.db.QueryContext(ctx, "SELECT id FROM users WHERE email = $1", user.Email)
	if err != nil {
		return err
	}

	// Stop reading rows
	defer CloseReadingRows(rows)

	// Validate
	if rows.Next() {
		return fmt.Errorf("email is already registered")
	}
	return nil
}

/* Insert a User
* @param {context} ctx context

* @param {*models.User} user struct that will be inserted in the database

* @return {error} If the insert fails, returns an error
 */
func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	/* Insertion */
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password)
	return err
}

// Get user sending an ID
func (repo *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	// Stop reading rows
	defer CloseReadingRows(rows)

	// Map column values of row into the user struct
	var user = models.User{}

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

// Get user sending the Email
func (repo *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", email)

	if err != nil {
		return nil, err
	}

	// Stop reading rows
	defer CloseReadingRows(rows)

	// Map column values of row into the user struct
	var user = models.User{}

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email, &user.Password); err == nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

// Table Posts Operations

/* Insert a Post
* @param {context} ctx context

* @param {*models.Post} post struct that will be inserted in the database

* @return {error} If the insert fails, returns an error
 */
func (repo *PostgresRepository) InsertPost(ctx context.Context, post *models.Post) error {
	// Insertion
	_, err := repo.db.ExecContext(ctx, "INSERT INTO posts (id, post_content, user_id) VALUES ($1, $2, $3)",
		post.Id, post.PostContent, post.UserId)
	return err
}

// Get user sending an ID
func (repo *PostgresRepository) GetPostById(ctx context.Context, id string) (*models.Post, error) {

	// Query
	rows, err := repo.db.QueryContext(ctx, "SELECT id, post_content, created_at, user_id FROM posts WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	// Stop reading rows
	defer CloseReadingRows(rows)

	// Map column values of row into the post struct
	var post = models.Post{}

	for rows.Next() {
		if err = rows.Scan(&post.Id, &post.PostContent, &post.CreatedAt, &post.UserId); err == nil {
			return &post, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &post, nil
}

func (repo *PostgresRepository) UpdatePost(ctx context.Context, post *models.Post) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE posts SET post_content = $1 WHERE id = $2 and user_id = $3",
		post.PostContent, post.Id, post.UserId)
	return err
}

func (repo *PostgresRepository) DeletePost(ctx context.Context, id string, userId string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM posts WHERE id = $1 and user_id = $2",
		id, userId)
	return err
}

func (repo *PostgresRepository) ListPost(ctx context.Context, page uint64, size uint64) ([]*models.Post, error) {
	// Query
	rows, err := repo.db.QueryContext(ctx, "SELECT id, post_content, user_id, created_at FROM posts LIMIT $1 OFFSET ($2-1)", size, page*size)
	if err != nil {
		return nil, err
	}

	// Stop Reading Rows
	defer CloseReadingRows(rows)

	// Add post into the slice
	var posts []*models.Post
	for rows.Next() {
		var post = models.Post{}
		if err = rows.Scan(&post.Id, &post.PostContent, &post.UserId, &post.CreatedAt); err == nil {
			posts = append(posts, &post)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
