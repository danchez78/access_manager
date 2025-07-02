package views

import "access_manager/internal/application/domain"

type CreateUserResponse struct {
	UserID string `json:"user_id"`
}

func NewCreateUserResponse(userID domain.UserID) CreateUserResponse {
	return CreateUserResponse{UserID: string(userID)}
}

type GetTokensResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewGetTokensResponse(
	accessToken domain.AccessToken,
	refreshToken *domain.RefreshToken,
) GetTokensResponse {
	return GetTokensResponse{
		AccessToken:  string(accessToken),
		RefreshToken: refreshToken.String,
	}
}

type RefreshTokensResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewRefreshTokensResponse(
	accessToken domain.AccessToken,
	refreshToken *domain.RefreshToken,
) RefreshTokensResponse {
	return RefreshTokensResponse{
		AccessToken:  string(accessToken),
		RefreshToken: refreshToken.String,
	}
}

type GetIDResponse struct {
	UserID string
}

func NewGetIDResponse(userID domain.UserID) GetIDResponse {
	return GetIDResponse{
		UserID: string(userID),
	}
}

type DeauthoriseUserResponse struct{}

func NewDeauthoriseUserResponse() DeauthoriseUserResponse {
	return DeauthoriseUserResponse{}
}
