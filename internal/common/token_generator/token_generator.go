package token_generator

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenGenerator struct {
	secretkey              []byte
	accessTokenTTLMinutes  int
	refreshTokenTTLMinutes int
}

func NewTokenGenerator(cfg Config) *TokenGenerator {
	return &TokenGenerator{
		secretkey:              []byte(cfg.SecretKey),
		accessTokenTTLMinutes:  cfg.AccessTokenTTLMinutes,
		refreshTokenTTLMinutes: cfg.RefreshTokenTTLMinutes,
	}
}

func (g *TokenGenerator) GenerateTokens(userID string) (*Token, *Token, error) {
	issuedTime := time.Now().Unix()
	accessTokenExpirationTime := time.Now().Add(time.Minute * time.Duration(g.accessTokenTTLMinutes)).Unix()
	refreshTokenExpirationTime := time.Now().Add(time.Minute * time.Duration(g.refreshTokenTTLMinutes)).Unix()

	accessTokenClaims := jwt.MapClaims{
		"user_id":         userID,
		"issued_time":     issuedTime,
		"expiration_time": accessTokenExpirationTime,
	}

	accessToken, err := g.generateToken(accessTokenClaims)
	if err != nil {
		return nil, nil, err
	}

	refreshTokenClaims := jwt.MapClaims{
		"user_id":         userID,
		"issued_time":     issuedTime,
		"expiration_time": refreshTokenExpirationTime,
	}

	refreshToken, err := g.generateToken(refreshTokenClaims)
	if err != nil {
		return nil, nil, err
	}

	return &Token{
			UserID:         userID,
			String:         accessToken,
			IssuedTime:     issuedTime,
			ExpirationTime: accessTokenExpirationTime,
		}, &Token{
			UserID:         userID,
			String:         refreshToken,
			IssuedTime:     issuedTime,
			ExpirationTime: refreshTokenExpirationTime,
		}, nil
}

func (g *TokenGenerator) ParseToken(tokenString string) (*Token, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(g.secretkey), nil
	})
	if err != nil {
		return nil, err
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.New("provided token with incorrect user id type")
	}

	issuedTime, ok := claims["issued_time"].(float64)
	if !ok {
		return nil, errors.New("provided token with incorrect issued time type")
	}

	expirationTime, ok := claims["expiration_time"].(float64)
	if !ok {
		return nil, errors.New("provided token with incorrect expiration type")
	}

	return &Token{
		UserID:         userID,
		String:         tokenString,
		IssuedTime:     int64(issuedTime),
		ExpirationTime: int64(expirationTime),
	}, nil
}

func (g *TokenGenerator) generateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenString, err := token.SignedString(g.secretkey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
