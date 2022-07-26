package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/ChrisCodeX/REST-API-Go/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

// Constructor of Postgres Repository
func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

// Methods that makes PostgresRepository a User Repository

/* Insert a User
* @param {context} ctx context

* @param {*models.User} user struct that will be inserted in postgres database

* @return {error} If the insert fails, returns an error
 */
func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	// Insertion
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password)
	return err
}

// Get user sending a ID
func (repo *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)

	// Stop reading rows
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

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

	// Stop reading rows
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

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

// Function that closes the connection
func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
