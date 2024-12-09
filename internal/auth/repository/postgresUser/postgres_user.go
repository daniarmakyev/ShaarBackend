package postgresUser

import (
	"context"
	"database/sql"
	"fmt"
	"myfirstBack/internal/auth/model"
)

type PostgresUserRepository struct {
	DB *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{DB: db}
}

func (r *PostgresUserRepository) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	query := "SELECT id, username, email, password, avatar FROM users WHERE id = $1"
	row := r.DB.QueryRowContext(ctx, query, id)
	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar)
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
	query := "UPDATE users SET username = $1, email = $2, password = $3, avatar = $4 WHERE id = $5"
	_, err := r.DB.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.Avatar, user.ID)
	return err
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *model.User) error {
	query := "INSERT INTO users (username, email, password, avatar) VALUES ($1, $2, $3, $4)"
	_, err := r.DB.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.Avatar)
	return err
}

func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := "SELECT id, username, email, password, avatar FROM users WHERE email = $1"
	rows, err := r.DB.QueryContext(ctx, query, email)
	if err != nil {
		fmt.Printf("Error while fetching user: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar)
		if err != nil {
			fmt.Printf("Error while scanning user: %v\n", err)
			return nil, err
		}
		return &user, nil
	}

	return nil, nil
}
