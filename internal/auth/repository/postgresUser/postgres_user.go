package postgresUser

import (
	"context"
	"database/sql"
	"myfirstBack/internal/auth/model"
)

type PostgresUserRepository struct {
	DB *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{DB: db}
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *model.User) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2)"
	_, err := r.DB.ExecContext(ctx, query, user.Name, user.Email)
	return err
}

func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := "SELECT id, name, email FROM users WHERE email = $1"
	row := r.DB.QueryRowContext(ctx, query, email)
	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = $1"
	row := r.DB.QueryRowContext(ctx, query, id)
	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) DeleteUser(ctx context.Context, id int64) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}

func (r *PostgresUserRepository) UpdateUser(ctx context.Context, user *model.User) error {
	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
	_, err := r.DB.ExecContext(ctx, query, user.Name, user.Email, user.ID)
	return err
}
