package models

import "access_manager/internal/application/domain"

type RefreshToken struct {
	UserID         string
	String         string
	IssuedTime     int64
	ExpirationTime int64
}

func RefreshTokenFromDomain(refreshToken *domain.RefreshToken) (*RefreshToken, error) {
	return &RefreshToken{
		UserID:         string(refreshToken.UserID),
		String:         refreshToken.String,
		IssuedTime:     refreshToken.IssuedTime,
		ExpirationTime: refreshToken.ExpirationTime,
	}, nil
}

func (t *RefreshToken) ToDomain() *domain.RefreshToken {
	return &domain.RefreshToken{
		UserID:         domain.UserID(t.UserID),
		String:         t.String,
		IssuedTime:     t.IssuedTime,
		ExpirationTime: t.ExpirationTime,
	}
}
