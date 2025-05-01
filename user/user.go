package user

import (
	"context"
	"fmt"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Users map[string]*User

func (u Users) GetUser(ctx context.Context, id string) (*User, error) {
	user, exists := u[id]
	if !exists {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
}

func (u Users) CreateUser(ctx context.Context, user *User) (*User, error) {
	if _, exists := u[user.ID]; exists {
		return nil, fmt.Errorf("user with ID %s already exists", user.ID)
	}
	u[user.ID] = user
	return user, nil
}

type CreateUserRequest struct {
	Name string `json:"name"`
}
type CreateUserResponse User

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
}

var user_repository UserRepository

func init() {
	user_repository = make(Users)
}

//encore:api public method=POST path=/user
func CreateUser(ctx context.Context, params *CreateUserRequest) (*CreateUserResponse, error) {
	user := &CreateUserResponse{
		ID:   "1",
		Name: params.Name,
	}

	createdUser, err := user_repository.CreateUser(ctx, &User{
		ID:   user.ID,
		Name: user.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return (*CreateUserResponse)(createdUser), nil
}

//encore:api public method=GET path=/user/:id
func GetUser(ctx context.Context, id string) (*User, error) {
	user, err := user_repository.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
}
