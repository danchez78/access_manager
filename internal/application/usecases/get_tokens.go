package usecases

import (
	"access_manager/internal/application/domain"
	"context"
)

type GetTokensHandler struct {
	userRepo  domain.UserRepository
	tokenRepo domain.TokenRepository
	tm        *domain.TokenManager
}

func NewGetTokensHandler(
	userRepo domain.UserRepository,
	tokenRepo domain.TokenRepository,
	tm *domain.TokenManager,
) *GetTokensHandler {
	if userRepo == nil {
		panic("user repository is nil")
	}

	if tokenRepo == nil {
		panic("token repository is nil")
	}

	if tm == nil {
		panic("token manager is nil")
	}

	return &GetTokensHandler{userRepo: userRepo, tokenRepo: tokenRepo, tm: tm}
}

func (h *GetTokensHandler) Execute(ctx context.Context, userID domain.UserID) (domain.AccessToken, *domain.RefreshToken, error) {
	_, err := h.userRepo.GetByID(ctx, userID)
	if err != nil {
		return "", nil, err
	}

	accessToken, refreshToken, err := h.tm.GenerateTokens(userID)
	if err != nil {
		return "", nil, err
	}

	previousRefreshToken, err := h.tokenRepo.GetByUserID(ctx, userID)
	if err != nil {
		return "", nil, err
	}

	if previousRefreshToken == nil {
		if err := h.tokenRepo.Save(ctx, refreshToken); err != nil {
			return "", nil, err
		}
	} else {
		if err := h.tokenRepo.Update(ctx, refreshToken); err != nil {
			return "", nil, err
		}
	}

	return accessToken, refreshToken, nil
}
