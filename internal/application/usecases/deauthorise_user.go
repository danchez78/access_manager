package usecases

import (
	"context"
	"time"

	"access_manager/internal/application/domain"
)

type DeauthoriseUserHandler struct {
	userRepo  domain.UserRepository
	tokenRepo domain.TokenRepository
	tm        *domain.TokenManager
}

func NewDeauthoriseUserHandler(
	userRepo domain.UserRepository,
	tokenRepo domain.TokenRepository,
	tm *domain.TokenManager,
) *DeauthoriseUserHandler {
	if userRepo == nil {
		panic("user repository is nil")
	}

	if tokenRepo == nil {
		panic("token repository is nil")
	}

	if tm == nil {
		panic("token manager is nil")
	}

	return &DeauthoriseUserHandler{userRepo: userRepo, tokenRepo: tokenRepo, tm: tm}
}

func (h *DeauthoriseUserHandler) Execute(
	ctx context.Context,
	accessTokenString string,
) error {
	accessToken, err := h.tm.ParseToken(accessTokenString)
	if err != nil {
		return err
	}

	refreshToken, err := h.tokenRepo.GetByUserID(ctx, domain.UserID(accessToken.UserID))
	if err != nil {
		return err
	}

	if !(accessToken.UserID == string(refreshToken.UserID) && accessToken.IssuedTime == refreshToken.IssuedTime) {
		return ErrTokensNotIssuedTogether
	}

	if accessToken.ExpirationTime < time.Now().Unix() {
		return ErrAccessTokenExpired
	}

	user, err := h.userRepo.GetByID(ctx, domain.UserID(accessToken.UserID))
	if err != nil {
		return err
	}

	user.IsDeauthorised = true
	return h.userRepo.Update(ctx, user)
}
