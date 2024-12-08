package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"myfirstBack/internal/auth/model"
	"myfirstBack/internal/auth/repository"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) CreateUser(ctx context.Context, user *model.User) error {
	existingUser, err := u.repo.GetUserByEmail(ctx, user.Email)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error checking existing user: %v", err)
	}
	if existingUser != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	return u.repo.CreateUser(ctx, user)
}
