package postgres

import (
	"context"
	"errors"
	"fmt"

	"access_manager/internal/application/domain"
	"access_manager/internal/application/infrastructure/postgres/models"
	"access_manager/internal/common/postgres_client"

	"github.com/jackc/pgx/v5"
)

type tokenRepository struct {
	client *postgres_client.Client
}

var (
	_ domain.TokenRepository = (*tokenRepository)(nil)
)

func NewTokenRepository(client *postgres_client.Client) domain.TokenRepository {
	if client == nil {
		panic("postgres client is nil")
	}

	return &tokenRepository{client: client}
}

func (r *tokenRepository) Save(ctx context.Context, refreshToken *domain.RefreshToken) error {
	mRefreshToken, err := models.RefreshTokenFromDomain(refreshToken)
	if err != nil {
		return err
	}

	q := `
	INSERT INTO access_manager.refresh_tokens (user_id, issued_time, expiration_time, string)
	VALUES ($1, $2, $3, $4)`

	if _, err := r.client.Exec(
		ctx,
		q,
		mRefreshToken.UserID,
		mRefreshToken.IssuedTime,
		mRefreshToken.ExpirationTime,
		mRefreshToken.String,
	); err != nil {
		return fmt.Errorf("inserting refresh token got error: %v", err)
	}

	return nil
}

func (r *tokenRepository) GetByUserID(ctx context.Context, userID domain.UserID) (*domain.RefreshToken, error) {
	var mRefreshToken models.RefreshToken

	q := `
	SELECT user_id, issued_time, expiration_time, string
	FROM access_manager.refresh_tokens
	WHERE user_id = $1`
	if err := r.client.QueryRow(ctx, q, userID).Scan(
		&mRefreshToken.UserID,
		&mRefreshToken.IssuedTime,
		&mRefreshToken.ExpirationTime,
		&mRefreshToken.String,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		} else {
			return nil, fmt.Errorf("getting user by id got error: %v", err)
		}
	}
	return mRefreshToken.ToDomain(), nil
}

func (r *tokenRepository) Update(ctx context.Context, refreshToken *domain.RefreshToken) error {
	mRefreshToken, err := models.RefreshTokenFromDomain(refreshToken)
	if err != nil {
		return err
	}

	q := `
	UPDATE access_manager.refresh_tokens
	SET issued_time = $2, expiration_time = $3, string = $4
	WHERE user_id = $1
	`
	if _, err := r.client.Exec(
		ctx,
		q,
		mRefreshToken.UserID,
		mRefreshToken.IssuedTime,
		mRefreshToken.ExpirationTime,
		mRefreshToken.String,
	); err != nil {
		return fmt.Errorf("updating refresh token got error: %v", err)
	}

	return nil
}
