package domain

import "context"

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	GetByID(ctx context.Context, userID UserID) (*User, error)
	Update(ctx context.Context, user *User) error
}
