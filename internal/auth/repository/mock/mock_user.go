package mock

import (
	"context"
	"myfirstBack/internal/auth/model"
)

type MockUserRepository struct{}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

func (r *MockUserRepository) CreateUser(ctx context.Context, user *model.User) error {

	return nil
}

func (r *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {

	return &model.User{ID: 1, Name: "John", Email: email}, nil
}

func (r *MockUserRepository) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	return &model.User{ID: id, Name: "John", Email: "john@example.com"}, nil
}

func (r *MockUserRepository) DeleteUser(ctx context.Context, id int64) error {
	return nil
}

func (r *MockUserRepository) UpdateUser(ctx context.Context, user *model.User) error {
	return nil
}
