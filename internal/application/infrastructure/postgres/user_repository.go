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

type userRepository struct {
	client *postgres_client.Client
}

var (
	_ domain.UserRepository = (*userRepository)(nil)
)

var (
	ErrNoProvidedUser = errors.New("user with provided id not found")
)

func NewUserRepository(client *postgres_client.Client) domain.UserRepository {
	if client == nil {
		panic("postgres client is nil")
	}

	return &userRepository{client: client}
}

func (r *userRepository) Save(ctx context.Context, user *domain.User) error {
	mUser := models.UserFromDomain(user)

	q := "INSERT INTO access_manager.users (id, user_agent, ip_address, is_deauthorised) VALUES ($1, $2, $3, $4)"
	if _, err := r.client.Exec(ctx, q, mUser.ID, mUser.UserAgent, mUser.IPAddress, mUser.IsDeauthorised); err != nil {
		return fmt.Errorf("inserting user got error: %v", err)
	}

	return nil
}

func (r *userRepository) GetByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	var mUser models.User

	q := `
	SELECT id, user_agent, ip_address, is_deauthorised
	FROM access_manager.users
	WHERE id = $1
	`
	if err := r.client.QueryRow(ctx, q, userID).Scan(&mUser.ID, &mUser.UserAgent, &mUser.IPAddress, &mUser.IsDeauthorised); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoProvidedUser
		} else {
			return nil, fmt.Errorf("getting user by id got error: %v", err)
		}
	}
	return mUser.ToDomain(), nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	mUser := models.UserFromDomain(user)

	q := `
	UPDATE access_manager.users
	SET user_agent = $2, ip_address = $3, is_deauthorised = $4
	WHERE id = $1
	`
	if _, err := r.client.Exec(
		ctx,
		q,
		mUser.ID,
		mUser.UserAgent,
		mUser.IPAddress,
		mUser.IsDeauthorised,
	); err != nil {
		return fmt.Errorf("updating user got error: %v", err)
	}

	return nil
}
