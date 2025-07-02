package controllers

import "access_manager/internal/application/domain"

type CreateUserRequest struct{}

type AccessTokenController struct {
	AccessToken string `json:"access_token"`
}

type RefreshTokenController struct {
	RefreshToken string `json:"refresh_token"`
}

type GetTokensRequest struct {
	UserID domain.UserID `param:"user_id"`
}

type RefreshTokensRequest struct {
	RefreshTokenController
}

type GetIDRequest struct {
	AccessTokenController
}

type DeauthoriseUserRequest struct {
	AccessTokenController
}
