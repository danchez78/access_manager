package usecases

import "errors"

var (
	ErrTokensNotIssuedTogether = errors.New("provided tokens issued not together")
	ErrAccessTokenExpired      = errors.New("provided access token expired")
	ErrRefreshTokenExpired     = errors.New("provided refresh token expired")
	ErrUnknownUserAgent        = errors.New("provided user agent not used before")
	ErrUnknownIPAddress        = errors.New("provided ip address not used before")
	ErrTokenNotInUse           = errors.New("provided refresh token is no longer in use")
	ErrDeauthorised            = errors.New("provided user was deauthorised")
)
