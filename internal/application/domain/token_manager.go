package domain

import (
	"access_manager/internal/common/token_generator"
)

type TokenManager struct {
	tg *token_generator.TokenGenerator
}

func NewTokenManager(tg *token_generator.TokenGenerator) *TokenManager {
	if tg == nil {
		panic("token generator is nil")
	}

	return &TokenManager{tg: tg}
}

func (tm *TokenManager) GenerateTokens(userID UserID) (AccessToken, *RefreshToken, error) {
	accessToken, refreshToken, err := tm.tg.GenerateTokens(string(userID))
	if err != nil {
		return "", nil, err
	}

	return AccessToken(accessToken.String), NewRefreshToken(
		userID,
		refreshToken.String,
		refreshToken.IssuedTime,
		refreshToken.ExpirationTime,
	), nil
}

func (tm *TokenManager) ParseToken(tokenString string) (*token_generator.Token, error) {
	token, err := tm.tg.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	return token, nil
}
