package usecases

import (
	"context"
	"time"

	"access_manager/internal/application/domain"
)

type GetIDHandler struct {
	userRepo  domain.UserRepository
	tokenRepo domain.TokenRepository
	tm        *domain.TokenManager
}

func NewGetIDHandler(
	userRepo domain.UserRepository,
	tokenRepo domain.TokenRepository,
	tm *domain.TokenManager,
) *GetIDHandler {
	if userRepo == nil {
		panic("user repository is nil")
	}

	if tokenRepo == nil {
		panic("token repository is nil")
	}

	if tm == nil {
		panic("token manager is nil")
	}

	return &GetIDHandler{userRepo: userRepo, tokenRepo: tokenRepo, tm: tm}
}

func (h *GetIDHandler) Execute(
	ctx context.Context,
	accessTokenString string,
) (domain.UserID, error) {
	accessToken, err := h.tm.ParseToken(accessTokenString)
	if err != nil {
		return "", err
	}

	user, err := h.userRepo.GetByID(ctx, domain.UserID(accessToken.UserID))
	if err != nil {
		return "", err
	}

	if user.IsDeauthorised {
		return "", ErrDeauthorised
	}

	refreshToken, err := h.tokenRepo.GetByUserID(ctx, domain.UserID(accessToken.UserID))
	if err != nil {
		return "", err
	}

	if !(accessToken.UserID == string(refreshToken.UserID) && accessToken.IssuedTime == refreshToken.IssuedTime) {
		return "", ErrTokensNotIssuedTogether
	}

	if accessToken.ExpirationTime < time.Now().Unix() {
		return "", ErrAccessTokenExpired
	}

	return domain.UserID(accessToken.UserID), nil
}
