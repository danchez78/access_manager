package usecases

import (
	"context"
	"fmt"
	"time"

	"access_manager/internal/application/domain"
)

type RefreshTokensHandler struct {
	userRepo  domain.UserRepository
	tokenRepo domain.TokenRepository
	tm        *domain.TokenManager
	as        *domain.AlertSender
}

func NewRefreshTokensHandler(
	userRepo domain.UserRepository,
	tokenRepo domain.TokenRepository,
	tm *domain.TokenManager,
	as *domain.AlertSender,
) *RefreshTokensHandler {
	if userRepo == nil {
		panic("user repository is nil")
	}

	if tokenRepo == nil {
		panic("token repository is nil")
	}

	if tm == nil {
		panic("token manager is nil")
	}

	if as == nil {
		panic("alert sender is nil")
	}

	return &RefreshTokensHandler{userRepo: userRepo, tokenRepo: tokenRepo, tm: tm, as: as}
}

func (h *RefreshTokensHandler) Execute(
	ctx context.Context,
	refreshTokenString string,
	ipAddress string,
) (domain.AccessToken, *domain.RefreshToken, error) {
	refreshToken, err := h.tm.ParseToken(refreshTokenString)
	if err != nil {
		return "", nil, err
	}

	user, err := h.userRepo.GetByID(ctx, domain.UserID(refreshToken.UserID))
	if err != nil {
		return "", nil, err
	}

	if user.IsDeauthorised {
		return "", nil, ErrDeauthorised
	}

	if user.IPAddress != ipAddress {
		// TODO: change ip address in db?
		err := h.as.SendAlert(string(user.ID), user.IPAddress, ipAddress)
		fmt.Println(err)
	}

	if refreshToken.ExpirationTime < time.Now().Unix() {
		return "", nil, ErrAccessTokenExpired
	}

	refreshTokenFromDB, err := h.tokenRepo.GetByUserID(ctx, domain.UserID(refreshToken.UserID))
	if err != nil {
		return "", nil, err
	}

	if refreshToken.IssuedTime != refreshTokenFromDB.IssuedTime {
		return "", nil, ErrTokenNotInUse
	}

	accessTokenDomain, refreshTokenDomain, err := h.tm.GenerateTokens(domain.UserID(refreshToken.UserID))
	if err != nil {
		return "", nil, err
	}

	if err := h.tokenRepo.Update(ctx, refreshTokenDomain); err != nil {
		return "", nil, err
	}

	return accessTokenDomain, refreshTokenDomain, nil

}
