package repository

import (
	"context"
	"myfirstBack/internal/auth/model"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id int64) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}
