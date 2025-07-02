package usecases

import (
	"access_manager/internal/application/domain"
)

type UseCases struct {
	CreateUserHandler      *CreateUserHandler
	GetIDHandler           *GetIDHandler
	GetTokensHandler       *GetTokensHandler
	RefreshTokensHandler   *RefreshTokensHandler
	DeauthoriseUserHandler *DeauthoriseUserHandler
}

func NewUseCases(
	userRepo domain.UserRepository,
	tokenRepo domain.TokenRepository,
	tm *domain.TokenManager,
	as *domain.AlertSender,
) *UseCases {
	return &UseCases{
		CreateUserHandler:      NewCreateUserHandler(userRepo),
		GetIDHandler:           NewGetIDHandler(userRepo, tokenRepo, tm),
		GetTokensHandler:       NewGetTokensHandler(userRepo, tokenRepo, tm),
		RefreshTokensHandler:   NewRefreshTokensHandler(userRepo, tokenRepo, tm, as),
		DeauthoriseUserHandler: NewDeauthoriseUserHandler(userRepo, tokenRepo, tm),
	}
}
