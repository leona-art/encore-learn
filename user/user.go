package user

import (
	"context"
	"fmt"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Users []*User

func (u *Users) GetUser(ctx context.Context, id int) (*User, error) {
	user := (*u)[id]
	if user == nil {
		return nil, fmt.Errorf("user with ID %d not found", id)
	}
	return user, nil
}

func (u *Users) CreateUser(ctx context.Context, name string) (*User, error) {
	id := len(*u) + 1
	user := &User{
		ID:   id,
		Name: name,
	}
	*u = append(*u, user)
	return user, nil
}

type CreateUserRequest struct {
	Name string `json:"name"`
}
type CreateUserResponse User

type UserRepository interface {
	CreateUser(ctx context.Context, name string) (*User, error)
	GetUser(ctx context.Context, id int) (*User, error)
}

var user_repository UserRepository

func init() {
	user_repository = &Users{}
}

//encore:api public method=POST path=/user
func CreateUser(ctx context.Context, params *CreateUserRequest) (*User, error) {
	createdUser, err := user_repository.CreateUser(ctx, params.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return createdUser, nil
}

//encore:api public method=GET path=/user/:id
func GetUser(ctx context.Context, id int) (*User, error) {
	user, err := user_repository.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user with ID %d not found", id)
	}
	return user, nil
}
