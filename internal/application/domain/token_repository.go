package domain

import "context"

type TokenRepository interface {
	Save(ctx context.Context, refreshToken *RefreshToken) error
	GetByUserID(ctx context.Context, userID UserID) (*RefreshToken, error)
	Update(ctx context.Context, refreshToken *RefreshToken) error
}
